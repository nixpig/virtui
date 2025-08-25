package storage

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/common"
	"github.com/nixpig/virtui/internal/messages"
)

var _ app.Screen = (*storageScreenModel)(nil)

type storageScreenModel struct {
	id       string
	title    string
	viewport viewport.Model
	width    int
	height   int
	keys     common.ScrollKeyMap
}

// NewStorageScreen returns and initialised storage screen model.
func NewStorageScreen() *storageScreenModel {
	return &storageScreenModel{
		id:       "storage",
		title:    "Storage",
		viewport: viewport.New(0, 0),
		keys:     common.DefaultScrollKeyMap(),
	}
}

func (m *storageScreenModel) Init() tea.Cmd {
	return nil
}

func (m *storageScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case messages.ScreenSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

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

func (m *storageScreenModel) View() string {
	content := "This is the Storage Screen.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"Press '1' for Manager, '2' for Network, 'q' or 'ctrl+c' to quit."

	renderedContent := fmt.Sprintf(content, m.width, m.height)

	m.viewport.SetContent(renderedContent)

	style := lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	m.viewport.Width = max(m.width-style.GetHorizontalFrameSize(), 0)
	m.viewport.Height = max(m.height-style.GetVerticalFrameSize(), 0)

	return style.Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(m.viewport.View())
}

// Title returns the title of the screen.
func (m *storageScreenModel) Title() string {
	return m.title
}

// Keybindings returns screen-specific keybindings.
func (m *storageScreenModel) Keybindings() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("z"), key.WithHelp("z", "action z")),
		key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "action y")),
		key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "action x")),
		key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "action w")),
	}
}

// ScrollKeys returns screen-specific keys used for scrolling.
func (m *storageScreenModel) ScrollKeys() common.ScrollKeyMap {
	return m.keys
}

// ID returns the screen ID.
func (m *storageScreenModel) ID() string {
	return m.id
}
