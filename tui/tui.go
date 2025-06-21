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
		// TODO: need to handle this
		log.Debug("get size of terminal", "fd", os.Stdin.Fd(), "err", err)
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
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		if err := d.Create(); err != nil {
			// TODO: surface error to user?
			log.Debug("failed to create domain", "uuid", msg.uuid, "err", err)
		}

	case pauseResumeGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		s, _, _ := d.GetState()
		if s == libvirt.DOMAIN_PAUSED {
			if err := d.Resume(); err != nil {
				log.Debug("failed to resume domain", "uuid", msg.uuid, "err", err)
			}
		} else if s == libvirt.DOMAIN_RUNNING {
			if err := d.Suspend(); err != nil {
				log.Debug("failed to pause domain", "uuid", msg.uuid, "err", err)
			}
		}

	case shutdownGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		s, _, _ := d.GetState()
		if s != libvirt.DOMAIN_SHUTOFF {
			if err := d.Shutdown(); err != nil {
				log.Debug("failed to shutdown domain", "uuid", msg.uuid, "err", err)
			}
		}

	case rebootGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		s, _, _ := d.GetState()
		if s == libvirt.DOMAIN_RUNNING {
			if err := d.Reboot(0); err != nil {
				log.Debug("failed to reboot domain", "uuid", msg.uuid, "err", err)
			}
		}

	case forceResetGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		if err := d.Reset(0); err != nil {
			log.Debug("failed to reset domain", "uuid", msg.uuid, "err", err)
		}

	case forceOffGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Debug("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		if err := d.Destroy(); err != nil {
			log.Debug("failed to destroy domain", "uuid", msg.uuid, "err", err)
		}

	case saveGuestMsg:
	// TODO: Save Guest
	// d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
	// if err != nil {
	// 	// TODO: surface error to user?
	// 	log.Debug("failed to lookup domain", "uuid", msg.uuid, "err", err)
	// }
	// if err := d.Save(/* SOME FILE TO SAVE TO */); err != nil {
	// 	log.Debug("failed to destroy domain", "uuid", msg.uuid, "err", err)
	// }

	case cloneGuestMsg:
		// TODO: Clone Guest

	case deleteGuestMsg:
	// TODO: Delete Guest (with confirmation)

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
