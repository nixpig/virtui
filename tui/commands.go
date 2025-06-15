package tui

import tea "github.com/charmbracelet/bubbletea"

// SelectGuestMsg is a message to communicate the currently selected guest by UUID
type SelectGuestMsg struct {
	SelectedUUID string
}

type GoBackMsg struct{}

func SelectGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return SelectGuestMsg{SelectedUUID: uuid}
	}
}

func GoBackCmd() tea.Cmd {
	return func() tea.Msg {
		return GoBackMsg{}
	}
}
