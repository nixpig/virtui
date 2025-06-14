package guest

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/internal/commands"
	"github.com/nixpig/virtui/internal/entity"
	"github.com/nixpig/virtui/internal/keys"
	"libvirt.org/go/libvirt"
)

type Model struct {
	activeGuestUUID string
	keys            keys.Keymap
	lv              *libvirt.Connect
	domain          *entity.Domain
}

// New creates a tea.Model for the guest view
func New(id string, lv *libvirt.Connect) tea.Model {
	dom, err := lv.LookupDomainByUUIDString(id)
	if err != nil {
		log.Fatal(err)
	}

	d, _ := entity.ToDomainStruct(dom)

	return Model{
		activeGuestUUID: id,
		keys:            keys.Keys,
		lv:              lv,
		domain:          d,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Back):
			return m, commands.GoBackCmd()

		}
	}
	return m, nil
}

func (m Model) View() string {
	var sb strings.Builder

	sb.WriteString(m.domain.Name + "\n")
	sb.WriteString(m.domain.UUID + "\n")
	sb.WriteString(m.domain.Type + "\n")
	sb.WriteString(m.domain.OS.Type.Arch + "\n")
	sb.WriteString(m.domain.OS.Type.Machine + "\n")
	sb.WriteString(m.domain.OS.Type.Type + "\n")

	return sb.String()
}
