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

	helpModels map[string][]tea.Model
	keyMaps    map[string]map[string]key.Binding

	tabs      []string
	activeTab int
}

func InitModel(store connection.ConnectionStore) appModel {
	model := appModel{
		store: store,
		help:  help.New(),
		keys:  keys.Global,

		tabs:      []string{"(1) Virtual Machines", "(2) Networks", "(3) Storage"},
		activeTab: 0,
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
		case key.Matches(msg, keys.Global.Dashboard):
			m.activeTab = 0
		case key.Matches(msg, m.keys.Networks):
			m.activeTab = 1
		case key.Matches(msg, keys.Global.Storage):
			m.activeTab = 2
		}
	}

	return m, tea.Batch(cmds...)
}

func (m appModel) View() string {
	// tabs
	renderedTabs := make([]string, len(m.tabs))

	for i, t := range m.tabs {

		borderForeground := lipgloss.Color("#999999")
		borderBottom := true

		if m.activeTab == i {
			borderForeground = lipgloss.Color("#ffffff")
			// borderBottom = false
		}

		tabStyle := lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderRight(true).
			BorderBottom(borderBottom).
			BorderLeft(true).
			Margin(0).
			BorderForeground(borderForeground)

		renderedTabs[i] = tabStyle.Render(t)
	}

	renderedTabRow := lipgloss.NewStyle().
		Padding(0).
		Margin(0).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...))
	tabHeight := lipgloss.Height(renderedTabRow)

	// container

	containerWidth := m.width
	containerHeight := m.height

	containerStyle := lipgloss.NewStyle().
		Width(containerWidth).
		Height(containerHeight).
		Margin(0)

	// help
	helpStyle := lipgloss.NewStyle().
		Width(containerWidth - lipgloss.ASCIIBorder().GetRightSize() - lipgloss.ASCIIBorder().GetLeftSize()).
		BorderStyle(lipgloss.NormalBorder())

	helpView := helpStyle.Render(m.help.View(m.keys))
	helpHeight := lipgloss.Height(helpView)

	// content
	content := m.tabs[m.activeTab]
	contentStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Height(containerHeight - helpHeight - tabHeight - lipgloss.ASCIIBorder().GetBottomSize() - lipgloss.ASCIIBorder().GetTopSize()).
		Width(containerWidth - lipgloss.ASCIIBorder().GetLeftSize() - lipgloss.ASCIIBorder().GetRightSize())

	contentView := contentStyle.Render(content)

	return containerStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, renderedTabRow, contentView, helpView),
	)
}
