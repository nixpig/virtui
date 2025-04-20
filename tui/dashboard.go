package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type dashboardModel struct {
	connections []qemuConnection
}

func initDashboard(connections []qemuConnection) dashboardModel {
	model := dashboardModel{
		connections: connections,
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

	for _, c := range m.connections {
		for _, d := range c.domains {
			v.WriteString(d.name + "\n")
		}
	}

	return v.String()
}
