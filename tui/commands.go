package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"libvirt.org/go/libvirt"
)

type openGuestMsg struct{ uuid string }
type startGuestMsg struct{ uuid string }
type pauseResumeGuestMsg struct{ uuid string }
type shutdownGuestMsg struct{ uuid string }
type rebootGuestMsg struct{ uuid string }
type forceResetGuestMsg struct{ uuid string }
type forceOffGuestMsg struct{ uuid string }
type saveGuestMsg struct{ uuid string }
type cloneGuestMsg struct{ uuid string }
type deleteGuestMsg struct{ uuid string }

type goBackMsg struct{}

type registerGuestMsg struct {
	dom *libvirt.Domain
}

func registerGuestCmd(dom *libvirt.Domain) tea.Cmd {
	return func() tea.Msg {
		return registerGuestMsg{dom}
	}
}

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
