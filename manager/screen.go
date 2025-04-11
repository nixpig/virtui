package manager

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/vm"
)

type ScreenModel struct {
	machine vm.VM
	basic   basic
}

func InitScreenModel(v *vm.VM) ScreenModel {
	return ScreenModel{
		basic: basic{vm: v},
	}
}

func (m ScreenModel) Init() tea.Cmd {
	return nil
}

func (m ScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m ScreenModel) View() string {
	return m.basic.View()
}
