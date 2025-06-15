package tui

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
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

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Left, k.Right, k.Back, k.Quit, k.Help}
}

func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

var Keys = keymap{
	Manager: key.NewBinding(
		key.WithKeys("1", "f1"),
		key.WithHelp("f1", "Manager"),
	),
	Network: key.NewBinding(
		key.WithKeys("2", "f2"),
		key.WithHelp("f2", "Network"),
	),
	Storage: key.NewBinding(
		key.WithKeys("3", "f3"),
		key.WithHelp("f3", "Storage"),
	),

	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("h", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("l", "right"),
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
