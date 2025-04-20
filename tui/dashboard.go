package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/digitalocean/go-libvirt"
)

type dashboardModel struct {
	connections []*libvirt.Libvirt
	domains     []libvirt.Domain
}

func initDashboard(connections []*libvirt.Libvirt) dashboardModel {
	model := dashboardModel{
		connections: connections,
	}

	for _, c := range model.connections {
		d, _, err := c.ConnectListAllDomains(1, 0)
		if err != nil {
			continue
		}

		model.domains = append(model.domains, d...)
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

	for i, d := range m.domains {
		v.WriteString(fmt.Sprintf("%d - %s\n", i, d.Name))
	}

	return v.String()
}
