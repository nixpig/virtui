package manager

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type basic struct {
	name  string
	id    string
	uuid  string
	state string
}

func (m basic) Init() tea.Cmd {
	return nil
}

func (m basic) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m basic) View() string {
	// conn, _ := m.vm.DomainGetConnect()
	//
	// t, _ := conn.GetType()
	// i, _ := conn.GetNodeInfo()
	// h, _ := conn.GetHostname()

	return fmt.Sprintf(
		"%s (%s) | %s | %s ",
		m.name,
		m.id,
		m.uuid,
		m.state,
	)
}
