package guest

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/entity"
	"github.com/nixpig/virtui/internal/service"
	"github.com/nixpig/virtui/internal/tui/common"
	"libvirt.org/go/libvirtxml"
)

type Model interface {
	tea.Model
}

type guestModel struct {
	uuid     string
	keys     common.Keymap
	viewport viewport.Model
	service  service.Service
	domain   *entity.Domain
}

func NewModel(
	id string,
	svc service.Service,
	width, height int,
) tea.Model {
	d, err := svc.LookupDomainByUUIDString(id)
	if err != nil {
		// TODO: handle this a bit better by surfacing an error to the user
		log.Debug("get domain", "id", id, "err", err)
	}

	vp := viewport.New(width, height)
	vp.Style = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(1).
		Width(width)

	return guestModel{
		uuid:     id,
		keys:     common.Keys,
		service:  svc,
		domain:   &d,
		viewport: vp,
	}
}

func (m guestModel) Init() tea.Cmd {
	return nil
}

func (m guestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

	sb.WriteString("\nCPUs\n")
	sb.WriteString(
		"vCPU allocation: " + fmt.Sprintf("%d", m.domain.VCPU.Value) + "\n",
	)
	sb.WriteString("Mode: " + m.domain.CPU.Mode + "\n")

	sb.WriteString("\nMemory\n")
	sb.WriteString(
		"Current allocation: " + fmt.Sprintf(
			"%d",
			m.domain.CurrentMemory.Value,
		) + m.domain.CurrentMemory.Unit + "\n",
	)
	sb.WriteString(
		"Maximum allocation?: " + fmt.Sprintf(
			"%d",
			m.domain.Memory.Value,
		) + m.domain.Memory.Unit + "\n",
	)

	var mouse libvirtxml.DomainInput
	var keyboard libvirtxml.DomainInput

	for _, in := range m.domain.Devices.Inputs {
		switch in.Type {
		case "mouse":
			mouse = in
		case "keyboard":
			keyboard = in
		}
	}

	sb.WriteString(
		"\nKeyboard: " + fmt.Sprintf(
			"%s %s",
			keyboard.Bus,
			keyboard.Type,
		) + "\n",
	)

	sb.WriteString(
		"\nMouse: " + fmt.Sprintf("%s %s", mouse.Bus, mouse.Type) + "\n",
	)

	m.viewport.SetContent(sb.String())

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Back):
			return m, common.GoBackCmd()
		case key.Matches(msg, m.keys.Down):
			m.viewport.ScrollDown(1)
		case key.Matches(msg, m.keys.Up):
			m.viewport.ScrollUp(1)

		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m guestModel) View() string {

	return m.viewport.View()
}
