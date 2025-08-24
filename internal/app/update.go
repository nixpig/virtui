package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"libvirt.org/go/libvirt"
)

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Debug(
		"AppModel.Update received message",
		"msg_type",
		fmt.Sprintf("%T", msg),
	)

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case *libvirt.DomainEventLifecycle:
		// TODO: handle domain event

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = m.width

		// calculate available height for the current screen
		// where header is 1 line, footer height is stored in m.keybindingsViewHeight
		availableScreenHeight := max(m.height-1-m.keybindingsViewHeight, 0)

		// pass dimensions to the current screen
		if m.currentScreen != nil {
			m.currentScreen.SetDimensions(m.width, availableScreenHeight) // Updated: Pass calculated height
			// also send WindowSizeMsg to the current screen's Update method
			// this allows the screen to react to resize events
			var screenCmd tea.Cmd
			var updatedModel tea.Model
			updatedModel, screenCmd = m.currentScreen.Update(msg)
			m.currentScreen = updatedModel.(Screen)
			cmds = append(cmds, screenCmd)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "?":
			m.showFullHelp = !m.showFullHelp

			// recalculate footer height and update screen dimensions
			// Create a temporary combined key map to get the new footer height
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
			m.keybindingsViewHeight = lipgloss.Height(tempFooter) // Update m.keybindingsViewHeight with the new height

			// Now recalculate availableScreenHeight with the updated footer height
			availableScreenHeight := max(m.height-1-m.keybindingsViewHeight, 0)
			m.currentScreen.SetDimensions(m.width, availableScreenHeight)

		default:
			// handle screen switching based on key press
			if screenName, ok := m.screenKeyMap[msg.String()]; ok {
				if nextScreen, ok := m.screens[screenName]; ok {
					m.currentScreen = nextScreen
					// when switching screens, recalculate and set dimensions for the new screen
					availableScreenHeight := max(m.height-1-m.keybindingsViewHeight, 0)
					m.currentScreen.SetDimensions(m.width, availableScreenHeight)
					// also send WindowSizeMsg to the new screen's Update method
					var screenCmd tea.Cmd
					updatedModel, screenCmd := m.currentScreen.Update(tea.WindowSizeMsg{
						Width:  m.width,
						Height: m.height,
					})
					m.currentScreen = updatedModel.(Screen)
					cmds = append(cmds, screenCmd)
				}
			}
		}
	}

	m.help, cmd = m.help.Update(msg)
	cmds = append(cmds, cmd)

	// delegate update to the current screen
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
