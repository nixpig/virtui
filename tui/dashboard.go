package tui

import (
	"net/url"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/term"
	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
	"github.com/nixpig/virtui/commander"
	"github.com/nixpig/virtui/keys"
	"github.com/nixpig/virtui/vm/domain"
)

var baseStyle = lipgloss.NewStyle()

type dashboardData map[string]map[libvirt.UUID]libvirt.Domain

type dashboardDomain struct {
	lvdom  *libvirt.Domain
	conn   *libvirt.Libvirt
	name   string
	uuid   string
	id     int
	status string
	host   string
	cpu    string
	mem    string
	net    string
	disk   string
}

type dashboardModel struct {
	connections map[string]*libvirt.Libvirt
	data        []dashboardDomain
	table       table.Model
	keys        keys.DashboardMap
}

func initDashboard(connections map[string]*libvirt.Libvirt) dashboardModel {
	model := dashboardModel{
		connections: connections,
		keys:        keys.Dashboard,
	}

	// ---

	var data []dashboardDomain

	for k, c := range model.connections {
		domains, _, err := c.ConnectListAllDomains(1, 0)
		if err != nil {
			continue
		}

		u, err := url.Parse(k)
		if err != nil {
			log.Error("failed to parse connection uri", "err", err)
			continue
		}

		for _, d := range domains {
			state, _, _ := c.DomainGetState(d, 0)
			uuid, _ := uuid.FromBytes(d.UUID[:])

			data = append(data, dashboardDomain{
				lvdom:  &d,
				conn:   c,
				name:   d.Name,
				uuid:   uuid.String(),
				id:     int(d.ID),
				status: domain.PresentableState(libvirt.DomainState(state)),
				host:   u.Host + u.Path,
			})
		}
	}

	model.data = data

	// ---

	w, _, _ := term.GetSize(0)

	hostW := 10
	stateW := 10
	cpuW := 5
	memW := 5
	diskW := 5
	netW := 5

	baselineW := 8 + 8

	nameW := w - baselineW - hostW - stateW - cpuW - memW - diskW - netW

	columns := []table.Column{
		{Title: "UUID", Width: 0},
		{Title: "Host", Width: hostW},
		{Title: "State", Width: stateW},
		{Title: "Name", Width: nameW},
		{Title: "CPU", Width: cpuW},
		{Title: "Mem", Width: memW},
		{Title: "Disk", Width: diskW},
		{Title: "Net", Width: netW},
	}

	var rows []table.Row

	s := table.DefaultStyles()

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	for _, d := range model.data {
		rows = append(rows, table.Row{
			d.uuid,
			d.host,
			d.status,
			d.name,
			"⣾⣷⣷⣷⣷",
			"⣷⣾⣤⣷⣤",
			"▄▄ ▆▆",
			"▇▇ ▃▃",
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	t.SetStyles(s)

	model.table = t

	return model
}

func (m dashboardModel) Init() tea.Cmd {
	return nil
}

func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	d := m.data[m.table.Cursor()]
	s, _, _ := d.conn.DomainGetState(*d.lvdom, 0)
	state := libvirt.DomainState(s)
	c := commander.NewCommander(d.conn)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Enter):
			return m, func() tea.Msg {
				selectedRow := m.table.SelectedRow()
				uuid := selectedRow[0]
				log.Debug("select domain", "uuid", uuid)
				return selectDomainMsg{uuid}
			}

		case key.Matches(msg, m.keys.AddConnection):
			log.Debug("add connection")

		case key.Matches(msg, m.keys.NewVM):
			log.Debug("new vm")

		case key.Matches(msg, m.keys.Start):
			if err := c.StartDomain(d.lvdom); err != nil {
				log.Error("start domain", "err", err)
			}

		case key.Matches(msg, m.keys.PauseResume):
			switch state {
			case libvirt.DomainRunning:
				if err := c.PauseDomain(d.lvdom); err != nil {
					log.Error("pause domain", "err", err)
				}
			case libvirt.DomainPaused:
				if err := c.ResumeDomain(d.lvdom); err != nil {
					log.Error("resume domain", "err", err)
				}
			default:
				// TODO: noop
			}

		case key.Matches(msg, m.keys.Shutdown):
			if state == libvirt.DomainRunning {
				if err := c.ShutdownDomain(d.lvdom); err != nil {
					log.Error("shutdown domain", "err", err)
				}
			} else {
				// TODO: noop
			}

		case key.Matches(msg, m.keys.Reboot):
			if state == libvirt.DomainRunning {
				if err := c.RebootDomain(d.lvdom); err != nil {
					log.Error("reboot domain", "err", err)
				}
			} else {
				// TODO: noop
			}

		case key.Matches(msg, m.keys.Reset):
			if err := c.ResetDomain(d.lvdom); err != nil {
				log.Error("reset domain", "err", err)
			}

		case key.Matches(msg, m.keys.PowerOff):
			if err := c.PoweroffDomain(d.lvdom); err != nil {
				log.Error("poweroff domain", "err", err)
			}

		case key.Matches(msg, m.keys.SaveRestore):
			log.Debug("save/restore vm")

		case key.Matches(msg, m.keys.Delete):
			if err := c.DeleteDomain(d.lvdom); err != nil {
				log.Error("delete domain", "err", err)
			}

		}
	}

	m.table, cmd = m.table.Update(msg)

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m dashboardModel) View() string {
	var v strings.Builder

	v.WriteString(baseStyle.Render(m.table.View()))

	return v.String()
}
