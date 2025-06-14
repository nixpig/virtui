package keys

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	Manager key.Binding
	Network key.Binding
	Storage key.Binding

	Help key.Binding
	Quit key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Manager, k.Network, k.Storage},
		{k.Help, k.Quit},
	}
}

var Keys = Keymap{
	Manager: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "Manager"),
	),
	Network: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "Network"),
	),
	Storage: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "Storage"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "Quit"),
	),
}
