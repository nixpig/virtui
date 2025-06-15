package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/tui/entity"
	"libvirt.org/go/libvirt"
)

type storageModel struct {
	conn    *libvirt.Connect
	storage map[entity.StoragePool][]entity.StorageVolume
}

func newStorageModel(conn *libvirt.Connect) tea.Model {
	pools, err := conn.ListAllStoragePools(0)
	if err != nil {
		// TODO: surface error to user?
		log.Debug("failed to list all storage pools", "err", err)
	}

	m := storageModel{
		conn:    conn,
		storage: make(map[entity.StoragePool][]entity.StorageVolume, len(pools)),
	}

	for _, p := range pools {
		x, err := entity.ToStoragePoolStruct(&p)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to convert storage pool to struct", "err", err)
		}
		m.storage[x] = []entity.StorageVolume{}

		vols, err := p.ListAllStorageVolumes(0)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to list all storage volumes", "err", err)
			continue
		}

		for _, v := range vols {
			y, err := entity.ToStorageVolume(&v)
			if err != nil {
				// TODO: surface error to user?
				log.Debug("failed to convert storage volume to struct", "err", err)
				continue
			}

			m.storage[x] = append(m.storage[x], y)
		}
	}

	return m
}

func (m storageModel) Init() tea.Cmd {
	return nil
}

func (m storageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Debug("storage received msg", "type", fmt.Sprintf("%T", msg), "data", msg)
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
