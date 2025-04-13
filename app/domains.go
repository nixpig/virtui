package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type domainsModel struct{}

func domainsScreen() domainsModel {
	return domainsModel{}
}

func (m domainsModel) Init() tea.Cmd {
	return nil
}

func (m domainsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "y":
			log.Info("y pressed in connections screen")

		}
	}
	return m, nil
}

func (m domainsModel) View() string {
	return "domains screen view"
}
