package guest

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg struct {
}

type Model struct {
	activeGuestID uint
}

func New(id uint) tea.Model {
	return Model{
		activeGuestID: id,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("guest model view - %d", m.activeGuestID)
}
