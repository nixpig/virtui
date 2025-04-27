package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type GlobalMap struct {
	Help      key.Binding
	Quit      key.Binding
	Dashboard key.Binding
	Networks  key.Binding
	Storage   key.Binding
}

func (k GlobalMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k GlobalMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			Dashboard.Enter,
			Dashboard.Up,
			Dashboard.Down,
			Global.Help,
			Global.Quit,
		},
		{
			Dashboard.AddConnection,
			Dashboard.NewVM,
			Dashboard.Start,
			Dashboard.PauseResume,
			Dashboard.Shutdown,
			Dashboard.Reboot,
			Dashboard.Reset,
			Dashboard.PowerOff,
			Dashboard.SaveRestore,
			Dashboard.Delete,
		},
	}
}

var Global = GlobalMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Dashboard: key.NewBinding(
		key.WithKeys("1"),
	),
	Networks: key.NewBinding(
		key.WithKeys("2"),
	),
	Storage: key.NewBinding(
		key.WithKeys("3"),
	),
}
