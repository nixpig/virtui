package storage

import (
	tea "github.com/charmbracelet/bubbletea"
	"libvirt.org/go/libvirt"
)

type Model struct {
	lv *libvirt.Connect
}

// New creates a tea.Model for the storage view
func New(lv *libvirt.Connect) tea.Model {
	return Model{lv}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "storage view"
}
