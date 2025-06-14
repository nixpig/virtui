package main

import (
	"context"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/term"
	"github.com/nixpig/virtui/internal/guest"
	"github.com/nixpig/virtui/internal/keys"
	"github.com/nixpig/virtui/internal/manager"
	"github.com/nixpig/virtui/internal/network"
	"github.com/nixpig/virtui/internal/storage"
	"libvirt.org/go/libvirt"
)

type state int

const (
	managerView state = iota // table that shows connections and domains under each
	guestView                // view of an individual domain
	networkView              // view of networks
	storageView              // view of storage
)

type MainModel struct {
	state         state
	keys          keys.Keymap
	help          help.Model
	managerModel  tea.Model
	guestModel    tea.Model
	networkModel  tea.Model
	storageModel  tea.Model
	lv            *libvirt.Connect
	activeGuestID uint
	width         int
	height        int
}

func initialModel(lv *libvirt.Connect) MainModel {
	domains, err := lv.ListAllDomains(0)
	if err != nil {
		log.Fatal(err)
	}

	managerModel := manager.New(domains)

	width, height, err := term.GetSize(os.Stdin.Fd())
	if err != nil {
		log.Fatal("failed to get size of terminal", "err", err)
	}

	return MainModel{
		state:        managerView,
		keys:         keys.Keys,
		help:         help.New(),
		managerModel: managerModel,
		lv:           lv,
		width:        width,
		height:       height,
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Manager):
			// TODO: initialise managerView?
			m.state = managerView
		case key.Matches(msg, m.keys.Network):
			m.networkModel = network.New()
			m.state = networkView
		case key.Matches(msg, m.keys.Storage):
			m.storageModel = storage.New()
			m.state = storageView
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

	case guest.BackMsg:
		m.state = managerView

	case manager.SelectMsg:
		m.activeGuestID = msg.ActiveGuestId
		m.guestModel = guest.New(m.activeGuestID)
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

	case networkView:
		networkModel, newCmd := m.networkModel.Update(msg)
		m.networkModel = networkModel
		cmd = newCmd

	case storageView:
		storageModel, newCmd := m.storageModel.Update(msg)
		m.storageModel = storageModel
		cmd = newCmd

	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	helpView := m.help.View(m.keys)

	var mainView string

	switch m.state {
	case managerView:
		mainView = m.managerModel.View()
	case guestView:
		mainView = m.guestModel.View()
	case networkView:
		mainView = m.networkModel.View()
	case storageView:
		mainView = m.storageModel.View()
	default:
		mainView = m.managerModel.View()
	}

	offset := 1 // who knows where this comes from 🤷

	padding := m.height - offset - strings.Count(mainView, "\n") - strings.Count(helpView, "\n")

	return mainView + strings.Repeat("\n", padding) + helpView
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
