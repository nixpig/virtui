package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/digitalocean/go-libvirt"
)

type networkModel struct {
	connections []*libvirt.Libvirt
	networks    []libvirt.Network
}

func initNetwork(connections []*libvirt.Libvirt) networkModel {
	model := networkModel{
		connections: connections,
	}

	for _, c := range model.connections {
		n, _, err := c.ConnectListAllNetworks(1, 0)
		if err != nil {
			log.Error("failed to list networks", "err", err)
			continue
		}

		model.networks = append(model.networks, n...)
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

	for i, n := range m.networks {
		v.WriteString(fmt.Sprintf("%d - %s\n", i, n.Name))
	}

	return v.String()
}
