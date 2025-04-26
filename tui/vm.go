package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/digitalocean/go-libvirt"
)

type vmModel struct {
	connections map[string]*libvirt.Libvirt
	uuid        string
}

func initVM(connections map[string]*libvirt.Libvirt, uuid string) vmModel {
	// TODO: use vmUUID to get domain details

	model := vmModel{
		connections: connections,
		uuid:        uuid,
	}

	return model
}

func (m vmModel) Init() tea.Cmd {
	return nil
}

func (m vmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m vmModel) View() string {
	var v strings.Builder

	v.WriteString("vm: " + m.uuid)

	return v.String()
}
