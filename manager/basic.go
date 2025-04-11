package manager

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/vm"
)

type basic struct {
	vm *vm.VM
}

func (m basic) Init() tea.Cmd {
	return nil
}

func (m basic) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m basic) View() string {
	uuid, _ := m.vm.GetUUIDString()
	return fmt.Sprintf(
		"%s | %s | %s | %s",
		m.vm.GetPresentableName(),
		uuid,
		m.vm.GetPresentableState(),
		"something",
	)
}
