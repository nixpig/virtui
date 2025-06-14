package manager

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/internal/entity"
	"libvirt.org/go/libvirt"
)

type SelectMsg struct {
	ActiveGuestId uint
}

type Model struct {
	domains []libvirt.Domain
}

func New(domains []libvirt.Domain) tea.Model {
	return Model{
		domains: domains,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var view strings.Builder

	for i, d := range m.domains {
		if i != 0 {
			view.WriteString("\n")
		}

		x, err := entity.ToDomainStruct(&d)
		if err != nil {
			view.WriteString("invalid domain: " + err.Error())
		}

		state, _, _ := d.GetState()

		view.WriteString(fmt.Sprintf("%s: %s (%s)", x.Name, state, x.UUID))
	}

	return view.String()
}
