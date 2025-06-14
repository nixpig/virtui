package manager

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/internal/entity"
	"github.com/nixpig/virtui/internal/keys"
	"libvirt.org/go/libvirt"
)

type SelectMsg struct {
	SelectedUUID string
}

var columns = []table.Column{
	// UUID is hidden as it's only used for identification
	{Title: "UUID", Width: 0},
	{Title: "ID", Width: 4},
	{Title: "Name", Width: 30},
	{Title: "State", Width: 10},
}

type Model struct {
	domains []libvirt.Domain
	lv      *libvirt.Connect
	keys    keys.Keymap
	table   table.Model
}

func New(lv *libvirt.Connect) tea.Model {
	domains, _ := lv.ListAllDomains(0)

	rows := make([]table.Row, len(domains))

	for i, d := range domains {
		x, _ := entity.ToDomainStruct(&d)
		state, _, _ := d.GetState()
		rows[i] = table.Row{x.UUID, fmt.Sprintf("%d", x.ID), x.Name, fmt.Sprintf("%v", state)}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithRows(rows),
	)

	return Model{
		domains: domains,
		table:   t,
		keys:    keys.Keys,
		lv:      lv,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// TODO: resize the table and stuff

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Select):
			return m, func() tea.Msg {
				return SelectMsg{
					SelectedUUID: m.table.SelectedRow()[0],
				}
			}
		}

	}

	m.table, cmd = m.table.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.table.View() + "\n"
}
