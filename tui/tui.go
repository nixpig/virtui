package tui

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/term"
	"github.com/nixpig/virtui/tui/mappers"
	"libvirt.org/go/libvirt"
)

type state int

const (
	managerView state = iota
	guestView
	networkView
	storageView
)

var (
	labelStyle = lipgloss.NewStyle().Bold(true)
	valueStyle = lipgloss.NewStyle().Faint(true)

	headingStyle = lipgloss.NewStyle().
			Underline(true).
			Bold(true)
)

func listenForEvent(ch chan *libvirt.DomainEventLifecycle) tea.Cmd {
	return func() tea.Msg {
		e := <-ch
		return e
	}
}

type model struct {
	state             state
	keys              keymap
	help              help.Model
	managerModel      tea.Model
	guestModel        tea.Model
	networkModel      tea.Model
	storageModel      tea.Model
	conn              *libvirt.Connect
	connectionDetails *connectionDetails
	width             int
	height            int
	events            chan *libvirt.DomainEventLifecycle
}

type connectionDetails struct {
	hostname  string
	uri       string
	connType  string
	hvVersion string
	lvVersion string
}

func New(conn *libvirt.Connect) model {
	var err error

	hostname, _ := conn.GetHostname()
	lvVersion, _ := conn.GetLibVersion()
	hvVersion, _ := conn.GetVersion()
	connectionType, _ := conn.GetType()
	connURI, _ := conn.GetURI()

	connDetails := &connectionDetails{
		uri:       connURI,
		hostname:  hostname,
		lvVersion: mappers.Version(lvVersion),
		hvVersion: mappers.Version(hvVersion),
		connType:  connectionType,
	}

	m := model{
		state:             managerView,
		keys:              keys,
		help:              help.New(),
		conn:              conn,
		connectionDetails: connDetails,
		events:            make(chan *libvirt.DomainEventLifecycle),
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
			// TODO: pass context and close event loop cleanly on exit and unregister handlers
			if err := libvirt.EventRunDefaultImpl(); err != nil {
				log.Error("failed to run event loop", "err", err)
			}
		}
	}()

	if _, err := conn.DomainEventLifecycleRegister(nil, func(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventLifecycle) {
		log.Debug("handing domain event", "event", event.Event, "detail", event.Detail, "data", event)
		m.events <- event
	}); err != nil {
		// TODO: surface error to user?
		log.Debug("failed to register domain event handler", "err", err)
	}

	manModel, _ := newManagerModel(conn).Update(&libvirt.DomainEventLifecycle{})
	netModel, _ := newNetworkModel(conn).Update(&libvirt.DomainEventLifecycle{})
	storModel, _ := newStorageModel(conn).Update(&libvirt.DomainEventLifecycle{})

	m.managerModel = manModel
	m.networkModel = netModel
	m.storageModel = storModel

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(listenForEvent(m.events))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case *libvirt.DomainEventLifecycle:
		switch m.state {
		case managerView:
			m.managerModel, cmd = m.managerModel.Update(msg)
			return m, cmd

		}

	case openGuestMsg:
		m.guestModel = newGuestModel(msg.uuid, m.conn, m.width-2, m.height-3)
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
		log.Debug("window size message in tui", "width", msg.Width, "height", msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.manager):
			if m.state == managerView {
				break
			}

			m.managerModel, cmd = m.managerModel.Update(msg)
			m.state = managerView

		case key.Matches(msg, m.keys.network):
			if m.state == networkView {
				break
			}

			m.networkModel, cmd = m.networkModel.Update(msg)
			m.state = networkView

		case key.Matches(msg, m.keys.storage):
			if m.state == storageView {
				break
			}

			m.storageModel, cmd = m.storageModel.Update(msg)
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

	systemInfo := lipgloss.NewStyle().Width(m.width/4).Render(
		labelStyle.Render(" Hostname: ")+valueStyle.Render(m.connectionDetails.hostname)+"\n",
		labelStyle.Render("URI: ")+valueStyle.Render(m.connectionDetails.uri)+"\n",
		labelStyle.Render("Hypervisor: ")+valueStyle.Render(m.connectionDetails.connType+" ("+m.connectionDetails.hvVersion+")")+"\n",
		labelStyle.Render("Libvirt: ")+valueStyle.Render(m.connectionDetails.lvVersion),
	)

	m.help.Styles.FullKey = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	m.help.Styles.FullDesc = lipgloss.NewStyle().Inherit(valueStyle)

	x := m.managerModel.(managerModel)
	x.help.Styles.FullKey = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
	x.help.Styles.FullDesc = lipgloss.NewStyle().Inherit(valueStyle)

	globalKeys := lipgloss.NewStyle().MarginBottom(1).Width(m.width / 4).Render(m.help.FullHelpView(m.keys.FullHelp()))
	localKeys := lipgloss.NewStyle().MarginBottom(1).Width(m.width / 2).Render(x.help.FullHelpView(x.keys.FullHelp()))

	aboveTable := lipgloss.NewStyle().MarginTop(1).Width(m.width-2).MarginLeft(0).Border(lipgloss.InnerHalfBlockBorder(), false, true).BorderForeground(lipgloss.Color("7")).Background(lipgloss.Color("7")).Align(lipgloss.Center).Foreground(lipgloss.Color("0")).Render("Guests")

	heading := lipgloss.JoinHorizontal(0, systemInfo, globalKeys, localKeys)

	// errorBlock := lipgloss.NewStyle().Width(m.width - 2).PaddingLeft(1).PaddingRight(1).Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("1")).Foreground(lipgloss.Color("1")).Render("󰅚  Error: some error has occurred, please try again. [C]lose")
	// successBlock := lipgloss.NewStyle().Width(m.width - 2).PaddingLeft(1).PaddingRight(1).Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("2")).Foreground(lipgloss.Color("2")).Render("󰗡  Success: something happened successfully! [O]kay")
	// warningBlock := lipgloss.NewStyle().Width(m.width - 2).PaddingLeft(1).PaddingRight(1).Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("3")).Foreground(lipgloss.Color("3")).Render("󰗖  Warning: something happened that might be an issue! [O]kay")

	panel := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true, true, true).
		Height(m.height - 2 - 2 - lipgloss.Height(heading)).
		Width(m.width - 2).
		Render(mainView)

	// return heading + "\n" + panel
	return lipgloss.JoinVertical(1, aboveTable, panel) + "\n" + heading
}
