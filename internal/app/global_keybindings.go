package app

import (
	"github.com/charmbracelet/bubbles/key"
)

// GlobalKeyMap defines the keymaps available globally.
type GlobalKeyMap struct {
	Quit            key.Binding
	Help            key.Binding
	DashboardScreen key.Binding
	StorageScreen   key.Binding
	NetworksScreen  key.Binding

	// Domain Actions
	DomainStart       key.Binding
	DomainPauseResume key.Binding
	DomainShutdown    key.Binding
	DomainReboot      key.Binding
	DomainReset       key.Binding
	DomainForceOff    key.Binding
	DomainSave        key.Binding
	DomainClone       key.Binding
	DomainDelete      key.Binding
	DomainOpen        key.Binding
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

		// Domain Actions
		DomainStart: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "start"),
		),
		DomainPauseResume: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "pause/resume"),
		),
		DomainShutdown: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "shutdown"),
		),
		DomainReboot: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "reboot"),
		),
		DomainReset: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "reset"),
		),
		DomainForceOff: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "force off"),
		),
		DomainSave: key.NewBinding(
			key.WithKeys("v"),
			key.WithHelp("v", "save"),
		),
		DomainClone: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "clone"),
		),
		DomainDelete: key.NewBinding(
			key.WithKeys("x"),
			key.WithHelp("x", "delete"),
		),
		DomainOpen: key.NewBinding(
			key.WithKeys("o"),
			key.WithHelp("o", "open"),
		),
	}
}

// combinedKeyMap implements help.KeyMap for global and screen-specific keybindings.
type combinedKeyMap struct {
	global GlobalKeyMap
	screen   interface{}
	currentScreenID string
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

	if k.currentScreenID == "manager" {
		fullHelp = append(fullHelp, []key.Binding{
			k.global.DomainStart,
			k.global.DomainPauseResume,
			k.global.DomainShutdown,
			k.global.DomainReboot,
			k.global.DomainReset,
			k.global.DomainForceOff,
			k.global.DomainSave,
			k.global.DomainClone,
			k.global.DomainDelete,
			k.global.DomainOpen,
		})
	}

	if screenKeys, ok := k.screen.([][]key.Binding); ok {
		for _, row := range screenKeys {
			fullHelp = append(fullHelp, row)
		}
	}

	return fullHelp
}
