package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	libvirtconn "github.com/nixpig/virtui/internal/libvirt"
	"github.com/nixpig/virtui/internal/messages" // New import
)

func (m *model) switchScreen(screenName string) tea.Cmd {
	if nextScreen, ok := m.screens[screenName]; ok {
		m.currentScreen = nextScreen
		availableScreenHeight := max(m.height-1-m.keybindingsViewHeight, 0)
		var screenCmd tea.Cmd
		var updatedModel tea.Model
		updatedModel, screenCmd = m.currentScreen.Update(ScreenSizeMsg{Width: m.width, Height: availableScreenHeight})
		m.currentScreen = updatedModel.(Screen)
		return screenCmd
	}
	return nil
}

func getDomainsCmd(service libvirtconn.Service) tea.Cmd {
	return func() tea.Msg {
		domains, err := service.ListAllDomains()
		if err != nil {
			log.Error("failed to list all domains", "err", err)
			return nil // Or return an error message
		}
		return messages.DomainsMsg(domains)
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Debug(
		"AppModel.Update received message",
		"msg_type",
		fmt.Sprintf("%T", msg),
	)

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case libvirtconn.DomainEvent: // Changed from *libvirt.DomainEventLifecycle
		// TODO: handle domain event

	case messages.DomainsMsg: // New case for DomainsMsg
		if m.currentScreen.ID() == "manager" {
			var screenCmd tea.Cmd
			var updatedModel tea.Model
			updatedModel, screenCmd = m.currentScreen.Update(msg)
			m.currentScreen = updatedModel.(Screen)
			cmds = append(cmds, screenCmd)
		}

	case ScreenSizeMsg:
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

		availableScreenHeight := max(m.height-1-m.keybindingsViewHeight, 0)

		if m.currentScreen != nil {
			var screenCmd tea.Cmd
			var updatedModel tea.Model
			updatedModel, screenCmd = m.currentScreen.Update(ScreenSizeMsg{Width: m.width, Height: availableScreenHeight})
			m.currentScreen = updatedModel.(Screen)
			cmds = append(cmds, screenCmd)
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.globalKeybindings.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.globalKeybindings.Help):
			m.showFullHelp = !m.showFullHelp

			combinedKeys := combinedKeyMap{
				global: m.globalKeybindings,
				screen: m.currentScreen.Keybindings(),
				scroll: m.currentScreen.ScrollKeys(),
			}

			var tempFooter string
			if m.showFullHelp {
				tempFooter = m.help.FullHelpView(combinedKeys.FullHelp())
			} else {
				tempFooter = m.help.ShortHelpView(combinedKeys.ShortHelp())
			}
			m.keybindingsViewHeight = lipgloss.Height(tempFooter)

			availableScreenHeight := max(m.height-1-m.keybindingsViewHeight, 0)
			if m.currentScreen != nil {
				var screenCmd tea.Cmd
				var updatedModel tea.Model
				updatedModel, screenCmd = m.currentScreen.Update(ScreenSizeMsg{Width: m.width, Height: availableScreenHeight})
				m.currentScreen = updatedModel.(Screen)
				cmds = append(cmds, screenCmd)
			}

		case key.Matches(msg, m.globalKeybindings.Screen1):
			cmd := m.switchScreen("manager")
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.globalKeybindings.Screen2):
			cmd := m.switchScreen("storage")
			cmds = append(cmds, cmd)

		case key.Matches(msg, m.globalKeybindings.Screen3):
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
