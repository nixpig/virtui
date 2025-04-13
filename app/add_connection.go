package app

import tea "github.com/charmbracelet/bubbletea"

type addConnectionModel struct{}

func addConnectionScreen() addConnectionModel {
	return addConnectionModel{}
}

func (m addConnectionModel) Init() tea.Cmd {
	return nil
}

func (m addConnectionModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m addConnectionModel) View() string {
	return "ADD CONNECTION SCREEN"
}
