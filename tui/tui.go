package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/term"
	"libvirt.org/go/libvirt"
)

type state int

const (
	managerView state = iota
	guestView
	networkView
	storageView
)

type model struct {
	state         state
	keys          keymap
	help          help.Model
	managerModel  tea.Model
	guestModel    tea.Model
	networkModel  tea.Model
	storageModel  tea.Model
	conn          *libvirt.Connect
	activeGuestID uint
	width         int
	height        int
}

func New(conn *libvirt.Connect) model {
	defaultModel := newManagerModel(conn)

	width, height, err := term.GetSize(os.Stdin.Fd())
	if err != nil {
		log.Fatal("failed to get size of terminal", "fd", os.Stdin.Fd(), "err", err)
	}

	return model{
		state:        managerView,
		keys:         keys,
		help:         help.New(),
		managerModel: defaultModel,
		conn:         conn,
		width:        width,
		height:       height,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	log.Debug("tui received msg", "type", fmt.Sprintf("%T", msg), "data", msg)

	switch msg := msg.(type) {
	case openGuestMsg:
		m.guestModel = newGuestModel(msg.uuid, m.conn)
		m.state = guestView

	case startGuestMsg:
		// TODO: Start Guest

	case pauseResumeGuestMsg:
		// TODO: PauseResume Guest

	case shutdownGuestMsg:
		// TODO: Shutdown Guest

	case rebootGuestMsg:
		// TODO: Reboot Guest

	case forceResetGuestMsg:
		// TODO: ForceReset Guest

	case forceOffGuestMsg:
		// TODO: ForceOff Guest

	case saveGuestMsg:
		// TODO: Save Guest

	case cloneGuestMsg:
		// TODO: Clone Guest

	case deleteGuestMsg:
	// TODO: Delete Guest

	case goBackMsg:
		switch m.state {
		case guestView:
			m.state = managerView
		}

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.manager):
			if m.state == managerView {
				break
			}

			m.managerModel = newManagerModel(m.conn)
			m.state = managerView

		case key.Matches(msg, m.keys.network):
			if m.state == networkView {
				break
			}

			m.networkModel = newNetworkModel(m.conn)
			m.state = networkView

		case key.Matches(msg, m.keys.storage):
			if m.state == storageView {
				break
			}

			m.storageModel = newStorageModel(m.conn)
			m.state = storageView

		case key.Matches(msg, m.keys.quit):
			return m, tea.Quit
		}
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

func (m model) View() string {
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

	helpView := m.help.View(m.keys)

	border := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(m.width - 2)
	borderCompensation := 2

	offset := 1 // who knows where this comes from 🤷

	padding := m.height - borderCompensation - offset - strings.Count(mainView, "\n") - strings.Count(helpView, "\n")

	return border.Render(mainView + strings.Repeat("\n", padding) + helpView)
}
