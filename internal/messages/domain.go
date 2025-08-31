package messages

import tea "github.com/charmbracelet/bubbletea"

type DomainActionMsg struct {
	Action    string
	Domain    string
	IsPrompt  bool
}

func NewDomainAction(action, domain string, isPrompt bool) tea.Cmd {
	return func() tea.Msg {
		return DomainActionMsg{
			Action:   action,
			Domain:   domain,
			IsPrompt: isPrompt,
		}
	}
}
