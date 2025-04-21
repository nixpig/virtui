package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/digitalocean/go-libvirt"
	"github.com/nixpig/virtui/vm/domain"
)

type dashboardData map[string]map[libvirt.UUID]libvirt.Domain

type dashboardModel struct {
	connections map[string]*libvirt.Libvirt
	data        dashboardData
}

func initDashboard(connections map[string]*libvirt.Libvirt) dashboardModel {
	model := dashboardModel{
		connections: connections,
		data:        make(dashboardData),
	}

	for k, c := range model.connections {
		m := make(map[libvirt.UUID]libvirt.Domain)

		domains, _, err := c.ConnectListAllDomains(1, 0)
		if err != nil {
			continue
		}

		for _, d := range domains {
			m[d.UUID] = d
		}

		model.data[k] = m
	}

	return model
}

func (m dashboardModel) Init() tea.Cmd {
	return nil
}

func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m dashboardModel) View() string {
	var v strings.Builder

	for k, c := range m.data {
		v.WriteString(fmt.Sprintf("Connection: %s\n", k))

		for _, d := range c {
			s, _, _ := m.connections[k].DomainGetState(d, 0)
			v.WriteString(fmt.Sprintf("- %s - %s\n", d.Name, domain.PresentableState(libvirt.DomainState(s))))
		}
	}

	return v.String()
}
