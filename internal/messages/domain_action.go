package messages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/internal/libvirtui"
)

type DomainActionWithFuncMsg struct {
	DomainUUID string
	Action     func(service libvirtui.Service, uuid string) error
}

func NewDomainActionWithFunc(uuid string, action func(service libvirtui.Service, uuid string) error) tea.Cmd {
	return func() tea.Msg {
		return DomainActionWithFuncMsg{
			DomainUUID: uuid,
			Action:     action,
		}
	}
}
