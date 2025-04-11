package manager

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/nixpig/virtui/keys"
	"github.com/nixpig/virtui/vm"
	"libvirt.org/go/libvirt"
)

func waitForActivity(event chan vm.Event) tea.Cmd {
	return func() tea.Msg {
		return <-event
	}
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	event  chan vm.Event
	conn   *libvirt.Connect
	vms    []vm.VM
	table  table.Model
	keys   keys.Keymap
	help   help.Model
	width  int
	height int
}

func cb(event chan vm.Event) libvirt.DomainEventLifecycleCallback {
	return func(
		conn *libvirt.Connect,
		domain *libvirt.Domain,
		lvEvent *libvirt.DomainEventLifecycle,
	) {
		event <- vm.Event{
			Event: lvEvent.Event,
		}
	}
}

func InitModel(conn *libvirt.Connect) Model {
	m := Model{
		conn:  conn,
		event: make(chan vm.Event),
	}

	vms := vm.GetAll(conn)
	m.vms = vms

	if _, err := conn.DomainEventLifecycleRegister(
		nil,
		cb(m.event),
	); err != nil {
		fmt.Println("ERR: ", err)
		os.Exit(1)
	}

	rows := vmsToRows(vms)

	t := buildTableModel(rows)

	m.table = t
	m.help = help.New()
	m.keys = keys.Keys

	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(waitForActivity(m.event))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	v := m.vms[m.table.Cursor()]

	switch msg := msg.(type) {
	// libvirt event
	case vm.Event:

		// we can do better than just re-fetching all the domains on every event
		// probably only get details for the specific domain and update it
		// but for now this will do
		m.vms = vm.GetAll(m.conn)
		rows := vmsToRows(m.vms)
		m.table.SetRows(rows)

		// ---

		return m, tea.Batch(waitForActivity(m.event))

	// resize
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// user input
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Focus):
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		// Start
		case key.Matches(msg, m.keys.Run):
			if err := v.Run(); err != nil {
				// TODO: present err and log
			}
			return m, nil

		// Pause/Resume
		case key.Matches(msg, m.keys.PauseResume):
			if err := v.PauseResume(); err != nil {
				// TODO: present err and log
			}
			return m, nil

		// Shutdown
		case key.Matches(msg, m.keys.Shutdown):
			if err := v.Shutdown(); err != nil {
				// TODO: present err and log
			}
			return m, nil

		// Reboot
		case key.Matches(msg, m.keys.Reboot):
			if err := v.Reboot(); err != nil {
				// TODO: present err and log
			}
			return m, nil

		// Reset
		case key.Matches(msg, m.keys.ForceReset):
			if err := v.ForceReset(); err != nil {
				// TODO: present err and log
			}
			return m, nil

		// Off
		case key.Matches(msg, m.keys.ForceOff):
			if err := v.ForceOff(); err != nil {
				fmt.Println("ERR: ", err)
				// TODO: present err and log
			}
			return m, nil

		// Save/Restore
		case key.Matches(msg, m.keys.SaveRestore):
			if err := v.SaveRestore(); err != nil {
				// TODO: present err and log
			}
			return m, nil

		// Delete
		case key.Matches(msg, m.keys.Delete):
			if err := v.Delete(); err != nil {
				fmt.Println("ERR: ", err)
				// TODO: present err and log
			}
			return m, nil

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return baseStyle.Render(m.table.View()+"\n"+m.help.View(m.keys)) + "\n"
}

func buildTableModel(rows []table.Row) table.Model {
	w, _, err := term.GetSize(0)
	if err != nil {
		fmt.Println("failed to get term size: ", err)
	}

	idWidth := 3
	stateWidth := 8
	nameWidth := w - idWidth - stateWidth - 4 - 4 - 2 - 4 - 4 - 4 - 7

	columns := []table.Column{
		{Title: "ID", Width: idWidth},
		{Title: "Name", Width: nameWidth},
		{Title: "State", Width: stateWidth},
		{Title: "CPU", Width: 3},
		{Title: "Mem", Width: 3},
		{Title: "Blk", Width: 3},
		{Title: "Net", Width: 3},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	return t
}

func vmsToRows(vms []vm.VM) []table.Row {
	rows := make([]table.Row, len(vms))

	for i, v := range vms {
		name := v.GetPresentableName()
		state := v.GetPresentableState()
		id := v.GetPresentableID()
		rows[i] = table.Row{id, name, state, "⣾⣷⣷", "⣷⣾⣷", "▄ ▆", "▃"}
	}

	return rows
}
