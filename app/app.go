package app

import (
	"database/sql"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/connection"
	"github.com/nixpig/virtui/event"
	"github.com/nixpig/virtui/keys"
)

type appModel struct {
	screenModel tea.Model

	cr connection.ConnectionRepository

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
	var currentModel tea.Model

	currentModel = connectionsScreen()

	m := appModel{
		screenModel: currentModel,

		cr: connection.NewConnectionRepositoryImpl(db),

		help: help.New(),
		keys: keys.Global,
	}

	return m
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
	// var cmd tea.Cmd

	switch msg := msg.(type) {
	case event.VM:
		log.Debug("handle vm event", "id", msg.ID, "event", msg.Event)
		// return m, tea.Batch(vmEventMsg(m.vmEventCh))

	case tea.WindowSizeMsg:
		log.Debug("handle window resize", "width", msg.Width, "height", msg.Height)
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case msg.String() == "d":
			m.screenModel = domainsScreen()
		case msg.String() == "c":
			m.screenModel = connectionsScreen()

		default:
			// pass remaining keys down to child screen model
			m.screenModel.Update(msg)
		}
	}

	return m, nil
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
	helpBorderStyle := lipgloss.NormalBorder()
	helpBorderColor := lipgloss.Color("gray")

	helpStyle := lipgloss.NewStyle().
		BorderTop(true).
		BorderForeground(helpBorderColor).
		BorderStyle(helpBorderStyle).
		Width(containerWidth)

	helpView := "\n" + helpStyle.Render(m.help.View(m.keys))

	// content
	contentHeight := containerHeight - strings.Count(helpView, "\n")
	contentStyle := lipgloss.NewStyle().
		Height(contentHeight)

	contentView := contentStyle.Render(m.screenModel.View())

	return containerStyle.Render(contentView + helpView)
}

func (m appModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.screenModel = model

	return m.screenModel, m.screenModel.Init()
}
