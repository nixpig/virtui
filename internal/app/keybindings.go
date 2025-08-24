package app

import "github.com/charmbracelet/bubbles/key"

var globalKeyBindings = []key.Binding{
	key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "Manager Screen"),
	),
	key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "Storage Screen"),
	),
	key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "Network Screen"),
	),
	key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}
