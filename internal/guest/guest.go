package guest

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	activeGuestUUID string
}

func New(uuid string) tea.Model {
	return Model{
		activeGuestUUID: uuid,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("guest model view - %s", m.activeGuestUUID)
}
