package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/tui/entity"
	"libvirt.org/go/libvirt"
)

type networkModel struct {
	conn     *libvirt.Connect
	networks []entity.Network
}

func newNetworkModel(conn *libvirt.Connect) tea.Model {
	networks, err := conn.ListAllNetworks(0)
	if err != nil {
		// TODO: surface error to user?
		log.Debug("list all networks", "err", err)
	}

	m := networkModel{
		conn:     conn,
		networks: make([]entity.Network, len(networks)),
	}

	for i, network := range networks {
		m.networks[i], err = entity.ToNetworkStruct(&network)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("covert entity to struct", "err", err, "network", network)
		}

		if err := network.Free(); err != nil {
			log.Warn("free ref counted network struct", "err", err)
		}
	}

	return m
}

func (m networkModel) Init() tea.Cmd {
	return nil
}

func (m networkModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Debug("network received msg", "type", fmt.Sprintf("%T", msg), "data", msg)
	return m, nil
}

func (m networkModel) View() string {
	var sb strings.Builder

	for _, network := range m.networks {
		sb.WriteString(network.Name + " " + network.UUID)
		sb.WriteString("\n")
	}

	return sb.String()
}
