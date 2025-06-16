package main

import (
	"context"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/nixpig/virtui/tui"
	"github.com/spf13/pflag"
	"libvirt.org/go/libvirt"
)

var qemuSystemURI = "qemu:///system"

func main() {
	ctx := context.Background()

	var debug bool
	var logPath string
	pflag.BoolVarP(&debug, "debug", "d", false, "set debug log level")
	pflag.StringVarP(&logPath, "log", "l", "/tmp/virtui.log", "path to log output file")
	pflag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Stderr.WriteString("Error: unable to open log file")
		os.Exit(1)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetPrefix(uuid.NewString())

	conn, err := libvirt.NewConnect(qemuSystemURI)
	if err != nil {
		log.Debug("connect to libvirt", "uri", qemuSystemURI, "err", err)
		os.Stderr.WriteString("Error: failed to connect to libvirt")
		os.Exit(1)
	}

	p := tea.NewProgram(
		tui.New(conn),
		tea.WithAltScreen(),
		tea.WithContext(ctx),
	)

	if model, err := p.Run(); err != nil {
		log.Debug("unrecoverable error", "err", err, "model", model)
		os.Stderr.WriteString("Error: encountered unrecoverable error and need to exit")
		os.Exit(1)
	}
}
