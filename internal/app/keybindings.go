package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/nixpig/virtui/internal/common"
)

type GlobalKeyMap struct {
	Quit    key.Binding
	Help    key.Binding
	Screen1 key.Binding
	Screen2 key.Binding
	Screen3 key.Binding
	Keys    []key.Binding
}

// DefaultGlobalKeyMap returns the default keybindings for actions that
// are available across all screens.
func DefaultGlobalKeyMap() GlobalKeyMap {
	keyMap := GlobalKeyMap{
		Screen1: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", "Dashboard"),
		),
		Screen2: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", "Storage"),
		),
		Screen3: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("3", "Networks"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}

	keyMap.Keys = []key.Binding{
		keyMap.Screen1,
		keyMap.Screen2,
		keyMap.Screen3,
		keyMap.Help,
		keyMap.Quit,
	}

	return keyMap
}

// combinedKeyMap implements help.KeyMap for global and screen-specific keybindings.
type combinedKeyMap struct {
	global GlobalKeyMap
	screen []key.Binding
	scroll common.ScrollKeyMap
}

func (k combinedKeyMap) ShortHelp() []key.Binding {
	return k.global.Keys
}

func (k combinedKeyMap) FullHelp() [][]key.Binding {
	fullHelp := append([][]key.Binding{}, k.global.Keys)
	fullHelp = append(fullHelp, k.scroll.FullHelp()...)
	fullHelp = append(fullHelp, k.screen)

	return fullHelp
}
