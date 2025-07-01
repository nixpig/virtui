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
		log.Debug("list all storage pools", "err", err)
	}

	m := storageModel{
		conn:    conn,
		storage: make(map[entity.StoragePool][]entity.StorageVolume, len(pools)),
	}

	for _, pool := range pools {
		p, err := entity.ToStoragePoolStruct(&pool)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("convert entity to struct", "err", err, "pool", pool)
		}
		m.storage[p] = []entity.StorageVolume{}

		volumes, err := pool.ListAllStorageVolumes(0)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("list all storage volumes", "err", err)
		}

		if err := pool.Free(); err != nil {
			log.Warn("free ref counted pool struct", "err", err)
		}

		if len(volumes) == 0 {
			continue
		}

		for _, volume := range volumes {
			v, err := entity.ToStorageVolume(&volume)
			if err != nil {
				// TODO: surface error to user?
				log.Debug("convert entity to struct", "err", err, "volume", volume)
			}

			if err := volume.Free(); err != nil {
				log.Warn("free ref counted volume struct", "err", err)
			}

			m.storage[p] = append(m.storage[p], v)
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
		sb.WriteString("Name: " + k.Name + "\n")
		sb.WriteString("UUID: " + k.UUID + "\n")
		sb.WriteString("Type: " + k.Type + "\n")
		sb.WriteString("Size: " + fmt.Sprintf("%d%s (%d%s available)", k.Capacity.Value, k.Capacity.Unit, k.Available.Value, k.Available.Unit) + "\n")
		sb.WriteString("Location: " + k.Target.Path + "\n")
		sb.WriteString("Volumes:\n")

		for i, x := range v {
			sb.WriteString(fmt.Sprintf("%d %s - %d%s - %s \n", i, x.Name, x.Capacity.Value, x.Capacity.Unit, x.Target.Format.Type))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
