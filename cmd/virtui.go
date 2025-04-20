package main

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/nixpig/virtui/config"
	"github.com/nixpig/virtui/connection"
	"github.com/nixpig/virtui/database"
	"github.com/nixpig/virtui/tui"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const libvirtURI = "qemu:///system"

func main() {
	var configPath string
	pflag.StringVarP(&configPath, "config", "c", "", "specify custom config path")

	var debug bool
	pflag.BoolVarP(&debug, "debug", "v", false, "set log level to debug")

	v := viper.New()
	if err := config.Initialise(configPath, v); err != nil {
		fatality("failed to initialise config", err)
	}

	logfile := v.GetString(config.LogfileKey)
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fatality("failed to open log file", err)
	}
	defer f.Close()

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	log.SetOutput(f)

	dbName := v.GetString(config.DatabaseKey)
	dbConn, err := database.NewConnection(dbName)
	if err != nil {
		fatality("failed to open database connection", err)
	}
	defer dbConn.Close()

	mig, err := database.NewMigration(dbConn, database.Migrations)
	if err != nil {
		fatality("failed to create migration", err)
	}

	if err := mig.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			fatality("failed to run migration", err)
		}
	}

	store := connection.NewConnectionStoreImpl(dbConn)

	if _, err := tea.NewProgram(
		tui.InitModel(store),
		tea.WithAltScreen(),
	).Run(); err != nil {
		fatality("failed to run virtui", err)
	}
}

func fatality(msg string, err error) {
	os.Stderr.Write(fmt.Appendf([]byte{}, "Error: %s: %s", msg, err))
	os.Exit(1)
}
