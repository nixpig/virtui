package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/connection"
	"github.com/nixpig/virtui/keys"
)

type errMsg struct{ error }

func (e errMsg) Error() string {
	return e.error.Error()
}

type appModel struct {
	store connection.ConnectionStore
	help  help.Model
	keys  keys.GlobalMap

	width  int
	height int
}

func InitModel(store connection.ConnectionStore) appModel {
	model := appModel{
		store: store,
		help:  help.New(),
		keys:  keys.Global,
	}

	return model
}

func (m appModel) Init() tea.Cmd {
	if _, err := m.store.GetConnectionByURI("qemu:///system"); err != nil {
		log.Debug("system connection not found; insert")
		if err := m.store.InsertConnection(&connection.Connection{
			URI: "qemu:///system",
		}); err != nil {
			log.Error("failed to insert system connection")
		}
	}

	if _, err := m.store.GetConnectionByURI("qemu:///session"); err != nil {
		log.Debug("session connection not found; insert")
		if err := m.store.InsertConnection(&connection.Connection{
			URI: "qemu:///session",
		}); err != nil {
			log.Error("failed to insert session connection")
		}
	}

	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.Global.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		}
	}

	return m, tea.Batch(cmds...)
}

func (m appModel) View() string {
	// container
	containerBorderStyle := lipgloss.RoundedBorder()
	containerBorderColor := lipgloss.Color("63")

	containerWidth :=
		m.width - (containerBorderStyle.GetLeftSize() + containerBorderStyle.GetRightSize())

	containerHeight :=
		m.height - (containerBorderStyle.GetTopSize() + containerBorderStyle.GetBottomSize())

	containerStyle := lipgloss.NewStyle().
		BorderStyle(containerBorderStyle).
		BorderForeground(containerBorderColor).
		Width(containerWidth).
		Height(containerHeight)

	// help
	helpStyle := lipgloss.NewStyle().
		Width(containerWidth)

	helpView := helpStyle.Render(m.help.View(m.keys))
	helpHeight := lipgloss.Height(helpView)

	// content
	content := "some content in here"
	contentStyle := lipgloss.NewStyle().
		Height(containerHeight - helpHeight)

	contentView := contentStyle.Render(content)

	return containerStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentView, helpView),
	)
}
