package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/entity"
	"github.com/nixpig/virtui/internal/mappers"
	"libvirt.org/go/libvirt"
)

var columns = []table.Column{
	// UUID is hidden as it's only used for identification
	{Title: "UUID", Width: 0},
	{Title: " Name", Width: 30},
	{Title: "State", Width: 12},
	{Title: "CPU", Width: 6},
	{Title: "Mem", Width: 12},
}

type managerModel struct {
	keys          managerKeyMap
	help          help.Model
	table         table.Model
	conn          *libvirt.Connect
	width, height int
}

type managerKeyMap struct {
	New         key.Binding
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
		mk.New,
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
	return [][]key.Binding{
		{
			mk.Open,
			mk.Start,
			mk.Reboot,
			mk.Shutdown,
		},
		{
			mk.PauseResume,
			mk.ForceReset,
			mk.ForceOff,
			mk.Delete,
		},
		{

			mk.New,
			// mk.Save,
			// mk.Clone,
		},
	}
}

var managerKeys = managerKeyMap{
	New: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new"),
	),

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
		key.WithHelp("p", "pause"),
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
	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithStyles(table.Styles{
			Header: lipgloss.NewStyle().
				Bold(true).
				Border(lipgloss.NormalBorder(), false, false, true),
			Selected: lipgloss.NewStyle().
				Background(lipgloss.Color("2")).
				Foreground(lipgloss.Color("0")),
		}),
	)

	m := &managerModel{
		table: t,
		keys:  managerKeys,
		help:  help.New(),
		conn:  conn,
	}

	return m
}

func (m managerModel) Init() tea.Cmd {
	return nil
}

func (m managerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case *libvirt.DomainEventLifecycle:
		domains, err := m.conn.ListAllDomains(0)
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
				// TODO: surface error to user?
				log.Debug("get domain state", "uuid", d.UUID, "err", err)
			}
			// if err := domain.Free(); err != nil {
			// 	log.Warn("free ref counted domain struct", "err", err)
			// }

			var icon string

			switch state {
			case libvirt.DOMAIN_RUNNING:
				icon = fonts.vm.play
			case libvirt.DOMAIN_BLOCKED:
				icon = fonts.vm.off
			case libvirt.DOMAIN_PAUSED:
				icon = fonts.vm.pause
			case libvirt.DOMAIN_SHUTDOWN, libvirt.DOMAIN_SHUTOFF:
				icon = fonts.vm.empty
			default:
				icon = fonts.vm.empty
			}

			rows[i] = table.Row{
				d.UUID,
				fmt.Sprintf(" %s  %s", icon, d.Name),
				mappers.FromState(state),
				fmt.Sprintf("%d", d.VCPU.Value),
				// FIXME: assumes the d.Memory.Value is always the default KiB, which it's not...
				// https://libvirt.org/formatdomain.html#memory-allocation
				fmt.Sprintf("%dMiB", d.Memory.Value/1024),
			}
		}

		m.table.SetRows(rows)

	case tea.WindowSizeMsg:
		log.Debug("window size msg in manager", "width", msg.Width, "height", msg.Height)
		m.width = msg.Width
		m.table.SetHeight(10)
		nameWidth := m.width - 32
		m.table.Columns()[1].Width = nameWidth
		// TODO: resize the table and stuff

	case tea.KeyMsg:
		var guestUUID string
		if len(m.table.Rows()) > 0 {
			guestUUID = m.table.SelectedRow()[0]
		}

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

	return m.table.View()
}
