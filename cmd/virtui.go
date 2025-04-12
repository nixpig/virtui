package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/nixpig/virtui/config"
	"github.com/nixpig/virtui/database"
	"github.com/nixpig/virtui/internal/app"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"libvirt.org/go/libvirt"
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

	db := v.GetString(config.DatabaseKey)
	conn, err := database.NewConnection(db)
	if err != nil {
		fatality("failed to open database connection", err)
	}
	defer conn.Close()

	mig, err := database.NewMigration(conn, database.Migrations)
	if err != nil {
		fatality("failed to create migration", err)
	}

	if err := mig.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			fatality("failed to run migration", err)
		}
	}

	if err := libvirt.EventRegisterDefaultImpl(); err != nil {
		fatality("failed to register event loop", err)
	}

	go func() {
		for {
			if err := libvirt.EventRunDefaultImpl(); err != nil {
				log.Error("run event loop", "err", err)
				fatality("failed to run event loop", err)
			}
		}
	}()

	if err := app.Run(conn); err != nil {
		fatality("failed to run virtui", err)
	}
}

func fatality(msg string, err error) {
	os.Stderr.Write(fmt.Appendf([]byte{}, "Error: %s: %s", msg, err))
	os.Exit(1)
}
