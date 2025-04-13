package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type connectionsModel struct{}

func connectionsScreen() connectionsModel {
	return connectionsModel{}
}

func (m connectionsModel) Init() tea.Cmd {
	return nil
}

func (m connectionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "x":
			log.Info("x pressed in connections screen")

		}
	}
	return m, nil
}

func (m connectionsModel) View() string {
	return "connections screen view"
}
