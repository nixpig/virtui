package tui

import tea "github.com/charmbracelet/bubbletea"

// selectGuestMsg is a message to communicate the currently selected guest by UUID
type selectGuestMsg struct {
	selectedUUID string
}

type goBackMsg struct{}

func selectGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return selectGuestMsg{selectedUUID: uuid}
	}
}

func goBackCmd() tea.Cmd {
	return func() tea.Msg {
		return goBackMsg{}
	}
}
