package keys

import "github.com/charmbracelet/bubbles/key"

type global struct {
	Help key.Binding
	Quit key.Binding
}

func (k global) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k global) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
	}
}

var Global = global{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "Q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
