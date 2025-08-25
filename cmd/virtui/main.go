package main

import (
	"context"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/libvirtui"
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

	// TODO: I'm not satisfied everything shuts down gracefully; I'm sure
	// BubbleTea is capturing, also we're not handling the cancelled context in
	// the libvirt connection

	// just a reminder that the event loop needs to be started _before_ the
	// connection is created
	if err := libvirtui.StartEventLoop(); err != nil {
		abort("failed to start libvirt event loop", err)
	}

	conn, err := libvirtui.NewConnection(ctx, qemuURI)
	if err != nil {
		abort("failed to establish connection with hypervisor", err)
	}
	defer func() {
		if _, err := conn.Close(); err != nil {
			log.Error("failed to close libvirt connection", "err", err)
		}
	}()

	appScreens := []app.Screen{
		manager.NewManagerScreen(),
		storage.NewStorageScreen(),
		network.NewNetworkScreen(),
	}

	service := libvirtui.NewService(conn)

	initialModel := app.NewAppModel(conn, service, appScreens)

	program := tea.NewProgram(
		initialModel,
		tea.WithContext(ctx),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := program.Run(); err != nil {
		abort("unrecoverable error", err)
	}
}

func abort(msg string, err error) {
	log.Error(msg, "err", err)
	os.Stderr.WriteString("Error: " + msg + "\n")
	os.Exit(1)
}
