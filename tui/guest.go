package tui

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/tui/entity"
	"libvirt.org/go/libvirt"
)

type guestModel struct {
	activeGuestUUID string
	keys            keymap
	conn            *libvirt.Connect
	domain          *entity.Domain
}

// New creates a tea.Model for the guest view
func newGuestModel(id string, conn *libvirt.Connect) tea.Model {
	dom, err := conn.LookupDomainByUUIDString(id)
	if err != nil {
		log.Fatal(err)
	}

	d, _ := entity.ToDomainStruct(dom)

	return guestModel{
		activeGuestUUID: id,
		keys:            keys,
		conn:            conn,
		domain:          &d,
	}
}

func (m guestModel) Init() tea.Cmd {
	return nil
}

func (m guestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.back):
			return m, goBackCmd()

		}
	}
	return m, nil
}

func (m guestModel) View() string {
	var sb strings.Builder

	sb.WriteString("Basic details\n")
	sb.WriteString("Name: " + m.domain.Name + "\n")
	sb.WriteString("UUID: " + m.domain.UUID + "\n")
	sb.WriteString("State: " + "\n")
	sb.WriteString("Title: " + m.domain.Title + "\n")
	sb.WriteString("Description: " + m.domain.Description + "\n")

	sb.WriteString("\nHypervisor details\n")
	sb.WriteString("Hypervisor: " + strings.ToUpper(m.domain.Type) + "\n")
	sb.WriteString("Architecture: " + m.domain.OS.Type.Arch + "\n")
	sb.WriteString("Emulator: " + m.domain.Devices.Emulator + "\n")
	sb.WriteString("Chipset: " + m.domain.OS.Type.Machine + "\n")
	sb.WriteString("Firmware: " + m.domain.OS.Firmware + "\n")

	return sb.String()
}
