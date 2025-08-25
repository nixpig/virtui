package app

import (
	"github.com/charmbracelet/lipgloss"
)

func (m *model) View() string {
	if m.currentScreen == nil {
		return "Loading..."
	}

	// --- Header (Title) ---
	titleStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#007bff")). // Blue background
		Foreground(lipgloss.Color("#ffffff")). // White text
		Padding(0, 1).
		Width(m.width).
		Align(lipgloss.Center)

	header := titleStyle.Render(m.currentScreen.Title())

	// --- Footer (Keybindings) ---
	// Create a combined key map for the help view
	combinedKeys := combinedKeyMap{
		global: m.globalKeybindings,
		screen: m.currentScreen.Keybindings(),
		scroll: m.currentScreen.ScrollKeys(),
	}

	var footer string
	if m.showFullHelp {
		footer = m.help.FullHelpView(combinedKeys.FullHelp())
	} else {
		footer = m.help.ShortHelpView(combinedKeys.ShortHelp())
	}

	// store the height of the keybindings view for content height calculation
	m.keybindingsViewHeight = lipgloss.Height(footer)

	// --- Main Content ---
	availableHeight := max(
		m.height-lipgloss.Height(header)-m.keybindingsViewHeight,
		0,
	)

	contentStyle := lipgloss.NewStyle().
		Height(availableHeight).
		Width(m.width)

	mainContent := contentStyle.Render(m.currentScreen.View())

	return lipgloss.JoinVertical(lipgloss.Left, header, mainContent, footer)
}
