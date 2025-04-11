package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/manager"
	"libvirt.org/go/libvirt"
)

const libvirtURI = "qemu:///system"

func main() {
	f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.SetOutput(f)

	if err := libvirt.EventRegisterDefaultImpl(); err != nil {
		log.Error("register libvirt event loop", "err", err)
		os.Stderr.Write(fmt.Appendf([]byte(""), "Error: failed to register event loop: %s", err.Error()))
		os.Exit(1)
	}

	go func() {
		for {
			if err := libvirt.EventRunDefaultImpl(); err != nil {
				log.Error("run libvirt event loop", "err", err)
				os.Stderr.Write(fmt.Appendf([]byte(""), "Error: failed to run event loop: %s", err.Error()))
				os.Exit(1)
			}
		}
	}()

	conn, err := libvirt.NewConnect(libvirtURI)
	if err != nil {
		log.Error("connect to qemu", "err", err)
		os.Stderr.Write(fmt.Appendf([]byte(""), "Error: failed to connect (%s): %s", libvirtURI, err.Error()))
		os.Exit(1)
	}
	defer conn.Close()

	m := manager.InitModel(conn)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Error("start bubbletea program", "err", err)
		os.Stderr.Write(fmt.Appendf([]byte(""), "Error: failed to start: %s", err.Error()))
		os.Exit(1)
	}
}
