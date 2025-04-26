package keys

import "github.com/charmbracelet/bubbles/key"

type DashboardMap struct {
	Quit          key.Binding
	Help          key.Binding
	AddConnection key.Binding
	NewVM         key.Binding
	Up            key.Binding
	Down          key.Binding
	Enter         key.Binding
	Start         key.Binding
	PauseResume   key.Binding
	Shutdown      key.Binding
	Reboot        key.Binding
	Reset         key.Binding
	PowerOff      key.Binding
	SaveRestore   key.Binding
	Delete        key.Binding
}

func (k DashboardMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
		k.Enter,
		k.Help,
	}
}

func (k DashboardMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Enter,
			k.Up,
			k.Down,
			k.Help,
		},
		{
			k.AddConnection,
			k.NewVM,
		},
	}
}

var Dashboard = DashboardMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "Q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "more help"),
	),
	AddConnection: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add connection"),
	),
	NewVM: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "create vm"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("j", "down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("⏎", "open"),
	),
	Start: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "start"),
	),
	PauseResume: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "pause/resume"),
	),

	Shutdown: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "shutdown"),
	),
	Reboot: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reboot"),
	),
	Reset: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "reset"),
	),
	PowerOff: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "poweroff"),
	),
	SaveRestore: key.NewBinding(
		key.WithKeys("v"),
		key.WithHelp("v", "save/restore"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
}
