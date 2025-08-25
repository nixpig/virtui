package common

import "github.com/charmbracelet/bubbles/key"

// ScrollKeyMap defines the keybindings for scrolling.
type ScrollKeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
}

// DefaultScrollKeyMap returns the default keybindings for scrolling.
func DefaultScrollKeyMap() ScrollKeyMap {
	return ScrollKeyMap{
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
	}
}

func (k ScrollKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k ScrollKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Left, k.Down, k.Up, k.Right},
	}
}
