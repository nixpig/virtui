package keys

import "github.com/charmbracelet/bubbles/key"

type DashboardMap struct {
	Help          key.Binding
	AddConnection key.Binding
	NewVM         key.Binding
	Up            key.Binding
	Down          key.Binding
	Enter         key.Binding
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
		key.WithHelp("n", "new vm"),
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
}
