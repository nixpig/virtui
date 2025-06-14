package network

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/internal/entity"
	"libvirt.org/go/libvirt"
)

type Model struct {
	lv       *libvirt.Connect
	networks []*entity.Network
}

// New creates tea.Model for the network view
func New(lv *libvirt.Connect) tea.Model {
	networks, _ := lv.ListAllNetworks(0)

	m := Model{
		lv:       lv,
		networks: make([]*entity.Network, len(networks)),
	}

	for i, n := range networks {
		m.networks[i], _ = entity.ToNetworkStruct(&n)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder

	for _, n := range m.networks {
		sb.WriteString(n.Name + " " + n.UUID)
		sb.WriteString("\n")
	}

	return sb.String()
}
