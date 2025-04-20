package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type storageModel struct {
	connections []qemuConnection
}

func initStorage(connections []qemuConnection) storageModel {
	model := storageModel{
		connections: connections,
	}

	return model
}

func (m storageModel) Init() tea.Cmd {
	return nil
}

func (m storageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m storageModel) View() string {
	var v strings.Builder

	for _, c := range m.connections {
		for _, s := range c.storage {
			v.WriteString(s.name + "\n")
		}
	}

	return v.String()
}
