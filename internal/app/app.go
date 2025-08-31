package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/charmtone"
	"github.com/nixpig/virtui/internal/libvirtui"
	"github.com/nixpig/virtui/internal/messages"
	"github.com/nixpig/virtui/internal/screen"
)

var _ tea.Model = (*model)(nil)

var footerStyle = lipgloss.NewStyle().Padding(0, 1, 1)

// model holds the global state and rendering logic.
type model struct {
	currentScreen    screen.Screen
	screens          map[string]screen.Screen
	width, height    int
	globalKeys       GlobalKeyMap
	keyMapViewHeight int
	help             help.Model
	events           chan libvirtui.DomainEvent
	conn             libvirtui.Connection
	service          libvirtui.Service
	combinedKeys     combinedKeyMap
}

func listenForEvents(ch <-chan libvirtui.DomainEvent) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}

// NewAppModel returns an initialised app model.
func NewAppModel(
	conn libvirtui.Connection,
	service libvirtui.Service,
	screens []screen.Screen,
) *model {
	m := &model{
		globalKeys: DefaultGlobalKeyMap(),
		help:       help.New(),
		screens:    make(map[string]screen.Screen),
		conn:       conn,
		service:    service,
		events:     make(chan libvirtui.DomainEvent),
	}

	for _, screen := range screens {
		m.screens[screen.ID()] = screen
	}

	m.currentScreen = m.screens[screens[0].ID()]

	m.combinedKeys = combinedKeyMap{
		global: m.globalKeys,
		screen: m.currentScreen.HelpKeys(),
		currentScreenID: m.currentScreen.ID(),
	}

	// temporarily rendering the footer to calculate the height offset
	tempFooter := footerStyle.Render(
		m.help.FullHelpView(m.combinedKeys.FullHelp()),
	)
	m.keyMapViewHeight = lipgloss.Height(tempFooter)

	if _, err := m.conn.DomainEventLifecycleRegister(
		func(event libvirtui.DomainEvent) {
			log.Debug("handling domain event", "event", event.Event, "detail", event.Detail)
			m.events <- event
		},
	); err != nil {
		log.Error("failed to register domain event handler", "err", err)
	}

	return m
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		tea.ClearScreen,
		listenForEvents(m.events),
		getDomainsCmd(m.service),
		// TODO: review whether I really want to be getting pools and networks on
		// init before either screen can be displayed or whether to lazily fetch
		getStoragePoolsCmd(m.service),
		getNetworksCmd(m.service),
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case messages.DomainActionWithFuncMsg:
		go func() {
			if err := msg.Action(m.service, msg.DomainUUID); err != nil {
				log.Error("failed to perform domain action", "err", err)
			}
		}()

	case messages.DomainActionMsg:
		if msg.IsPrompt {
			log.Info("prompt action not yet implemented")
		} else {
			switch msg.Action {
			case "start":
				go func() {
					if err := m.service.DomainStart(msg.Domain); err != nil {
						log.Error("failed to start domain", "err", err)
					}
				}()
			case "pause/resume":
				go func() {
					if err := m.service.ToggleDomainState(msg.Domain); err != nil {
						log.Error("failed to toggle domain state", "err", err)
					}
				}()
			case "shutdown":
				go func() {
					if err := m.service.DomainShutdown(msg.Domain); err != nil {
						log.Error("failed to shutdown domain", "err", err)
					}
				}()
			case "reboot":
				go func() {
					if err := m.service.DomainReboot(msg.Domain); err != nil {
						log.Error("failed to reboot domain", "err", err)
					}
				}()
			case "reset":
				go func() {
					if err := m.service.ResetDomain(msg.Domain); err != nil {
						log.Error("failed to reset domain", "err", err)
					}
				}()
			case "off":
				go func() {
					if err := m.service.ForceOffDomain(msg.Domain); err != nil {
						log.Error("failed to force off domain", "err", err)
					}
				}()
			}
		}

	case libvirtui.DomainEvent:
		// just re-fetch all the domain details for the time-being
		cmds = append(cmds, getDomainsCmd(m.service))

	case messages.DomainsMsg:
		if m.currentScreen.ID() == "manager" {
			var screenCmd tea.Cmd
			var updatedModel tea.Model

			updatedModel, screenCmd = m.currentScreen.Update(msg)
			m.currentScreen = updatedModel.(screen.Screen)

			cmds = append(cmds, screenCmd)
		}

	case messages.StoragePoolsMsg, messages.StorageVolumesMsg:
		storageScreen, ok := m.screens["storage"]
		if ok {
			var screenCmd tea.Cmd
			var updatedModel tea.Model

			updatedModel, screenCmd = storageScreen.Update(msg)
			m.screens["storage"] = updatedModel.(screen.Screen)
			cmds = append(cmds, screenCmd)
		}

		if m.currentScreen.ID() == "storage" {
			// if storage is the current screen, also update the currentScreen reference
			m.currentScreen = m.screens["storage"]
		}

	case messages.NetworksMsg:
		networkScreen, ok := m.screens["network"]
		if ok {
			var screenCmd tea.Cmd
			var updatedModel tea.Model

			updatedModel, screenCmd = networkScreen.Update(msg)
			m.screens["network"] = updatedModel.(screen.Screen)
			cmds = append(cmds, screenCmd)
		}

		if m.currentScreen.ID() == "network" {
			// if network is the current screen, also update the currentScreen reference
			m.currentScreen = m.screens["network"]
		}

	case messages.ScreenSizeMsg:
		if m.currentScreen != nil {
			var screenCmd tea.Cmd
			var updatedModel tea.Model

			updatedModel, screenCmd = m.currentScreen.Update(msg)
			m.currentScreen = updatedModel.(screen.Screen)

			cmds = append(cmds, screenCmd)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = m.width

		availableScreenHeight := max(m.height-1-m.keyMapViewHeight, 0)

		if m.currentScreen != nil {
			var screenCmd tea.Cmd
			var updatedModel tea.Model

			updatedModel, screenCmd = m.currentScreen.Update(messages.ScreenSizeMsg{
				Width:  m.width,
				Height: availableScreenHeight,
			})

			m.currentScreen = updatedModel.(screen.Screen)

			cmds = append(cmds, screenCmd)
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.globalKeys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.globalKeys.Help):
		// TODO: pop a dialog

		case key.Matches(msg, m.globalKeys.DashboardScreen):
			cmd := m.switchScreen("manager")
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.globalKeys.StorageScreen):
			cmd := m.switchScreen("storage")
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.globalKeys.NetworksScreen):
			cmd := m.switchScreen("network")
			cmds = append(cmds, cmd)
		}
	}

	m.help, cmd = m.help.Update(msg)
	cmds = append(cmds, cmd)

	if m.currentScreen != nil {
		var screenCmd tea.Cmd
		var updatedModel tea.Model

		updatedModel, screenCmd = m.currentScreen.Update(msg)
		m.currentScreen = updatedModel.(screen.Screen)

		cmds = append(cmds, screenCmd)
	}

	cmds = append(cmds, listenForEvents(m.events))

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	if m.currentScreen == nil {
		return "Loading..."
	}

	titleStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(charmtone.Charple.Hex())).
		Foreground(lipgloss.Color(charmtone.Zest.Hex())).
		Padding(0).
		Width(m.width).
		Align(lipgloss.Center)

	header := titleStyle.Render(m.currentScreen.Title())

	footer := footerStyle.Render(
		m.help.FullHelpView(m.combinedKeys.FullHelp()),
	)

	availableScreenHeight := max(
		m.height-lipgloss.Height(header)-lipgloss.Height(footer),
		0,
	)

	contentStyle := lipgloss.NewStyle().
		Height(availableScreenHeight).
		Width(m.width)

	mainContent := contentStyle.Render(m.currentScreen.View())

	return lipgloss.JoinVertical(lipgloss.Left, header, mainContent, footer)
}

