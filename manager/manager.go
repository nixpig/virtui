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

type responseMsg struct {
	vm    *vm.VM
	event libvirt.DomainEventType
}

func waitForActivity(sub chan responseMsg) tea.Cmd {
	return func() tea.Msg {
		return <-sub
	}
}

var rows = []table.Row{}
var domains []libvirt.Domain

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	sub   chan responseMsg
	conn  *libvirt.Connect
	vms   []vm.VM
	table table.Model
	keys  keys.Keymap
	help  help.Model
}

func cb(sub chan responseMsg, v *vm.VM) libvirt.DomainEventLifecycleCallback {
	return func(
		conn *libvirt.Connect,
		domain *libvirt.Domain,
		event *libvirt.DomainEventLifecycle,
	) {
		sub <- responseMsg{
			vm:    v,
			event: event.Event,
		}
	}
}

func InitModel(conn *libvirt.Connect) Model {
	m := Model{
		conn: conn,
		sub:  make(chan responseMsg),
	}

	vms := vm.GetAll(conn)
	m.vms = vms

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

	for _, v := range m.vms {
		if _, err := conn.DomainEventLifecycleRegister(
			v.Domain,
			cb(m.sub, &v),
		); err != nil {
			fmt.Println("ERR: ", err)
			os.Exit(1)
		}

		name := v.GetPresentableName()
		state := v.GetPresentableState()
		id := v.GetPresentableID()

		// https://en.wikipedia.org/wiki/Braille_Patterns
		// https://en.wikipedia.org/wiki/Block_Elements
		rows = append(
			rows,
			table.Row{
				id,
				name,
				state,
				"⣾⣷⣷",
				"⣷⣾⣤",
				"▄ ▆",
				"▇ ▃",
			})

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

	m.table = t
	m.help = help.New()
	m.keys = keys.Keys

	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(waitForActivity(m.sub))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	v := m.vms[m.table.Cursor()]

	switch msg := msg.(type) {
	case responseMsg:
		id, _ := msg.vm.GetID()
		d, _ := m.conn.LookupDomainById(uint32(id))
		m.vms[m.table.Cursor()] = *vm.FromDomain(d)

		// ---

		existingRows := m.table.Rows()
		existingRows[m.table.Cursor()] = table.Row{
			msg.vm.GetPresentableID(),
			msg.vm.GetPresentableName(),
			msg.vm.GetPresentableState(),
			"", "", "", "",
		}

		m.table.SetRows(existingRows)

		// ---

		return m, tea.Batch(waitForActivity(m.sub))

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
			// return m, nil

		// Pause/Resume
		case key.Matches(msg, m.keys.PauseResume):
			if err := v.PauseResume(); err != nil {
				// TODO: present err and log
			}
			// return m, nil

		// Shutdown
		case key.Matches(msg, m.keys.Shutdown):
			if err := v.Shutdown(); err != nil {
				// TODO: present err and log
			}
			// return m, nil

		// Reboot
		case key.Matches(msg, m.keys.Reboot):
			if err := v.Reboot(); err != nil {
				// TODO: present err and log
			}
			// return m, nil

		// Reset
		case key.Matches(msg, m.keys.ForceReset):
			if err := v.ForceReset(); err != nil {
				// TODO: present err and log
			}
			// return m, nil

		// Off
		case key.Matches(msg, m.keys.ForceOff):
			if err := v.ForceOff(); err != nil {
				fmt.Println("ERR: ", err)
				// TODO: present err and log
			}
			// return m, nil

		// Save/Restore
		case key.Matches(msg, m.keys.SaveRestore):
			if err := v.SaveRestore(); err != nil {
				// TODO: present err and log
			}
			// return m, nil

		// Delete
		case key.Matches(msg, m.keys.Delete):
			if err := v.Delete(); err != nil {
				fmt.Println("ERR: ", err)
				// TODO: present err and log
			}
			// TODO: remove from array??
			// return m, nil

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		}
	}

	if err := v.Update(m.conn); err != nil {
		fmt.Println("ERR: ", err)
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return baseStyle.Render(m.table.View()+"\n"+m.help.View(m.keys)) + "\n"
}

