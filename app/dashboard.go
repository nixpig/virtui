package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/connection"
	"github.com/nixpig/virtui/keys"
)

type hasConnectionsMsg bool

type dashboardModel struct {
	help help.Model
	keys keys.DashboardMap

	updater func(tea.Msg) (tea.Model, tea.Cmd)

	cr connection.ConnectionRepository

	content string

	hasConnections bool

	width  int
	height int
}

func dashboardScreen(cr connection.ConnectionRepository, updater func(tea.Msg) (tea.Model, tea.Cmd)) dashboardModel {
	return dashboardModel{
		keys:    keys.Dashboard,
		help:    help.New(),
		updater: updater,

		cr: cr,

		content: "initial",
	}
}

func (m dashboardModel) checkConnections() tea.Msg {
	c, err := m.cr.HasConnections()
	if err != nil {
		return errMsg{err}
	}

	return hasConnectionsMsg(c)
}

func (m dashboardModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.checkConnections)
	return tea.Batch(cmds...)
}

func (m dashboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case hasConnectionsMsg:
		m.hasConnections = bool(msg)
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.AddConnection):
			return m.updater(ADD_CONNECTION_SCREEN)

		case key.Matches(msg, m.keys.NewVM):
			m.content = "new vm!!"

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		}
	}

	return m, nil
}

func (m dashboardModel) View() string {
	s := "content"
	if !m.hasConnections {
		s += " no connections"
	}

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		BorderStyle(lipgloss.RoundedBorder()).
		Render("dashboard screen" + s + "\n" + m.help.View(m.keys))
}
