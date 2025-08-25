package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/charmtone"
	"github.com/nixpig/virtui/internal/common"
	"github.com/nixpig/virtui/internal/libvirtui"
	"github.com/nixpig/virtui/internal/messages"
)

var _ tea.Model = (*model)(nil)

var footerStyle = lipgloss.NewStyle().Padding(0, 1, 1)

// Screen is the interface that all application screens must implement.
type Screen interface {
	tea.Model
	ID() string
	Title() string
	Keybindings() []key.Binding
	ScrollKeys() common.ScrollKeyMap
}

// model holds the global state and rendering logic.
type model struct {
	currentScreen    Screen
	screens          map[string]Screen
	width, height    int
	globalKeyMap     GlobalKeyMap
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
	screens []Screen,
) *model {
	m := &model{
		globalKeyMap: DefaultGlobalKeyMap(),
		help:         help.New(),
		screens:      make(map[string]Screen),
		conn:         conn,
		service:      service,
		events:       make(chan libvirtui.DomainEvent),
	}

	for _, screen := range screens {
		m.screens[screen.ID()] = screen
	}

	m.currentScreen = m.screens[screens[0].ID()]

	m.combinedKeys = combinedKeyMap{
		global: m.globalKeyMap,
		screen: m.currentScreen.Keybindings(),
		scroll: m.currentScreen.ScrollKeys(),
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
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case libvirtui.DomainEvent:
		// just re-fetch all the domain details for the time-being
		cmds = append(cmds, getDomainsCmd(m.service))

	case messages.DomainsMsg:
		if m.currentScreen.ID() == "manager" {
			var screenCmd tea.Cmd
			var updatedModel tea.Model

			updatedModel, screenCmd = m.currentScreen.Update(msg)
			m.currentScreen = updatedModel.(Screen)

			cmds = append(cmds, screenCmd)
		}

	case messages.ScreenSizeMsg:
		if m.currentScreen != nil {
			var screenCmd tea.Cmd
			var updatedModel tea.Model

			updatedModel, screenCmd = m.currentScreen.Update(msg)
			m.currentScreen = updatedModel.(Screen)

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

			updatedModel, screenCmd = m.currentScreen.Update(messages.ScreenSizeMsg{Width: m.width, Height: availableScreenHeight})
			m.currentScreen = updatedModel.(Screen)

			cmds = append(cmds, screenCmd)
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.globalKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.globalKeyMap.Help):
		// TODO: pop a dialog

		case key.Matches(msg, m.globalKeyMap.Screen1):
			cmd := m.switchScreen("manager")
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.globalKeyMap.Screen2):
			cmd := m.switchScreen("storage")
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.globalKeyMap.Screen3):
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
		m.currentScreen = updatedModel.(Screen)

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

	combinedKeys := combinedKeyMap{
		global: m.globalKeyMap,
		screen: m.currentScreen.Keybindings(),
		scroll: m.currentScreen.ScrollKeys(),
	}

	footer := footerStyle.Render(
		m.help.FullHelpView(combinedKeys.FullHelp()),
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

	availableScreenHeight := max(m.height-1-m.keyMapViewHeight, 0)

	var screenCmd tea.Cmd
	var updatedModel tea.Model

	updatedModel, screenCmd = m.currentScreen.Update(
		messages.ScreenSizeMsg{
			Width:  m.width,
			Height: availableScreenHeight,
		},
	)

	m.currentScreen = updatedModel.(Screen)

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
