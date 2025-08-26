package manager

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/exp/charmtone"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/common"
	"github.com/nixpig/virtui/internal/icons"
	"github.com/nixpig/virtui/internal/libvirtui"
	"github.com/nixpig/virtui/internal/messages"
)

var _ app.Screen = (*managerScreenModel)(nil)

type managerScreenModel struct {
	id       string
	title    string
	viewport viewport.Model
	width    int
	height   int
	keys     common.ScrollKeyMap
	table    table.Model
}

// NewManagerScreen returns an initialised manager screen model.
func NewManagerScreen() *managerScreenModel {
	columns := []table.Column{
		// TODO: make this column dynamic so the table takes the full width
		{Title: "Name", Width: 20},
		{Title: "State", Width: 10},
		{Title: "Memory", Width: 10},
		{Title: "CPU", Width: 5},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{}),
		table.WithFocused(true),
		// TODO: height needs to take the available height
		table.WithHeight(10),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color(charmtone.Zest.Hex())).
		Background(lipgloss.Color(charmtone.Dolly.Hex())).
		Bold(false)

	t.SetStyles(s)

	return &managerScreenModel{
		id:       "manager",
		title:    "Dashboard",
		viewport: viewport.New(0, 0),
		keys:     common.DefaultScrollKeyMap(),
		table:    t,
	}
}

func (m *managerScreenModel) Init() tea.Cmd {
	return nil
}

func (m *managerScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case messages.ScreenSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		style := lipgloss.NewStyle().Border(lipgloss.HiddenBorder())
		m.table.SetWidth(m.width - style.GetHorizontalFrameSize())
		m.table.SetHeight(m.height - style.GetVerticalFrameSize())

		fixedColumnsWidth := 0
		for _, col := range m.table.Columns()[1:] {
			fixedColumnsWidth += col.Width + 2
		}

		nameColumnWidth := m.table.Width() - fixedColumnsWidth - style.GetHorizontalFrameSize()
		m.table.Columns()[0].Width = nameColumnWidth

	case messages.DomainsMsg:
		rows := make([]table.Row, len(msg.Domains))

		for i, domain := range msg.Domains {
			var icon string
			domainState, _, _ := domain.GetState()

			switch libvirtui.DomainState(domainState) {
			case libvirtui.DomainStateRunning:
				icon = icons.Icons.VM.Running
			case libvirtui.DomainStateBlocked:
				icon = icons.Icons.VM.Blocked
			case libvirtui.DomainStatePaused:
				icon = icons.Icons.VM.Paused
			case libvirtui.DomainStateShutdown, libvirtui.DomainStateShutoff:
				icon = icons.Icons.VM.Off
			default:
				icon = icons.Icons.VM.Off
			}

			rows[i] = table.Row{fmt.Sprintf("%s  %s", icon, domain.Name()), domain.State(), fmt.Sprintf("%dMB", domain.Memory()/1024), fmt.Sprintf("%d", domain.VCPU())}
		}

		m.table.SetRows(rows)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.table.MoveUp(1)
		case key.Matches(msg, m.keys.Down):
			m.table.MoveDown(1)
		}
	}

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *managerScreenModel) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
	return style.Render(m.table.View())
}

// Title returns the title of the screen.
func (m *managerScreenModel) Title() string {
	return m.title
}

// Keybindings returns screen-specific keybindings.
func (m *managerScreenModel) Keybindings() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "action a")),
		key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "action b")),
		key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "action c")),
		key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "action d")),
	}
}

// ScrollKeys returns the screen-specific keybindings for scrolling.
func (m *managerScreenModel) ScrollKeys() common.ScrollKeyMap {
	return m.keys
}

// ID returns the screen ID.
func (m *managerScreenModel) ID() string {
	return m.id
}
