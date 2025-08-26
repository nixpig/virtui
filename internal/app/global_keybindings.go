package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/nixpig/virtui/internal/common"
)

// GlobalKeyMap defines the keymaps available globally.
type GlobalKeyMap struct {
	Quit            key.Binding
	Help            key.Binding
	DashboardScreen key.Binding
	StorageScreen   key.Binding
	NetworksScreen  key.Binding
}

// DefaultGlobalKeyMap returns the default keybindings for actions that
// are available across all screens.
func DefaultGlobalKeyMap() GlobalKeyMap {
	return GlobalKeyMap{
		DashboardScreen: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", "Dashboard"),
		),
		StorageScreen: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", "Storage"),
		),
		NetworksScreen: key.NewBinding(
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
}

// combinedKeyMap implements help.KeyMap for global and screen-specific keybindings.
type combinedKeyMap struct {
	global GlobalKeyMap
	screen []key.Binding
	scroll common.ScrollKeyMap
}

func (k combinedKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k combinedKeyMap) FullHelp() [][]key.Binding {
	fullHelp := [][]key.Binding{
		{
			k.global.DashboardScreen,
			k.global.StorageScreen,
			k.global.NetworksScreen,
			k.global.Help,
			k.global.Quit,
		},
	}

	fullHelp = append(fullHelp, k.scroll.FullHelp()...)
	fullHelp = append(fullHelp, k.screen)

	return fullHelp
}
