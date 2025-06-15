package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/tui/entity"
	"github.com/nixpig/virtui/tui/mappers"
	"libvirt.org/go/libvirt"
)

var columns = []table.Column{
	// UUID is hidden as it's only used for identification
	{Title: "UUID", Width: 0},
	{Title: "Name", Width: 30},
	{Title: "State", Width: 10},
	{Title: "CPU", Width: 4},
	{Title: "Mem", Width: 12},
	{Title: "Connection", Width: 20},
}

type managerModel struct {
	domains []libvirt.Domain
	conn    *libvirt.Connect
	keys    managerKeyMap
	help    help.Model
	table   table.Model
}

type managerKeyMap struct {
	Open        key.Binding
	Run         key.Binding
	PauseResume key.Binding
	Shutdown    key.Binding
	Reboot      key.Binding
	ForceReset  key.Binding
	ForceOff    key.Binding
	Save        key.Binding
	Clone       key.Binding
	Delete      key.Binding
}

func (mk managerKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		mk.Open,
		mk.Run,
		mk.PauseResume,
		mk.Shutdown,
		mk.Reboot,
		mk.ForceReset,
		mk.ForceOff,
		mk.Save,
		mk.Clone,
		mk.Delete,
	}
}

func (mk managerKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

var managerKeys = managerKeyMap{
	Open: key.NewBinding(
		key.WithKeys("o", "enter"),
		key.WithHelp("o", "open"),
	),
	Run: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "run"),
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
	ForceReset: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "reset"),
	),
	ForceOff: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "off"),
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
}

// New creates a tea.Model for the manager view
func newManagerModel(conn *libvirt.Connect) tea.Model {
	domains, _ := conn.ListAllDomains(0)

	rows := make([]table.Row, len(domains))

	for i, d := range domains {
		x, _ := entity.ToDomainStruct(&d)
		state, _, _ := d.GetState()

		rows[i] = table.Row{
			x.UUID,
			x.Name,
			mappers.FromState(state),
			fmt.Sprintf("%d", x.VCPU.Value),
			fmt.Sprintf("%d%s", x.Memory.Value, x.Memory.Unit),
			"QEMU/KVM (system)",
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithRows(rows),
	)

	return managerModel{
		domains: domains,
		table:   t,
		keys:    managerKeys,
		help:    help.New(),
		conn:    conn,
	}
}

func (m managerModel) Init() tea.Cmd {
	return nil
}

func (m managerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		// TODO: resize the table and stuff

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Open):
			return m, selectGuestCmd(m.table.SelectedRow()[0])
		case key.Matches(msg, m.keys.Run):
			fmt.Println("START!!!")
		}
	}

	m.table, cmd = m.table.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m managerModel) View() string {
	helpView := m.help.View(m.keys)
	return m.table.View() + "\n" + helpView
}
