package tui

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	manager key.Binding
	network key.Binding
	storage key.Binding

	up    key.Binding
	down  key.Binding
	left  key.Binding
	right key.Binding

	back key.Binding

	quit key.Binding
	help key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.up, k.down, k.left, k.right, k.back, k.quit, k.help}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.manager,
			k.network,
			k.storage,
			k.quit,
		},
		{
			k.left,
			k.down,
			k.up,
			k.right,
		},
	}
}

var keys = Keymap{
	manager: key.NewBinding(
		key.WithKeys("1", "f1"),
		key.WithHelp("1", "Guests"),
	),
	network: key.NewBinding(
		key.WithKeys("2", "f2"),
		key.WithHelp("2", "Networks"),
	),
	storage: key.NewBinding(
		key.WithKeys("3", "f3"),
		key.WithHelp("3", "Storage"),
	),

	up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("󰁞/k", "up"),
	),
	down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("󰁆/j", "down"),
	),
	left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("󰁎/h", "left"),
	),
	right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("󰁕/l", "right"),
	),

	back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),

	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	help: key.NewBinding(
		key.WithKeys("", ""),
		key.WithHelp("", "https://github.com/nixpig/virtui"),
	),
}
