package manager

import (
	tea "github.com/charmbracelet/bubbletea"
	"libvirt.org/go/libvirt"
)

type SelectMsg struct {
	ActiveGuestId uint
}

type Model struct {
	lv *libvirt.Connect
}

func New(lv *libvirt.Connect) tea.Model {
	return Model{
		lv: lv,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "manager model view"
}
