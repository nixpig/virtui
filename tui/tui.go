package tui

import (
	"net/url"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/digitalocean/go-libvirt"
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

	activeModel tea.Model
	connections map[string]*libvirt.Libvirt

	width  int
	height int

	helpModels map[string][]tea.Model
	keyMaps    map[string]map[string]key.Binding

	tabs      []string
	activeTab int
}

func InitTUI(store connection.ConnectionStore) appModel {
	model := appModel{
		store: store,
		help:  help.New(),
		keys:  keys.Global,

		connections: make(map[string]*libvirt.Libvirt),

		tabs:      []string{"(1) Virtual Machines", "(2) Networks", "(3) Storage"},
		activeTab: 0,
	}

	if _, err := model.store.GetConnectionByURI(string(libvirt.QEMUSystem)); err != nil {
		log.Debug("system connection not found; insert")
		if err := model.store.InsertConnection(&connection.Connection{
			URI: string(libvirt.QEMUSystem),
		}); err != nil {
			log.Error("failed to insert system connection")
		}
	}

	// if _, err := model.store.GetConnectionByURI(string(libvirt.QEMUSession)); err != nil {
	// 	log.Debug("session connection not found; insert")
	// 	if err := model.store.InsertConnection(&connection.Connection{
	// 		URI: string(libvirt.QEMUSession),
	// 	}); err != nil {
	// 		log.Error("failed to insert session connection")
	// 	}
	// }

	conns, err := model.store.GetConnections()
	if err != nil {
		log.Error("failed to get connections from store", "err", err)
	}

	for _, c := range conns {
		uri, err := url.Parse(c.URI)
		if err != nil {
			log.Error("failed to parse uri", "err", err)
			continue
		}

		// TODO: / FIXME: don't forget to close these connections on app exit!!!
		lv, err := libvirt.ConnectToURI(uri)
		if err != nil {
			log.Error("failed to connect to uri", "uri", uri, "err", err)
			continue
		}

		model.connections[c.URI] = lv
	}

	model.activeModel = initDashboard(model.connections)

	return model
}

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.activeModel, cmd = m.activeModel.Update(msg)
		cmds = append(cmds, cmd)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.Global.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, keys.Global.Dashboard):
			m.activeTab = 0
			m.activeModel = initDashboard(m.connections)
			m.activeModel.Init()
		case key.Matches(msg, m.keys.Networks):
			m.activeTab = 1
			m.activeModel = initNetwork(m.connections)
			m.activeModel.Init()
		case key.Matches(msg, keys.Global.Storage):
			m.activeTab = 2
			m.activeModel = initStorage(m.connections)
			m.activeModel.Init()
		default:
			m.activeModel, cmd = m.activeModel.Update(msg)
			cmds = append(cmds, cmd)
		}
	default:
		m.activeModel, cmd = m.activeModel.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m appModel) View() string {
	// tabs
	renderedTabs := make([]string, len(m.tabs))

	for i, t := range m.tabs {

		borderForeground := lipgloss.Color("#999999")
		// borderBottom := true

		// if m.activeTab == i {
		// 	borderForeground = lipgloss.Color("#ffffff")
		// 	// borderBottom = false
		// }

		tabStyle := lipgloss.NewStyle().
			BorderRight(true).
			BorderStyle(lipgloss.NormalBorder()).
			// BorderTop(true).
			// BorderRight(true).
			// BorderBottom(borderBottom).
			// BorderLeft(true).
			// Margin(0).
			PaddingLeft(1).
			PaddingRight(1).
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
		Width(containerWidth - lipgloss.ASCIIBorder().GetLeftSize() - lipgloss.ASCIIBorder().GetRightSize()).
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(lipgloss.Color("#555555"))

	helpView := helpStyle.Render(m.help.View(m.keys))
	helpHeight := lipgloss.Height(helpView)

	// content
	content := m.activeModel.View()

	contentStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Height(containerHeight - tabHeight - lipgloss.ASCIIBorder().GetBottomSize() - lipgloss.ASCIIBorder().GetTopSize()).
		Width(containerWidth - lipgloss.ASCIIBorder().GetLeftSize() - lipgloss.ASCIIBorder().GetRightSize())

	padding := lipgloss.NewStyle().Height(contentStyle.GetHeight() - lipgloss.Height(content) - helpHeight).Render("")

	contentView := contentStyle.Render(lipgloss.JoinVertical(lipgloss.Top, content, padding, helpView))

	return containerStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, renderedTabRow, contentView),
	)
}
