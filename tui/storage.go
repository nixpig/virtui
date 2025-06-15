package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/tui/entity"
	"libvirt.org/go/libvirt"
)

type storageModel struct {
	conn    *libvirt.Connect
	storage map[*entity.StoragePool][]*entity.StorageVolume
}

// New creates a tea.Model for the storage view
func newStorageModel(conn *libvirt.Connect) tea.Model {
	pools, _ := conn.ListAllStoragePools(0)

	m := storageModel{
		conn:    conn,
		storage: make(map[*entity.StoragePool][]*entity.StorageVolume, len(pools)),
	}

	for _, p := range pools {
		x, _ := entity.ToStoragePoolStruct(&p)
		m.storage[x] = []*entity.StorageVolume{}

		vols, _ := p.ListAllStorageVolumes(0)
		for _, v := range vols {
			y, _ := entity.ToStorageVolume(&v)
			m.storage[x] = append(m.storage[x], y)
		}
	}

	return m
}

func (m storageModel) Init() tea.Cmd {
	return nil
}

func (m storageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m storageModel) View() string {
	var sb strings.Builder

	for k, v := range m.storage {
		sb.WriteString(k.Name + " - ")
		for i, x := range v {
			sb.WriteString(fmt.Sprintf("\n%d %s", i, x.Name))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
