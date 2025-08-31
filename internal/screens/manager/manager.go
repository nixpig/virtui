package manager

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/charmtone"
	"github.com/nixpig/virtui/internal/icons"
	"github.com/nixpig/virtui/internal/libvirtui"
	"github.com/nixpig/virtui/internal/messages"
	"github.com/nixpig/virtui/internal/screen"
)

var _ screen.Screen = (*managerScreenModel)(nil)

type managerScreenModel struct {
	id       string
	title    string
	viewport viewport.Model
	width    int
	height   int
	keys     KeyMap
	table    table.Model
	domains  []libvirtui.Domain
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
		Foreground(lipgloss.Color(charmtone.Ash.Hex())).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(charmtone.Charcoal.Hex())).
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
		keys:     defaultKeyMap,
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
		m.domains = msg.Domains
		rows := make([]table.Row, len(msg.Domains))

		for i, domain := range msg.Domains {
			var icon string
			domainState, _, err := domain.GetState()
			log.Error("failed to get domain state", "name", domain.Name(), "err", err)

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

			rows[i] = table.Row{
				fmt.Sprintf("%s  %s", icon, domain.Name()),
				domain.State(),
				fmt.Sprintf("%dMB", domain.Memory()/1024),
				fmt.Sprintf("%d", domain.VCPU()),
			}
		}

		m.table.SetRows(rows)

	case tea.KeyMsg:
		return m.handleKeypress(msg)
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
		BorderForeground(lipgloss.Color(charmtone.Charcoal.Hex()))
	return style.Render(m.table.View())
}

// Title returns the title of the screen.
func (m *managerScreenModel) Title() string {
	return m.title
}

func (m *managerScreenModel) HelpKeys() [][]key.Binding {
	return m.keys.FullHelp()
}

// ID returns the screen ID.
func (m *managerScreenModel) ID() string {
	return m.id
}

func (m *managerScreenModel) handleKeypress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	selectedDomainIndex := m.table.Cursor()
	selectedDomain := m.domains[selectedDomainIndex]
	domainUUID, err := selectedDomain.GetUUIDString()
	if err != nil {
		log.Error("failed to get domain UUID", "err", err)
		return m, nil
	}

	switch {
	case key.Matches(msg, m.keys.Start):
		return m, messages.NewDomainActionWithFunc(domainUUID, func(service libvirtui.Service, uuid string) error {
			return service.DomainStart(uuid)
		})
	case key.Matches(msg, m.keys.PauseResume):
		return m, messages.NewDomainActionWithFunc(domainUUID, func(service libvirtui.Service, uuid string) error {
			return service.ToggleDomainState(uuid)
		})
	case key.Matches(msg, m.keys.Shutdown):
		return m, messages.NewDomainActionWithFunc(domainUUID, func(service libvirtui.Service, uuid string) error {
			return service.DomainShutdown(uuid)
		})
	case key.Matches(msg, m.keys.Reboot):
		return m, messages.NewDomainActionWithFunc(domainUUID, func(service libvirtui.Service, uuid string) error {
			return service.DomainReboot(uuid)
		})
	case key.Matches(msg, m.keys.Reset):
		return m, messages.NewDomainActionWithFunc(domainUUID, func(service libvirtui.Service, uuid string) error {
			return service.ResetDomain(uuid)
		})
	case key.Matches(msg, m.keys.ForceOff):
		return m, messages.NewDomainActionWithFunc(domainUUID, func(service libvirtui.Service, uuid string) error {
			return service.ForceOffDomain(uuid)
		})
	case key.Matches(msg, m.keys.Save):
		// TODO: implement save
		return m, messages.NewDomainAction("save", domainUUID, true)
	case key.Matches(msg, m.keys.Clone):
		// TODO: implement clone
		return m, messages.NewDomainAction("clone", domainUUID, true)
	case key.Matches(msg, m.keys.Delete):
		// TODO: implement delete
		return m, messages.NewDomainAction("delete", domainUUID, true)
	case key.Matches(msg, m.keys.Open):
		// TODO: implement open
		return m, messages.NewDomainAction("open", domainUUID, true)
	case key.Matches(msg, m.keys.KeyUp):
		m.table.MoveUp(1)
	case key.Matches(msg, m.keys.KeyDown):
		m.table.MoveDown(1)
	}
	return m, nil
}
