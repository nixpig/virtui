package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type networkModel struct {
	connections []qemuConnection
}

func initNetwork(connections []qemuConnection) networkModel {
	model := networkModel{
		connections: connections,
	}

	return model
}

func (m networkModel) Init() tea.Cmd {
	return nil
}

func (m networkModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m networkModel) View() string {
	var v strings.Builder

	for _, c := range m.connections {
		for _, n := range c.networks {
			v.WriteString(n.name + "\n")
		}
	}

	return v.String()
}
