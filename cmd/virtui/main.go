package main

import (
	"context"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/nixpig/virtui/internal/tui"
	"github.com/spf13/pflag"
	"libvirt.org/go/libvirt"
)

func main() {
	var debug bool
	var logPath string
	var qemuURI string

	pflag.BoolVarP(&debug, "debug", "d", false, "set debug log level")
	pflag.StringVarP(
		&logPath,
		"log",
		"l",
		"/tmp/virtui.log",
		"path to log output file",
	)
	pflag.StringVarP(&qemuURI, "uri", "u", "qemu:///system", "QEMU URI")
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
		os.Stderr.WriteString("Error: unable to open log file")
		os.Exit(1)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetPrefix(uuid.NewString())

	log.Debug("settings", "debug", debug, "logPath", logPath)

	conn, err := libvirt.NewConnect(qemuURI)
	if err != nil {
		log.Error("connect to libvirt", "err", err)
		os.Stderr.WriteString("Error: failed to connect to libvirt\n")
		os.Exit(1)
	}

	defer conn.Close()

	hostname, _ := conn.GetHostname()
	lvVersion, _ := conn.GetLibVersion()
	hvVersion, _ := conn.GetVersion()
	connectionType, _ := conn.GetType()
	hostinfo, _ := conn.GetNodeInfo()

	log.Debug(
		"connection",
		"hostname", hostname,
		"qemuURI", qemuURI,
		"connectionType", connectionType,
		"hvVersion", hvVersion,
		"lvVersion", lvVersion,
		"hostinfo", hostinfo,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	initialModel, _ := tui.New(conn, ctx)

	p := tea.NewProgram(
		initialModel,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
		tea.WithContext(ctx),
	)

	if model, err := p.Run(); err != nil {
		log.Error("unrecoverable error", "err", err, "model", model)
		os.Stderr.WriteString(
			"Error: encountered unrecoverable error and need to exit\n",
		)
		os.Exit(1)
	}
}
