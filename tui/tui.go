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

func listenForEvent(ch chan *libvirt.DomainEventLifecycle) tea.Cmd {
	return func() tea.Msg {
		e := <-ch
		log.Debug("⏪ reading domain event from channel", "event", e.Event, "detail", e.Detail)
		return e
	}
}

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
	events        chan *libvirt.DomainEventLifecycle
}

func New(conn *libvirt.Connect) model {
	var err error

	m := model{
		state:  managerView,
		keys:   keys,
		help:   help.New(),
		conn:   conn,
		events: make(chan *libvirt.DomainEventLifecycle),
	}

	m.width, m.height, err = term.GetSize(os.Stdin.Fd())
	if err != nil {
		// TODO: need to handle this
		log.Error("get size of terminal", "fd", os.Stdin.Fd(), "err", err)
	}

	if err := libvirt.EventRegisterDefaultImpl(); err != nil {
		log.Error("failed to register default event loop impl", "err", err)
	}

	go func() {
		for {
			if err := libvirt.EventRunDefaultImpl(); err != nil {
				log.Error("failed to run event loop", "err", err)
			}
		}
	}()

	domains, err := conn.ListAllDomains(0)
	if err != nil {
		// TODO: surface error to user?
		log.Debug("list all domains", "err", err)
	}

	if _, err := conn.DomainEventLifecycleRegister(nil, func(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventLifecycle) {
		log.Debug("⏩ writing domain event to channel", "event", event.Event, "detail", event.Detail)
		m.events <- event
	}); err != nil {
		log.Debug("failed to register domain event handler", "err", err)
	}

	defaultModel := newManagerModel(domains, 0)

	m.managerModel = defaultModel

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(listenForEvent(m.events))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	log.Debug("tui received msg", "type", fmt.Sprintf("%T", msg), "data", msg)

	switch msg := msg.(type) {
	case *libvirt.DomainEventLifecycle:
		switch m.state {
		case managerView:
			domains, err := m.conn.ListAllDomains(0)
			if err != nil {
				// TODO: surface error to user?
				log.Debug("list all domains", "err", err)
			}

			mx, _ := m.managerModel.(managerModel)
			m.managerModel = newManagerModel(domains, mx.table.Cursor())
			// TODO: what about persisting active selection in ui instead of it jumping to top

			// TODO: other model views

		}

	case openGuestMsg:
		m.guestModel = newGuestModel(msg.uuid, m.conn, m.width, m.height)
		m.state = guestView

	case startGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Error("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		if err := d.Create(); err != nil {
			// TODO: surface error to user?
			log.Error("failed to create domain", "uuid", msg.uuid, "err", err)
		}

	case pauseResumeGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Error("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		s, _, _ := d.GetState()
		if s == libvirt.DOMAIN_PAUSED {
			if err := d.Resume(); err != nil {
				log.Error("failed to resume domain", "uuid", msg.uuid, "err", err)
			}
		} else if s == libvirt.DOMAIN_RUNNING {
			if err := d.Suspend(); err != nil {
				log.Error("failed to pause domain", "uuid", msg.uuid, "err", err)
			}
		}

	case shutdownGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Error("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		s, _, _ := d.GetState()
		if s != libvirt.DOMAIN_SHUTOFF {
			if err := d.Shutdown(); err != nil {
				log.Error("failed to shutdown domain", "uuid", msg.uuid, "err", err)
			}
		}

	case rebootGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Error("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		s, _, _ := d.GetState()
		if s == libvirt.DOMAIN_RUNNING {
			if err := d.Reboot(0); err != nil {
				log.Error("failed to reboot domain", "uuid", msg.uuid, "err", err)
			}
		}

	case forceResetGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Error("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		if err := d.Reset(0); err != nil {
			log.Error("failed to reset domain", "uuid", msg.uuid, "err", err)
		}

	case forceOffGuestMsg:
		d, err := m.conn.LookupDomainByUUIDString(msg.uuid)
		if err != nil {
			// TODO: surface error to user?
			log.Error("failed to lookup domain", "uuid", msg.uuid, "err", err)
		}
		if err := d.Destroy(); err != nil {
			log.Error("failed to destroy domain", "uuid", msg.uuid, "err", err)
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

			domains, err := m.conn.ListAllDomains(0)
			if err != nil {
				// TODO: surface error to user?
				log.Debug("list all domains", "err", err)
			}

			mx, _ := m.managerModel.(managerModel)
			m.managerModel = newManagerModel(domains, mx.table.Cursor())
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

	cmds = append(cmds, cmd, listenForEvent(m.events))
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
