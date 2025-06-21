package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
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
	Start       key.Binding
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
		mk.Start,
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

func newManagerModel(conn *libvirt.Connect) tea.Model {
	domains, err := conn.ListAllDomains(0)
	if err != nil {
		// TODO: surface error to user?
		log.Debug("list all domains", "err", err)
	}

	rows := make([]table.Row, len(domains))

	for i, domain := range domains {
		d, err := entity.ToDomainStruct(&domain)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("convert entity to struct", "err", err, "domain", domain)
		}

		state, _, err := domain.GetState()
		if err != nil {
			log.Debug("get domain state", "uuid", d.UUID, "err", err)
		}

		if err := domain.Free(); err != nil {
			log.Warn("free ref counted domain struct", "err", err)
		}

		rows[i] = table.Row{
			d.UUID,
			d.Name,
			mappers.FromState(state),
			fmt.Sprintf("%d", d.VCPU.Value),
			// FIXME: assumes the d.Memory.Value is always the default KiB, which it's not...
			// https://libvirt.org/formatdomain.html#memory-allocation
			fmt.Sprintf("%dMiB", d.Memory.Value/1024),
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

	log.Debug("manager received msg", "type", fmt.Sprintf("%T", msg), "data", msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		// TODO: resize the table and stuff

	case tea.KeyMsg:
		guestUUID := m.table.SelectedRow()[0]

		switch {
		case key.Matches(msg, m.keys.Open):
			return m, openGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Start):
			return m, startGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.PauseResume):
			return m, pauseResumeGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Shutdown):
			return m, shutdownGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Reboot):
			return m, rebootGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.ForceReset):
			return m, forceResetGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.ForceOff):
			return m, forceOffGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Save):
			return m, saveGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Clone):
			return m, cloneGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Delete):
			return m, deleteGuestCmd(guestUUID)

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
