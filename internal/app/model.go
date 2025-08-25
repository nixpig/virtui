package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/common"
	libvirtconn "github.com/nixpig/virtui/internal/libvirt"
)

var _ tea.Model = (*model)(nil)

// ScreenSizeMsg is a message to notify a screen of a size change.
type ScreenSizeMsg struct {
	Width  int
	Height int
}

// Screen is the interface that all application screens must implement.
type Screen interface {
	tea.Model
	ID() string
	Title() string
	Keybindings() []key.Binding
	ScrollKeys() common.ScrollKeyMap
}

// model holds the global state and rendering logic.
type model struct {
	currentScreen         Screen
	screens               map[string]Screen
	width, height         int
	globalKeybindings     KeyMap
	keybindingsViewHeight int
	help                  help.Model
	showFullHelp          bool
	events                chan libvirtconn.DomainEvent
	conn                  libvirtconn.Connection
	service               libvirtconn.Service
}

func listenForEvents(ch chan libvirtconn.DomainEvent) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}

func NewAppModel(
	conn libvirtconn.Connection,
	service libvirtconn.Service,
	screens []Screen,
) *model {
	m := &model{
		globalKeybindings: GlobalKeyMap,
		help:              help.New(),
		showFullHelp:      false,
		screens:           make(map[string]Screen),
		conn:              conn,
		service:           service,
	}

	for _, screen := range screens {
		m.screens[screen.ID()] = screen
	}

	m.currentScreen = m.screens[screens[0].ID()]

	if _, err := m.conn.DomainEventLifecycleRegister(
		func(event libvirtconn.DomainEvent) {
			log.Debug(
				"handling domain event",
				"event",
				event.Event,
				"detail",
				event.Detail,
			)

			m.events <- event
		},
	); err != nil {
		log.Error("failed to register domain event handler", "err", err)
	}

	return m
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		tea.ClearScreen,
		listenForEvents(m.events),
		getDomainsCmd(m.service), // Call getDomainsCmd here
	)
}
