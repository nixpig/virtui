package tui

import (
	"context"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/term"
	"github.com/nixpig/virtui/internal/entity"
	"github.com/nixpig/virtui/internal/service"
	"github.com/nixpig/virtui/internal/tui/common"
	"github.com/nixpig/virtui/internal/tui/guest"
	"github.com/nixpig/virtui/internal/tui/manager"
	"github.com/nixpig/virtui/internal/tui/network"
	"github.com/nixpig/virtui/internal/tui/storage"
	"libvirt.org/go/libvirt"
)

type state int

const (
	managerView state = iota
	guestView
	networkView
	storageView
)

type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}

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
	ctx               context.Context
	state             state
	keys              Keymap
	help              help.Model
	managerModel      tea.Model
	guestModel        tea.Model
	networkModel      tea.Model
	storageModel      tea.Model
	service           service.Service
	connectionDetails entity.ConnectionDetails
	width             int
	height            int
	events            chan *libvirt.DomainEventLifecycle
	eventCallbackID   int
	currentError      error
	initialCmd        tea.Cmd
}

func New(
	svc service.Service,
	ctx context.Context,
) (model, error) {
	connDetails, err := svc.GetConnectionDetails()
	if err != nil {
		log.Error("failed to get connection details", "err", err)
		return model{}, err
	}

	m := model{
		ctx:               ctx,
		state:             managerView,
		keys:              keys,
		help:              help.New(),
		service:           svc,
		connectionDetails: connDetails,
		events:            make(chan *libvirt.DomainEventLifecycle),
	}

	m.width, m.height, err = term.GetSize(os.Stdin.Fd())
	if err != nil {
		log.Warn(
			"failed to get size of terminal",
			"fd",
			os.Stdin.Fd(),
			"err",
			err,
		)
		m.width = 80
		m.height = 24

		m.initialCmd = func() tea.Msg {
			return errMsg{err}
		}
	}

	callbackID, err := svc.RegisterDomainEventCallback(m.events)
	if err != nil {
		log.Error("failed to register domain event callback", "err", err)
		return model{}, err
	}
	m.eventCallbackID = callbackID

	if err := svc.EventLoop(ctx); err != nil {
		log.Error("failed to start event loop", "err", err)
		return model{}, err
	}

	initialManagerModel := manager.NewModel(svc).(manager.Model)
	manModel, _ := initialManagerModel.Update(&libvirt.DomainEventLifecycle{})
	netModel, _ := network.NewModel(svc).Update(&libvirt.DomainEventLifecycle{})
	storModel, _ := storage.NewModel(svc).
		Update(&libvirt.DomainEventLifecycle{})

	m.managerModel = manModel
	m.networkModel = netModel
	m.storageModel = storModel

	return m, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(listenForEvent(m.events), m.initialCmd)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// clear any previous error on receipt of new messages
	m.currentError = nil

	switch msg := msg.(type) {
	case errMsg:
		m.currentError = msg.err
		return m, nil

	case *libvirt.DomainEventLifecycle:
		switch m.state {
		case managerView:
			m.managerModel, cmd = m.managerModel.Update(msg)
			return m, cmd

		}

	case common.OpenGuestMsg:
		m.guestModel = guest.NewModel(msg.UUID, m.service, m.width-2, m.height-3)
		m.state = guestView

	case common.StartGuestMsg:
		if err := m.service.StartDomain(msg.UUID); err != nil {
			log.Error("failed to start domain", "uuid", msg.UUID, "err", err)
			return m, func() tea.Msg {
				return errMsg{err}
			}
		}

	case common.PauseResumeGuestMsg:
		if err := m.service.PauseResumeDomain(msg.UUID); err != nil {
			log.Error("failed to pause/resume domain", "uuid", msg.UUID, "err", err)
			return m, func() tea.Msg {
				return errMsg{err}
			}
		}

	case common.ShutdownGuestMsg:
		if err := m.service.ShutdownDomain(msg.UUID); err != nil {
			log.Error("failed to shutdown domain", "uuid", msg.UUID, "err", err)
			return m, func() tea.Msg {
				return errMsg{err}
			}
		}

	case common.RebootGuestMsg:
		if err := m.service.RebootDomain(msg.UUID); err != nil {
			log.Error("failed to reboot domain", "uuid", msg.UUID, "err", err)
			return m, func() tea.Msg {
				return errMsg{err}
			}
		}

	case common.ForceResetGuestMsg:
		if err := m.service.ForceResetDomain(msg.UUID); err != nil {
			log.Error("failed to reset domain", "uuid", msg.UUID, "err", err)
			return m, func() tea.Msg {
				return errMsg{err}
			}
		}

	case common.ForceOffGuestMsg:
		if err := m.service.ForceOffDomain(msg.UUID); err != nil {
			log.Error("failed to destroy domain", "uuid", msg.UUID, "err", err)
			return m, func() tea.Msg {
				return errMsg{err}
			}
		}

	case common.SaveGuestMsg:
	// TODO: Save Guest

	case common.CloneGuestMsg:
	// TODO: Clone Guest

	case common.DeleteGuestMsg:
	// TODO: Delete Guest (with confirmation)

	case common.GoBackMsg:
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

	cmds = append(cmds, cmd, common.ListenForEvent(m.events))
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var mainView string

	switch m.state {
	case managerView:
		log.Debug("Rendering managerModel")
		mainView = m.managerModel.View()
	case guestView:
		log.Debug("Rendering guestModel")
		mainView = m.guestModel.View()
	case networkView:
		log.Debug("Rendering networkModel")
		mainView = m.networkModel.View()
	case storageView:
		log.Debug("Rendering storageModel")
		mainView = m.storageModel.View()
	default:
		log.Debug("Rendering default managerModel")
		mainView = m.managerModel.View()
	}

	systemInfo := lipgloss.NewStyle().Width(m.width/4).Render(
		labelStyle.Render(" Hostname: ")+
			valueStyle.Render(m.connectionDetails.Hostname)+
			"\n",
		labelStyle.Render("URI: ")+
			valueStyle.Render(m.connectionDetails.URI)+
			"\n",
		labelStyle.Render("Hypervisor: ")+
			valueStyle.Render(
				m.connectionDetails.ConnType+" ("+m.connectionDetails.HvVersion+")",
			)+
			"\n",
		labelStyle.Render("Libvirt: ")+
			valueStyle.Render(m.connectionDetails.LvVersion),
	)

	m.help.Styles.FullKey = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	m.help.Styles.FullDesc = lipgloss.NewStyle().Inherit(valueStyle)

	x := m.managerModel.(manager.Model)
			x.Help().Styles.FullKey = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
			x.Help().Styles.FullDesc = lipgloss.NewStyle().Inherit(valueStyle)

	globalKeys := lipgloss.NewStyle().
		MarginBottom(1).
		Width(m.width / 4).
		Render(m.help.FullHelpView(m.keys.FullHelp()))
	localKeys := lipgloss.NewStyle().
		MarginBottom(1).
		Width(m.width / 2).
		Render(x.Help().FullHelpView(x.Keys().FullHelp()))

	aboveTable := lipgloss.NewStyle().
		MarginTop(1).
		Width(m.width-2).
		MarginLeft(0).
		Border(lipgloss.InnerHalfBlockBorder(), false, true).
		BorderForeground(lipgloss.Color("7")).
		Background(lipgloss.Color("7")).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("0")).
		Render("Guests")

	heading := lipgloss.JoinHorizontal(0, systemInfo, globalKeys, localKeys)

	panel := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, true, true).
		Height(m.height - 2 - 2 - lipgloss.Height(heading)).
		Width(m.width - 2).
		Render(mainView)

	var errorBlock string
	if m.currentError != nil {
		errorBlock = lipgloss.NewStyle().
			Width(m.width - 2).
			PaddingLeft(1).
			PaddingRight(1).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("1")).
			Foreground(lipgloss.Color("1")).
			Render("󰅚  Error: " + m.currentError.Error() + "[C]lose")
	}

	finalView := lipgloss.JoinVertical(1, aboveTable, panel) + "\n" + heading

	if errorBlock != "" {
		finalView = lipgloss.JoinVertical(1, finalView, errorBlock)
	}

	return finalView
}
