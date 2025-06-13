package main

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/guest"
	"github.com/nixpig/virtui/internal/manager"
	"libvirt.org/go/libvirt"
)

type state int

const (
	managerView state = iota // table that shows connections and domains under each
	guestView                // view of an individual domain
)

type MainModel struct {
	state         state
	managerModel  tea.Model
	guestModel    tea.Model
	lv            *libvirt.Connect
	activeGuestID uint
}

func initialModel(lv *libvirt.Connect) MainModel {
	managerModel := manager.New(lv)

	return MainModel{
		state:        managerView,
		managerModel: managerModel,
		lv:           lv,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case guest.BackMsg:
		m.state = managerView

	case manager.SelectMsg:
		m.activeGuestID = msg.ActiveGuestId
		m.state = guestView

	}

	switch m.state {
	case managerView:
		managerModel, newCmd := m.managerModel.Update(msg)
		m.managerModel = managerModel
		cmd = newCmd

	case guestView:
		guestModel, newCmd := m.guestModel.Update(msg)
		m.guestModel = guestModel
		cmd = newCmd

	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	switch m.state {
	case managerView:
		return m.managerModel.View()
	case guestView:
		return m.guestModel.View()
	default:
		return m.managerModel.View()
	}
}

func main() {
	ctx := context.Background()

	uri := "qemu:///system"
	lv, err := libvirt.NewConnect(uri)
	if err != nil {
		log.Fatal("failed to connect to libvirt", "uri", uri, "err", err)
	}

	p := tea.NewProgram(
		initialModel(lv),
		tea.WithAltScreen(),
		tea.WithContext(ctx),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal("failed to start program", "err", err)
	}
}
