package app

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/common"
	libvirtconn "github.com/nixpig/virtui/internal/libvirt/conn"
	"libvirt.org/go/libvirt"
)

var _ tea.Model = (*model)(nil)

// Screen is the interface that all application screens must implement.
type Screen interface {
	tea.Model
	ID() string
	SetDimensions(width, height int)
	Title() string
	Keybindings() []key.Binding
	ScrollKeys() common.ScrollKeyMap
}

// model holds the global state and rendering logic.
type model struct {
	currentScreen         Screen
	screens               map[string]Screen
	width, height         int
	globalKeybindings     []key.Binding
	keybindingsViewHeight int
	screenKeyMap          map[string]string
	help                  help.Model
	showFullHelp          bool
	events                chan *libvirt.DomainEventLifecycle
	libvirt               libvirtconn.LibvirtConnection
}

func listenForEvents(ch chan *libvirt.DomainEventLifecycle) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}

func NewAppModel(conn libvirtconn.LibvirtConnection, screens []Screen) *model {
	m := &model{
		globalKeybindings: globalKeyBindings,
		help:              help.New(),
		showFullHelp:      false,
		screens:           make(map[string]Screen),
		libvirt:           conn,
	}

	for _, screen := range screens {
		m.screens[screen.ID()] = screen
	}

	m.currentScreen = m.screens[screens[0].ID()]

	m.screenKeyMap = map[string]string{
		"1": "manager",
		"2": "storage",
		"3": "network",
	}

	if _, err := m.libvirt.DomainEventLifecycleRegister(
		nil,
		func(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventLifecycle) {
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
	return tea.Batch(tea.ClearScreen, listenForEvents(m.events))
}
