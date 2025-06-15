package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/tui/entity"
	"libvirt.org/go/libvirt"
)

type networkModel struct {
	conn     *libvirt.Connect
	networks []entity.Network
}

// New creates tea.Model for the network view
func newNetworkModel(conn *libvirt.Connect) tea.Model {
	networks, _ := conn.ListAllNetworks(0)

	m := networkModel{
		conn:     conn,
		networks: make([]entity.Network, len(networks)),
	}

	for i, n := range networks {
		m.networks[i], _ = entity.ToNetworkStruct(&n)
	}

	return m
}

func (m networkModel) Init() tea.Cmd {
	return nil
}

func (m networkModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m networkModel) View() string {
	var sb strings.Builder

	for _, n := range m.networks {
		sb.WriteString(n.Name + " " + n.UUID)
		sb.WriteString("\n")
	}

	return sb.String()
}
