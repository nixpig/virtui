package storage

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/internal/entity"
	"libvirt.org/go/libvirt"
)

type Model struct {
	lv      *libvirt.Connect
	storage map[*entity.StoragePool][]*entity.StorageVolume
}

// New creates a tea.Model for the storage view
func New(lv *libvirt.Connect) tea.Model {
	pools, _ := lv.ListAllStoragePools(0)

	m := Model{
		lv:      lv,
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

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
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
