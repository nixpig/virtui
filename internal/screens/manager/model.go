package manager

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/common" // Updated import
)

var _ app.Screen = (*model)(nil)

type model struct {
	id       string
	viewport viewport.Model
	width    int
	height   int
	keys     common.ScrollKeyMap
}

func NewManagerScreen() *model {
	return &model{
		id:       "manager",
		viewport: viewport.New(0, 0),
		keys:     common.DefaultScrollKeyMap(),
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// handle scrolling using the common key map
		if key.Matches(msg, m.keys.ScrollUp) {
			m.viewport.ScrollUp(1)
		} else if key.Matches(msg, m.keys.ScrollDown) {
			m.viewport.ScrollUp(1)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	content := "Welcome to virtui!\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"This is a multi-screen Bubble Tea application template.\n\n" +
		"Current screen dimensions: Width = %d, Height = %d\n\n" +
		"Press 'q' or 'ctrl+c' to quit."

	renderedContent := fmt.Sprintf(content, m.width, m.height)

	m.viewport.SetContent(renderedContent)

	style := lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	// calculate content width and height by subtracting frame size
	m.viewport.Width = max(m.width-style.GetHorizontalFrameSize(), 0)
	m.viewport.Height = max(m.height-style.GetVerticalFrameSize(), 0)

	// Render the viewport within the styled box
	return style.Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(m.viewport.View())
}

func (m *model) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

func (m *model) Title() string {
	return "Manager Screen"
}

func (m *model) Keybindings() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "action a")),
		key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "action b")),
		key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "action c")),
		key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "action d")),
	}
}

func (m *model) ScrollKeys() common.ScrollKeyMap { // Updated: Use common.ScrollKeyMap
	return m.keys
}

func (m *model) ID() string {
	return m.id
}
