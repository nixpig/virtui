package manager

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/common"
	"github.com/nixpig/virtui/internal/messages" // New import
)

var _ app.Screen = (*model)(nil)

type model struct {
	id       string
	viewport viewport.Model
	width    int
	height   int
	keys     common.ScrollKeyMap
	table    table.Model // New field for the table
}

func NewManagerScreen() *model {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "State", Width: 10},
		{Title: "Memory", Width: 10},
		{Title: "CPU", Width: 5},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}), // Initially empty
		table.WithFocused(true),
		table.WithHeight(10), // Placeholder height
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return &model{
		id:       "manager",
		viewport: viewport.New(0, 0),
		keys:     common.DefaultScrollKeyMap(),
		table:    t,
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
		m.table.SetWidth(m.width)
		m.table.SetHeight(m.height - 2) // Adjust for header/footer
	case messages.DomainsMsg: // New case for DomainsMsg
		rows := make([]table.Row, len(msg))
		for i, domain := range msg {
			rows[i] = table.Row{domain.Name, domain.State, fmt.Sprintf("%dMB", domain.Memory/1024), fmt.Sprintf("%d", domain.VCPU)}
		}
		m.table.SetRows(rows)
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.ScrollUp) {
			m.table.MoveUp(1) // Use MoveUp
		} else if key.Matches(msg, m.keys.ScrollDown) {
			m.table.MoveDown(1) // Use MoveDown
		}
	}

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *model) View() string {
	return m.table.View()
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

func (m *model) ScrollKeys() common.ScrollKeyMap {
	return m.keys
}

func (m *model) ID() string {
	return m.id
}
