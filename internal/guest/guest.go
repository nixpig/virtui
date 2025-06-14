package guest

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/internal/commands"
	"github.com/nixpig/virtui/internal/keys"
)

type Model struct {
	activeGuestUUID string
	keys            keys.Keymap
}

// New creates a tea.Model for the guest view
func New(uuid string) tea.Model {
	return Model{
		activeGuestUUID: uuid,
		keys:            keys.Keys,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Back):
			return m, commands.GoBackCmd()

		}
	}
	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("guest model view - %s", m.activeGuestUUID)
}
