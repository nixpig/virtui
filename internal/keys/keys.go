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

	Help key.Binding
	Quit key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Manager, k.Network, k.Storage},
		{k.Up, k.Down, k.Left, k.Right},
		{k.Back, k.Help, k.Quit},
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

	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k/↑", "Up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j/↓", "Down"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("h/←", "Left"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("l/→", "Right"),
	),

	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Back"),
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
