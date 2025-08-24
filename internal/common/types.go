package common

import "github.com/charmbracelet/bubbles/key"

// ScrollKeyMap defines the keybindings for scrolling.
type ScrollKeyMap struct {
	ScrollUp   key.Binding
	ScrollDown key.Binding
}

func (k ScrollKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k ScrollKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.ScrollUp, k.ScrollDown},
	}
}

// DefaultScrollKeyMap returns the default keybindings for scrolling.
func DefaultScrollKeyMap() ScrollKeyMap {
	return ScrollKeyMap{
		ScrollUp: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k", "scroll up"),
		),
		ScrollDown: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j", "scroll down"),
		),
	}
}
