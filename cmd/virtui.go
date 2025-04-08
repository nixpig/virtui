package main

import (
	"fmt"
	"os"

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
		switch msg.String() {

		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}

		// Quit
		case "q", "Q", "ctrl+c":
			return m, tea.Quit

		// Start
		case "t", "T":
			if err := d.Create(); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Start")
			return m, nil

		// Pause/Resume
		case "p", "P":
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
		case "s", "S":
			if err := d.Shutdown(); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Shutdown")
			return m, nil

		// Reboot
		case "r", "R":
			if err := d.Reboot(0); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Reboot")
			return m, nil

		// Reset
		case "o", "O":
			if err := d.Reset(0); err != nil {
				fmt.Println("ERR: ", err)
			}
			fmt.Println("Reset")
			return m, nil

		// Off
		case "f", "F":
			return m, nil

		// Save/Restore
		case "v", "V":
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
		case "d", "D":
			if err := d.Destroy(); err != nil {
				fmt.Println("ERR: ", err)
			}

			return m, nil
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
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

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("failed to run program: ", err)
		os.Exit(1)
	}
}
