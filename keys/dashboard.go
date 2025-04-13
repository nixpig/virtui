package keys

import "github.com/charmbracelet/bubbles/key"

type DashboardMap struct {
	Help          key.Binding
	AddConnection key.Binding
	NewVM         key.Binding
}

func (k DashboardMap) ShortHelp() []key.Binding {
	return []key.Binding{k.AddConnection, k.NewVM}
}

func (k DashboardMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.AddConnection},
		{k.NewVM},
	}
}

var Dashboard = DashboardMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	AddConnection: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add connection"),
	),
	NewVM: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new vm"),
	),
}