func (m *model) switchScreen(screenName string) tea.Cmd {
	nextScreen, ok := m.screens[screenName]
	if !ok {
		return nil
	}

	m.currentScreen = nextScreen
	m.combinedKeys.currentScreenID = nextScreen.ID()

	availableScreenHeight := max(m.height-1-m.keyMapViewHeight, 0)

	var screenCmd tea.Cmd
	var updatedModel tea.Model

	updatedModel, screenCmd = m.currentScreen.Update(
		messages.ScreenSizeMsg{
			Width:  m.width,
			Height: availableScreenHeight,
		},
	)

	m.currentScreen = updatedModel.(screen.Screen)

	return screenCmd
}

func getDomainsCmd(service libvirtui.Service) tea.Cmd {
	return func() tea.Msg {
		domains, err := service.ListAllDomains()
		if err != nil {
			log.Error("failed to list all domains", "err", err)
			return nil
		}

		return messages.DomainsMsg{Domains: domains}
	}
}

func getStoragePoolsCmd(service libvirtui.Service) tea.Cmd {
	return func() tea.Msg {
		pools, err := service.ListAllStoragePools()
		if err != nil {
			log.Error("failed to list all storage pools", "err", err)
			return nil
		}

		return messages.StoragePoolsMsg{Pools: pools}
	}
}

func getNetworksCmd(service libvirtui.Service) tea.Cmd {
	return func() tea.Msg {
		networks, err := service.ListAllNetworks()
		if err != nil {
			log.Error("failed to list all networks", "err", err)
			return nil
		}

		return messages.NetworksMsg{Networks: networks}
	}
}
