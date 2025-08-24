package manager

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/mappers"
	"github.com/nixpig/virtui/internal/service"
	"github.com/nixpig/virtui/internal/tui/common"
	"github.com/nixpig/virtui/internal/tui/icons"
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

type Model interface {
	tea.Model
	Help() *help.Model
	Keys() *managerKeyMap
}

type managerModel struct {
	keys          managerKeyMap
	help          help.Model
	table         table.Model
	service       service.Service
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

func NewModel(svc service.Service) tea.Model {
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
		table:   t,
		keys:    managerKeys,
		help:    help.New(),
		service: svc,
	}

	return m
}

func (m managerModel) Help() *help.Model {
	return &m.help
}

func (m managerModel) Keys() *managerKeyMap {
	return &m.keys
}

func (m *managerModel) Init() tea.Cmd {
	return nil
}

func (m *managerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case *libvirt.DomainEventLifecycle:
		domains, err := m.service.ListAllDomains()
		if err != nil {
			log.Debug("list all domains", "err", err)
			return m, func() tea.Msg {
				return common.ErrMsg{Err: err}
			}
		}

		rows := make([]table.Row, len(domains))

		for i, d := range domains {
			var icon string

			switch d.State {
			case libvirt.DOMAIN_RUNNING:
				icon = icons.Vm.Running
			case libvirt.DOMAIN_BLOCKED:
				icon = icons.Vm.Blocked
			case libvirt.DOMAIN_PAUSED:
				icon = icons.Vm.Paused
			case libvirt.DOMAIN_SHUTDOWN, libvirt.DOMAIN_SHUTOFF:
				icon = icons.Vm.Off
			default:
				icon = icons.Vm.Off
			}

			rows[i] = table.Row{
				d.Domain.UUID,
				fmt.Sprintf(" %s  %s", icon, d.Domain.Name),
				mappers.FromState(d.State),
				fmt.Sprintf("%d", d.Domain.VCPU.Value),
				// FIXME: assumes the d.Memory.Value is always the default KiB, which it's not...
				// https://libvirt.org/formatdomain.html#memory-allocation
				fmt.Sprintf("%dMiB", d.Domain.Memory.Value/1024),
			}
		}

		m.table.SetRows(rows)

	case tea.WindowSizeMsg:
		nameWidth := msg.Width - 32
		m.table.Columns()[1].Width = nameWidth
		m.table.SetHeight(10)
		m.table.SetWidth(msg.Width)

	case tea.KeyMsg:
		log.Debug("KeyMsg received in managerModel.Update", "key", msg.String())
		var guestUUID string
		if len(m.table.Rows()) > 0 {
			guestUUID = m.table.SelectedRow()[0]
			log.Debug("Selected guest UUID", "uuid", guestUUID)
		} else {
			log.Debug("No rows in table, guestUUID is empty")
		}

		switch {
		case key.Matches(msg, m.keys.Open):
			log.Debug("Key matched: Open")
			return m, common.OpenGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Start):
			log.Debug("Key matched: Start")
			return m, common.StartGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.PauseResume):
			log.Debug("Key matched: PauseResume")
			return m, common.PauseResumeGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Shutdown):
			log.Debug("Key matched: Shutdown")
			return m, common.ShutdownGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Reboot):
			log.Debug("Key matched: Reboot")
			return m, common.RebootGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.ForceReset):
			log.Debug("Key matched: ForceReset")
			return m, common.ForceResetGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.ForceOff):
			log.Debug("Key matched: ForceOff")
			return m, common.ForceOffGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Save):
			log.Debug("Key matched: Save")
			return m, common.SaveGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Clone):
			log.Debug("Key matched: Clone")
			return m, common.CloneGuestCmd(guestUUID)

		case key.Matches(msg, m.keys.Delete):
			log.Debug("Key matched: Delete")
			return m, common.DeleteGuestCmd(guestUUID)

		}
	}

	m.table, cmd = m.table.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *managerModel) View() string {
	return m.table.View()
}
