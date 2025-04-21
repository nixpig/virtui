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
	Up        key.Binding
	Down      key.Binding
}

func (k GlobalMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k GlobalMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
	}
}

var Global = GlobalMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "Q", "ctrl+c"),
		key.WithHelp("q", "quit"),
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
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
	),
}
