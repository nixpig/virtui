package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/nixpig/virtui/internal/common"
)

type KeyMap struct {
	Quit    key.Binding
	Help    key.Binding
	Screen1 key.Binding
	Screen2 key.Binding
	Screen3 key.Binding
	Keys    []key.Binding
}

var GlobalKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Screen1: key.NewBinding(
		key.WithKeys("1"),
		key.WithHelp("1", "Manager Screen"),
	),
	Screen2: key.NewBinding(
		key.WithKeys("2"),
		key.WithHelp("2", "Storage Screen"),
	),
	Screen3: key.NewBinding(
		key.WithKeys("3"),
		key.WithHelp("3", "Network Screen"),
	),
}

func init() {
	GlobalKeyMap.Keys = []key.Binding{
		GlobalKeyMap.Quit,
		GlobalKeyMap.Help,
		GlobalKeyMap.Screen1,
		GlobalKeyMap.Screen2,
		GlobalKeyMap.Screen3,
	}
}

// combinedKeyMap implements help.KeyMap for global and screen-specific keybindings.
type combinedKeyMap struct {
	global KeyMap
	screen []key.Binding
	scroll common.ScrollKeyMap
}

func (k combinedKeyMap) ShortHelp() []key.Binding {
	return k.global.Keys
}

func (k combinedKeyMap) FullHelp() [][]key.Binding {
	fullHelp := [][]key.Binding{}
	fullHelp = append(fullHelp, k.global.Keys)
	fullHelp = append(fullHelp, k.screen)
	fullHelp = append(
		fullHelp,
		k.scroll.FullHelp()...) // Assuming ScrollKeyMap has FullHelp
	return fullHelp
}
