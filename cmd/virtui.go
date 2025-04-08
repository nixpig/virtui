package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"libvirt.org/go/libvirt"
)

const libvirtURI = "qemu:///system"

var rows = []table.Row{}
var domains []libvirt.Domain

var domainState = map[libvirt.DomainState]string{
	libvirt.DOMAIN_NOSTATE:     "None",
	libvirt.DOMAIN_RUNNING:     "Running",
	libvirt.DOMAIN_BLOCKED:     "Blocked",
	libvirt.DOMAIN_PAUSED:      "Paused",
	libvirt.DOMAIN_SHUTDOWN:    "Shutdown",
	libvirt.DOMAIN_CRASHED:     "Crashed",
	libvirt.DOMAIN_PMSUSPENDED: "Suspended",
	libvirt.DOMAIN_SHUTOFF:     "Shutoff",
}

type keymap struct {
	Up          key.Binding
	Down        key.Binding
	Open        key.Binding
	Start       key.Binding
	PauseResume key.Binding
	Shutdown    key.Binding
	Reboot      key.Binding
	Reset       key.Binding
	PowerOff    key.Binding
	SaveRestore key.Binding
	Migrate     key.Binding
	Delete      key.Binding
	Clone       key.Binding
	New         key.Binding
	Help        key.Binding
	Quit        key.Binding
	Focus       key.Binding
}

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Open, k.Help, k.Quit},
		{k.Start, k.PauseResume, k.Shutdown, k.Reboot, k.Reset, k.PowerOff},
		{k.New, k.SaveRestore, k.Migrate, k.Clone, k.Delete},
	}
}

var keys = keymap{
	Up: key.NewBinding(
		key.WithKeys("up", "k", "K"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j", "J"),
		key.WithHelp("↓/j", "down"),
	),
	Open: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "open"),
	),
	Start: key.NewBinding(
		key.WithKeys("t", "T"),
		key.WithHelp("t", "start"),
	),
	PauseResume: key.NewBinding(
		key.WithKeys("p", "P"),
		key.WithHelp("p", "pause/resume"),
	),
	Shutdown: key.NewBinding(
		key.WithKeys("s", "S"),
		key.WithHelp("s", "shutdown"),
	),
	Reboot: key.NewBinding(
		key.WithKeys("r", "R"),
		key.WithHelp("r", "reboot"),
	),
	Reset: key.NewBinding(
		key.WithKeys("o", "O"),
		key.WithHelp("o", "reset"),
	),
	PowerOff: key.NewBinding(
		key.WithKeys("f", "F"),
		key.WithHelp("f", "poweroff"),
	),
	SaveRestore: key.NewBinding(
		key.WithKeys("v", "V"),
		key.WithHelp("v", "save/restore"),
	),
	Migrate: key.NewBinding(
		key.WithKeys("m", "M"),
		key.WithHelp("m", "migrate"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d", "D"),
		key.WithHelp("d", "delete"),
	),
	Clone: key.NewBinding(
		key.WithKeys("c", "C"),
		key.WithHelp("c", "clone"),
	),
	Help: key.NewBinding(
		key.WithKeys("h", "H", "?"),
		key.WithHelp("h", "help"),
	),
	New: key.NewBinding(
		key.WithKeys("n", "N"),
		key.WithHelp("n", "new"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "Q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Focus: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "focus"),
	),
}

func fmtState(d *libvirt.Domain) string {
	s, _, err := d.GetState()
	if err != nil {
		return "-"
	}
	state, ok := domainState[s]
	if !ok {
		return "-"
	}
	return state
}

func fmtID(d *libvirt.Domain) string {
	id := "-"
	s, _, err := d.GetState()
	if err != nil {
		return "-"
	}
	if s != libvirt.DOMAIN_SHUTOFF {
		u, _ := d.GetID()
		id = fmt.Sprintf("%d", u)
	}

	return id
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
	keys  keymap
	help  help.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	r := m.table.Cursor()
	d := domains[r]
	s, _, _ := d.GetState()

	switch msg := msg.(type) {

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
		case key.Matches(msg, m.keys.Start):
			if err := d.Create(); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Start")
			return m, nil

		// Pause/Resume
		case key.Matches(msg, m.keys.PauseResume):
			if s == libvirt.DOMAIN_PAUSED {
				if err := d.Resume(); err != nil {
					fmt.Println("ERR: ", err)
				}
				fmt.Println("Resume")
			} else {
				if err := d.Suspend(); err != nil {
					fmt.Println("ERR: ", err)
				}
				fmt.Println("Suspend")
			}
			return m, nil

		// Shutdown
		case key.Matches(msg, m.keys.Shutdown):
			if err := d.Shutdown(); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Shutdown")
			return m, nil

		// Reboot
		case key.Matches(msg, m.keys.Reboot):
			if err := d.Reboot(0); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Reboot")
			return m, nil

		// Reset
		case key.Matches(msg, m.keys.Reset):
			if err := d.Reset(0); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Reset")
			return m, nil

		// Off
		case key.Matches(msg, m.keys.PowerOff):
			return m, nil

		// Save/Restore
		case key.Matches(msg, m.keys.SaveRestore):
			if s == libvirt.DOMAIN_RUNNING || s == libvirt.DOMAIN_PAUSED {
				if err := d.ManagedSave(0); err != nil {
					fmt.Println("ERR: ", err)
				}
			} else {
				if err := d.Create(); err != nil {
					fmt.Println("ERR: ", err)
				}
			}
			return m, nil

		// Delete
		case key.Matches(msg, m.keys.Delete):
			if err := d.Destroy(); err != nil {
				fmt.Println("ERR: ", err)
			}

			return m, nil

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()+"\n"+m.help.View(m.keys)) + "\n"
}

func main() {
	conn, err := libvirt.NewConnect(libvirtURI)
	if err != nil {
		fmt.Println("new connection: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	flags := libvirt.ConnectListAllDomainsFlags(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	domains, _ = conn.ListAllDomains(flags)

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

	for _, d := range domains {
		name, _ := d.GetName()
		state := fmtState(&d)
		id := fmtID(&d)

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

	m := model{
		table: t,
		help:  help.New(),
		keys:  keys,
	}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("failed to run program: ", err)
		os.Exit(1)
	}
}
