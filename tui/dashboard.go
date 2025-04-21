package tui

import (
	"net/url"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/digitalocean/go-libvirt"
	"github.com/nixpig/virtui/vm/domain"
)

var baseStyle = lipgloss.NewStyle()

type dashboardData map[string]map[libvirt.UUID]libvirt.Domain

type dashboardModel struct {
	connections map[string]*libvirt.Libvirt
	table       table.Model
	width       int
	// keys        keys.GlobalMap
}

func initDashboard(connections map[string]*libvirt.Libvirt) dashboardModel {
	model := dashboardModel{
		connections: connections,
		// keys:        keys.Global,
	}

	s := table.DefaultStyles()

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	columns := []table.Column{
		{Title: "Host", Width: 10},
		{Title: "State", Width: 10},
		{Title: "Name", Width: 30},
		{Title: "CPU", Width: 5},
		{Title: "Mem", Width: 5},
		{Title: "Disk", Width: 9},
		{Title: "Net", Width: 9},
	}

	var rows []table.Row

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
			rows = append(rows, table.Row{
				u.Host + u.Path,
				domain.PresentableState(libvirt.DomainState(state)),
				d.Name,
				"100%",
				"100%",
				"↓99 ↑999",
				"↓999 ↑99",
			})
		}

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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.table.SetWidth(msg.Width)
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
