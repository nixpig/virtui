package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/keys"
)

type dashboardModel struct {
	keys    keys.DashboardMap
	help    help.Model
	content string
}

func dashboardScreen() dashboardModel {
	return dashboardModel{
		help:    help.New(),
		keys:    keys.Dashboard,
		content: "initial",
	}
}

func (m dashboardModel) Init() tea.Cmd {
	return nil
}

func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.AddConnection):
			log.Info("add connection")

		case key.Matches(msg, m.keys.NewVM):
			m.content = "new vm!!"
			log.Info("new vm")

		case key.Matches(msg, m.keys.Help):
			log.Info("toggle help", "showall", m.help.ShowAll)
			m.help.ShowAll = !m.help.ShowAll

		}
	}

	return m, nil
}

func (m dashboardModel) View() string {
	return "dashboard screen" + m.content + "\n" + m.help.View(m.keys)
}
