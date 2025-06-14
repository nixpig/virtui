package keys

import "github.com/charmbracelet/bubbles/key"

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

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Left, k.Right, k.Back, k.Quit, k.Help}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
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
