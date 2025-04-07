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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
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
	domains, _ := conn.ListAllDomains(flags)

	w, _, err := term.GetSize(0)
	if err != nil {
		fmt.Println("failed to get term size: ", err)
	}

	idWidth := 4
	stateWidth := 8
	nameWidth := w - idWidth - stateWidth - 4 - 4 - 2

	columns := []table.Column{
		{Title: "ID", Width: idWidth},
		{Title: "Name", Width: nameWidth},
		{Title: "State", Width: stateWidth},
	}

	rows := make([]table.Row, len(domains))

	for i, d := range domains {
		name, _ := d.GetName()
		state := fmtState(&d)
		id := fmtID(&d)

		rows[i] = table.Row{id, name, state}
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
		Bold(false)

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
