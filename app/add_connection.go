package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/connection"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ OK ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("OK"))
)

type addConnectionModel struct {
	connection *connection.Connection
	form       []textinput.Model
	selected   int
}

func addConnectionScreen() addConnectionModel {
	hvInput := textinput.New()
	hvInput.Placeholder = "QEMU/KVM"
	hvInput.CharLimit = 10
	hvInput.Width = 30
	hvInput.Focus()

	urlInput := textinput.New()
	urlInput.Placeholder = "qemu:///system"
	urlInput.CharLimit = 255
	urlInput.Width = 30

	model := addConnectionModel{
		connection: &connection.Connection{},
		selected:   0,
		form: []textinput.Model{
			hvInput, urlInput,
		},
	}

	return model
}

func (m addConnectionModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m addConnectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab":
			s := msg.String()

			if s == "tab" {
				m.selected++
			} else {
				m.selected--
			}

			if m.selected > len(m.form) {
				m.selected = 0
			} else if m.selected < 0 {
				m.selected = len(m.form)
			}

			for i := range m.form {
				if i == m.selected {
					// focus etc..
					m.form[i].Focus()
					continue
				}
				// unfocus etc...
				m.form[i].Blur()
			}

		case "enter":
			if m.selected == len(m.form) {
				log.Info("submit form...")
			}

		default:
			if m.selected != len(m.form) {
				m.form[m.selected], cmd = m.form[m.selected].Update(msg)
				return m, cmd
			}
		}
	}

	return m, nil
}

func (m addConnectionModel) View() string {
	var b strings.Builder

	for i, v := range m.form {
		b.WriteString(v.View())
		if i < len(m.form)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.selected == len(m.form) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return lipgloss.NewStyle().Render(
		b.String(),
	)
}
