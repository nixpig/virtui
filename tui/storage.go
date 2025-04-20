package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/digitalocean/go-libvirt"
)

type storageModel struct {
	connections []*libvirt.Libvirt
	pools       []libvirt.StoragePool
}

func initStorage(connections []*libvirt.Libvirt) storageModel {
	model := storageModel{
		connections: connections,
	}

	for _, c := range model.connections {
		p, _, err := c.ConnectListAllStoragePools(1, 0)
		if err != nil {
			log.Error("failed to list storage pools", "err", err)
			continue
		}

		model.pools = append(model.pools, p...)
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

	for i, p := range m.pools {
		v.WriteString(fmt.Sprintf("%d - %s\n", i, p.Name))
	}

	return v.String()
}
