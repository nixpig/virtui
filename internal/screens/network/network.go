package network

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/common"
	"github.com/nixpig/virtui/internal/libvirtui"
	"github.com/nixpig/virtui/internal/messages"
)

var _ app.Screen = (*networkScreenModel)(nil)

type networkScreenModel struct {
	id       string
	title    string
	viewport viewport.Model
	width    int
	height   int
	keys     common.ScrollKeyMap
	networks []libvirtui.Network
}

// NewNetworkScreen returns an initialised network screen model.
func NewNetworkScreen() *networkScreenModel {
	return &networkScreenModel{
		id:       "network",
		title:    "Networks",
		viewport: viewport.New(0, 0),
		keys:     common.DefaultScrollKeyMap(),
		networks: []libvirtui.Network{},
	}
}

func (m *networkScreenModel) Init() tea.Cmd {
	return nil
}

func (m *networkScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case messages.ScreenSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case messages.NetworksMsg:
		m.networks = msg.Networks

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.viewport.ScrollUp(1)
		case key.Matches(msg, m.keys.Down):
			m.viewport.ScrollDown(1)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *networkScreenModel) View() string {
	var sb strings.Builder
	if len(m.networks) == 0 {
		sb.WriteString("No networks found.")
	} else {
		for _, network := range m.networks {
			sb.WriteString("Name: " + network.Name() + "\n")
			sb.WriteString("UUID: " + network.UUID() + "\n")
			sb.WriteString("Bridge: " + network.Bridge() + "\n")
			sb.WriteString(fmt.Sprintf("Active: %t\n", network.Active()))
			sb.WriteString("\n")
		}
	}

	renderedContent := sb.String()

	m.viewport.SetContent(renderedContent)

	style := lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	m.viewport.Width = max(m.width-style.GetHorizontalFrameSize(), 0)
	m.viewport.Height = max(m.height-style.GetVerticalFrameSize(), 0)

	return style.Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(m.viewport.View())
}

// Title returns the title of the screen.
func (m *networkScreenModel) Title() string {
	return m.title
}

// Keybindings returns screen-specific keybindings.
func (m *networkScreenModel) Keybindings() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("z"), key.WithHelp("z", "action z")),
		key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "action y")),
		key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "action x")),
		key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "action w")),
	}
}

// ScrollKeys returns the the screen-specific keybindings for scrolling.
func (m *networkScreenModel) ScrollKeys() common.ScrollKeyMap {
	return m.keys
}

// ID returns the screen ID used to reference the screen in other areas of app.
func (m *networkScreenModel) ID() string {
	return m.id
}
