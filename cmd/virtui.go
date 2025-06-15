package main

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/tui"
	"libvirt.org/go/libvirt"
)

func main() {
	ctx := context.Background()

	uri := "qemu:///system"
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		log.Fatal("failed to connect to libvirt", "uri", uri, "err", err)
	}

	p := tea.NewProgram(
		tui.New(conn),
		tea.WithAltScreen(),
		tea.WithContext(ctx),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal("failed to start program", "err", err)
	}
}
