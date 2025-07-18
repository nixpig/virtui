package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type guestMsg struct{ uuid string }

type openGuestMsg guestMsg
type startGuestMsg guestMsg
type pauseResumeGuestMsg guestMsg
type shutdownGuestMsg guestMsg
type rebootGuestMsg guestMsg
type forceResetGuestMsg guestMsg
type forceOffGuestMsg guestMsg
type saveGuestMsg guestMsg
type cloneGuestMsg guestMsg
type deleteGuestMsg guestMsg

type goBackMsg struct{}

func openGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return openGuestMsg{uuid}
	}
}

func startGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return startGuestMsg{uuid}
	}
}

func pauseResumeGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return pauseResumeGuestMsg{uuid}
	}
}

func shutdownGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return shutdownGuestMsg{uuid}
	}
}

func rebootGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return rebootGuestMsg{uuid}
	}
}

func forceResetGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return forceResetGuestMsg{uuid}
	}
}

func forceOffGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return forceOffGuestMsg{uuid}
	}
}

func saveGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return saveGuestMsg{uuid}
	}
}

func cloneGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return cloneGuestMsg{uuid}
	}
}

func deleteGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return deleteGuestMsg{uuid}
	}
}

func goBackCmd() tea.Cmd {
	return func() tea.Msg {
		return goBackMsg{}
	}
}
