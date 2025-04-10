package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nixpig/virtui/manager"
	"libvirt.org/go/libvirt"
)

const libvirtURI = "qemu:///system"

func main() {
	// TODO: set up logging to file
	if err := libvirt.EventRegisterDefaultImpl(); err != nil {
		fmt.Println("register default event loop: ", err)
		os.Exit(1)
	}

	f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.SetOutput(f)

	go func() {
		for {
			if err := libvirt.EventRunDefaultImpl(); err != nil {
				fmt.Println("run default event loop: ", err)
				os.Exit(1)
			}
		}
	}()

	conn, err := libvirt.NewConnect(libvirtURI)
	if err != nil {
		// TODO: print and log err
		fmt.Println("failed to connect to: ", libvirtURI)
		os.Exit(1)
	}
	defer conn.Close()

	m := manager.InitModel(conn)

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("failed to run program: ", err)
		// TODO: print and log err
		os.Exit(1)
	}
}
