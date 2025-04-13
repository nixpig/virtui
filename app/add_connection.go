package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/connection"
)

type addConnectionModel struct {
	connection *connection.Connection
	hvInput    textinput.Model
	urlInput   textinput.Model
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
		urlInput:   urlInput,
		hvInput:    hvInput,
	}

	return model
}

func (m addConnectionModel) Init() tea.Cmd {
	return nil
}

func (m addConnectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.urlInput, cmd = m.urlInput.Update(msg)
		return m, cmd

	}
	return m, nil
}

func (m addConnectionModel) View() string {
	hvLabel := lipgloss.NewStyle().Inline(true).Render("Hypervisor type ")
	urlLabel := lipgloss.NewStyle().Inline(true).Render("Connection URL ")

	autoConnect := lipgloss.NewStyle().Inline(true).Render("[ ] Autoconnect")

	return lipgloss.NewStyle().Render(
		hvLabel + m.hvInput.View() + "\n" +
			urlLabel + m.urlInput.View() + "\n" +
			autoConnect,
	)
}
