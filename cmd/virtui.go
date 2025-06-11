package main

import (
	"context"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

type model struct {
	connections []connection
}

type connection struct {
	name               string
	uri                string
	connected          bool
	autoconnect        bool
	domains            []domain
	networkConnections []networkConnection
	storagePools       []storagePool
}

type domain struct {
	name  string
	uuid  string
	state string
}

type networkConnection struct{}

type storagePool struct{}

func initialModel() model {
	return model{
		connections: []connection{
			{
				name:        "first",
				uri:         "hsdaofsadoifs",
				connected:   true,
				autoconnect: false,
				domains: []domain{
					{
						name:  "some domain",
						uuid:  "some-uuid-in-this-place",
						state: "Running",
					},
				},
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	for _, c := range m.connections {
		s.WriteString(c.name)
	}

	return fmt.Sprintf("%s\n", s.String())
}

func main() {
	ctx := context.Background()

	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithContext(ctx),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal("failed to start program", "err", err)
	}
}
