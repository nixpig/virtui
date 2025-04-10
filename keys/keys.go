package keys

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	Up          key.Binding
	Down        key.Binding
	Open        key.Binding
	Run         key.Binding
	PauseResume key.Binding
	Shutdown    key.Binding
	Reboot      key.Binding
	ForceReset  key.Binding
	ForceOff    key.Binding
	SaveRestore key.Binding
	Migrate     key.Binding
	Delete      key.Binding
	Clone       key.Binding
	New         key.Binding
	Help        key.Binding
	Quit        key.Binding
	Focus       key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Open, k.Help, k.Quit}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.New, k.SaveRestore, k.Migrate, k.Clone, k.Delete},
		{k.Run, k.PauseResume, k.Shutdown, k.Reboot, k.ForceReset, k.ForceOff},
		{k.Up, k.Down, k.Open, k.Help, k.Quit},
	}
}

var Keys = Keymap{
	Up: key.NewBinding(
		key.WithKeys("up", "k", "K"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j", "J"),
		key.WithHelp("↓/j", "down"),
	),
	Open: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "open"),
	),
	Run: key.NewBinding(
		key.WithKeys("t", "T"),
		key.WithHelp("t", "run"),
	),
	PauseResume: key.NewBinding(
		key.WithKeys("p", "P"),
		key.WithHelp("p", "pause/resume"),
	),
	Shutdown: key.NewBinding(
		key.WithKeys("s", "S"),
		key.WithHelp("s", "shutdown"),
	),
	Reboot: key.NewBinding(
		key.WithKeys("r", "R"),
		key.WithHelp("r", "reboot"),
	),
	ForceReset: key.NewBinding(
		key.WithKeys("o", "O"),
		key.WithHelp("o", "force reset"),
	),
	ForceOff: key.NewBinding(
		key.WithKeys("f", "F"),
		key.WithHelp("f", "force off"),
	),
	SaveRestore: key.NewBinding(
		key.WithKeys("v", "V"),
		key.WithHelp("v", "save/restore"),
	),
	Migrate: key.NewBinding(
		key.WithKeys("m", "M"),
		key.WithHelp("m", "migrate"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d", "D"),
		key.WithHelp("d", "delete"),
	),
	Clone: key.NewBinding(
		key.WithKeys("c", "C"),
		key.WithHelp("c", "clone"),
	),
	Help: key.NewBinding(
		key.WithKeys("h", "H", "?"),
		key.WithHelp("h", "help"),
	),
	New: key.NewBinding(
		key.WithKeys("n", "N"),
		key.WithHelp("n", "new"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "Q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Focus: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "focus"),
	),
}
