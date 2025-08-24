package main

import (
	"context"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/nixpig/virtui/internal/app"
	libvirtconn "github.com/nixpig/virtui/internal/libvirt/conn"
	"github.com/nixpig/virtui/internal/libvirt/events"
	"github.com/nixpig/virtui/internal/screens/manager"
	"github.com/nixpig/virtui/internal/screens/network"
	"github.com/nixpig/virtui/internal/screens/storage"
	"github.com/spf13/pflag"
)

func main() {
	var debug bool
	var logPath string
	var qemuURI string

	pflag.BoolVarP(&debug, "debug", "d", false, "set DEBUG log level")
	pflag.StringVarP(&qemuURI, "uri", "u", "qemu:///system", "set QEMU URI")
	pflag.StringVarP(&logPath, "log", "l", "/tmp/virtui.log", "set log path")
	pflag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	logFile, err := os.OpenFile(
		logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		os.Stderr.WriteString("Error: unable to open log file\n")
		os.Exit(1)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetPrefix(uuid.NewString())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := libvirtconn.New(ctx, qemuURI)
	if err != nil {
		log.Error("failed to establish connection with hypervisor", "err", err)
		os.Stderr.WriteString(
			"Error: failed to establish hypervisor connection and need to exit\n",
		)
		os.Exit(1)
	}

	if err := events.RegisterDefaultEventLoopImpl(); err != nil {
		log.Error("failed to register default event loop", "err", err)
	}

	go func() {
		for {
			// TODO: pass context and close event loop cleanly on exit and unregister handlers
			if err := events.RunDefaultEventLoopImpl(); err != nil {
				log.Error("failed to run default event loop", "err", err)
			}
		}
	}()

	initialModel := app.NewAppModel(
		conn,
		[]app.Screen{
			manager.NewManagerScreen(),
			storage.NewStorageScreen(),
			network.NewNetworkScreen(),
		},
	)

	p := tea.NewProgram(
		initialModel,
		tea.WithContext(ctx),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if model, err := p.Run(); err != nil {
		log.Error("unrecoverable error", "err", err, "model", model)
		os.Stderr.WriteString(
			"Error: encountered unrecoverable error and need to exit\n",
		)
		os.Exit(1)
	}

}
