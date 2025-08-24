package common

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/key"
	"libvirt.org/go/libvirt"
)

type ErrMsg struct {
	Err error
}

func (e ErrMsg) Error() string {
	return e.Err.Error()
}

type Keymap struct {
	Manager key.Binding
	Network key.Binding
	Storage key.Binding

	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding

	Back key.Binding

	Quit key.Binding
	Help key.Binding
}

var Keys = Keymap{
	Manager: key.NewBinding(
		key.WithKeys("1", "f1"),
		key.WithHelp("1", "Guests"),
	),
	Network: key.NewBinding(
		key.WithKeys("2", "f2"),
		key.WithHelp("2", "Networks"),
	),
	Storage: key.NewBinding(
		key.WithKeys("3", "f3"),
		key.WithHelp("3", "Storage"),
	),

	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("󰁞/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("󰁆/j", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),	// TODO: need to fix this
		key.WithHelp("󰁎/h", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("󰁕/l", "right"),
	),

	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),

	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("", ""),
		key.WithHelp("", "https://github.com/nixpig/virtui"),
	),
}

func ListenForEvent(ch chan *libvirt.DomainEventLifecycle) tea.Cmd {
	return func() tea.Msg {
		e := <-ch
		return e
	}
}

type OpenGuestMsg struct{ UUID string }

func OpenGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return OpenGuestMsg{UUID: uuid}
	}
}

type StartGuestMsg struct{ UUID string }

func StartGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return StartGuestMsg{UUID: uuid}
	}
}

type PauseResumeGuestMsg struct{ UUID string }

func PauseResumeGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return PauseResumeGuestMsg{UUID: uuid}
	}
}

type ShutdownGuestMsg struct{ UUID string }

func ShutdownGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return ShutdownGuestMsg{UUID: uuid}
	}
}

type RebootGuestMsg struct{ UUID string }

func RebootGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return RebootGuestMsg{UUID: uuid}
	}
}

type ForceResetGuestMsg struct{ UUID string }

func ForceResetGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return ForceResetGuestMsg{UUID: uuid}
	}
}

type ForceOffGuestMsg struct{ UUID string }

func ForceOffGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return ForceOffGuestMsg{UUID: uuid}
	}
}

type SaveGuestMsg struct{ UUID string }

func SaveGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return SaveGuestMsg{UUID: uuid}
	}
}

type CloneGuestMsg struct{ UUID string }

func CloneGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return CloneGuestMsg{UUID: uuid}
	}
}

type DeleteGuestMsg struct{ UUID string }

func DeleteGuestCmd(uuid string) tea.Cmd {
	return func() tea.Msg {
		return DeleteGuestMsg{UUID: uuid}
	}
}

type GoBackMsg struct{}

func GoBackCmd() tea.Cmd {
	return func() tea.Msg {
		return GoBackMsg{}
	}
}
