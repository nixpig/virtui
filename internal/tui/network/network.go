package network

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/entity"
	"github.com/nixpig/virtui/internal/service"
)

type Model interface {
	tea.Model
}

type networkModel struct {
	service  service.Service
	networks []entity.Network
}

func NewModel(svc service.Service) tea.Model {
	networks, err := svc.ListAllNetworks()
	if err != nil {
		log.Debug("list all networks", "err", err)
		return networkModel{
			service:  svc,
			networks: []entity.Network{},
		}
	}

	m := networkModel{
		service:  svc,
		networks: networks,
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

	for _, network := range m.networks {
		sb.WriteString("Name: " + network.Name + "\n")
		sb.WriteString("UUID: " + network.UUID + "\n")
		sb.WriteString("Bridge device: " + network.Bridge.Name + "\n")

		for _, ip := range network.IPs {
			sb.WriteString("Address: " + ip.Address + "\n")
			sb.WriteString("Netmask: " + ip.Netmask + "\n")
			for _, dhcp := range ip.DHCP.Ranges {
				sb.WriteString(
					"DHCP range: " + dhcp.Start + " - " + dhcp.End + "\n",
				)
			}
		}

		sb.WriteString(
			"Forwarding mode: " + strings.ToUpper(network.Forward.Mode) + "\n",
		)
		sb.WriteString("\n\n")
	}

	return sb.String()
}
