package storage

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/common"
)

var _ app.Screen = (*model)(nil)

type model struct {
	id       string
	viewport viewport.Model
	width    int
	height   int
	keys     common.ScrollKeyMap
}

func NewStorageScreen() *model {
	return &model{
		id:       "storage",
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
	case app.ScreenSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.ScrollUp) {
			m.viewport.ScrollUp(1)
		} else if key.Matches(msg, m.keys.ScrollDown) {
			m.viewport.ScrollDown(1)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
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

func (m *model) Title() string {
	return "Storage Screen"
}

func (m *model) Keybindings() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("z"), key.WithHelp("z", "action z")),
		key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "action y")),
		key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "action x")),
		key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "action w")),
	}
}

func (m *model) ScrollKeys() common.ScrollKeyMap {
	return m.keys
}

func (m *model) ID() string {
	return m.id
}
