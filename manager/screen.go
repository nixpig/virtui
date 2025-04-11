package manager

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/vm"
)

type ScreenModel struct {
	basic basic
}

func InitScreenModel(v *vm.VM) ScreenModel {
	return ScreenModel{
		basic{
			name:  v.GetPresentableName(),
			id:    v.GetPresentableID(),
			uuid:  v.GetPresentableUUID(),
			state: v.GetPresentableState(),
		},
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
