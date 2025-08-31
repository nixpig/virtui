package manager

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Start        key.Binding
	PauseResume  key.Binding
	Shutdown     key.Binding
	Reboot       key.Binding
	Reset        key.Binding
	ForceOff     key.Binding
	Save         key.Binding
	Clone        key.Binding
	Delete       key.Binding
	Open         key.Binding
	KeyUp        key.Binding
	KeyDown      key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.KeyUp, k.KeyDown, k.Start, k.PauseResume, k.Shutdown, k.Reboot, k.Reset, k.ForceOff, k.Save, k.Clone, k.Delete, k.Open}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.KeyUp, k.KeyDown},
	}
}

func (k KeyMap) Up() key.Binding {
	return k.KeyUp
}

func (k KeyMap) Down() key.Binding {
	return k.KeyDown
}

var defaultKeyMap = KeyMap{
	Start: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "start"),
	),
	PauseResume: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "pause/resume"),
	),
	Shutdown: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "shutdown"),
	),
	Reboot: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reboot"),
	),
	Reset: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "reset"),
	),
	ForceOff: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "force off"),
	),
	Save: key.NewBinding(
		key.WithKeys("v"),
		key.WithHelp("v", "save"),
	),
	Clone: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "clone"),
	),
	Delete: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "delete"),
	),
	Open: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open"),
	),
	KeyUp: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	KeyDown: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
}