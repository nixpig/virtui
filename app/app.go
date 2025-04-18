package app

import (
	"database/sql"

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

type switchScreenMsg int

const (
	DASHBOARD_SCREEN switchScreenMsg = iota
	CONNECTIONS_SCREEN
	DOMAINS_SCREEN
	ADD_CONNECTION_SCREEN
	NEW_VM_SCREEN
)

type appModel struct {
	screenModel tea.Model

	cr connection.ConnectionStore

	help help.Model
	keys keys.GlobalMap

	width  int
	height int

	/*
		confirm
		oops

		dashboard
		vms
		connections

		addConnection
		newVM

	*/
}

// func vmEventMsg(ch chan event.VM) tea.Cmd {
// 	return func() tea.Msg {
// 		return <-ch
// 	}
// }

// func vmEventCallback(ch chan event.VM) libvirt.DomainEventLifecycleCallback {
// 	return func(
// 		c *libvirt.Connect,
// 		d *libvirt.Domain,
// 		l *libvirt.DomainEventLifecycle,
// 	) {
// 		ch <- event.VM{
// 			Event: l.Event,
// 		}
// 	}
// }

func InitModel(db *sql.DB) appModel {
	cr := connection.NewConnectionStoreImpl(db)

	model := appModel{
		help: help.New(),
		keys: keys.Global,
	}

	model.screenModel = dashboardScreen(cr, model.Update)

	return model
}

func (m appModel) Init() tea.Cmd {

	// need to create a new libvirt connection for each connection
	// then need to do this for _each_ of the connections
	// how do we 'map' these so events from one connection map to the vms for that connection?
	// for _, c := range m.connections {
	// 	if _, err := .DomainEventLifecycleRegister(
	// 		nil,
	// 		vmEventCallback(m.vmEventCh),
	// 	); err != nil {
	// 		log.Error("register domain event handler", "err", err)
	// 		// TODO: handle err
	// 		fmt.Println(err)
	// 	}
	// }

	// return tea.Batch(vmEventMsg(m.vmEventCh))
	return m.screenModel.Init()
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case switchScreenMsg:
		switch msg {
		case DASHBOARD_SCREEN:
			return m.switchScreen(dashboardScreen(m.cr, m.Update))
		case CONNECTIONS_SCREEN:
			return m.switchScreen(connectionsScreen())
		case DOMAINS_SCREEN:
			log.Error("todo!")
		case ADD_CONNECTION_SCREEN:
			return m.switchScreen(addConnectionScreen())
		case NEW_VM_SCREEN:
			log.Error("todo!")

		}

	// case event.VM:
	// 	log.Debug("handle vm event", "id", msg.ID, "event", msg.Event)
	// 	// return m, tea.Batch(vmEventMsg(m.vmEventCh))

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		default:
			// pass remaining keys down to child screen model
			m.screenModel, cmd = m.screenModel.Update(msg)
			cmds = append(cmds, cmd)
		}

	default:
		m.screenModel, cmd = m.screenModel.Update(msg)
		cmds = append(cmds, cmd)
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
	content := m.screenModel.View()
	contentStyle := lipgloss.NewStyle().
		Height(containerHeight - helpHeight)

	contentView := contentStyle.Render(content)

	return containerStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top, contentView, helpView),
	)
}

func (m appModel) switchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.screenModel = model

	return m.screenModel, m.screenModel.Init()
}
