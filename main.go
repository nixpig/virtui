package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/digitalocean/go-libvirt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/nixpig/virtui/connection"
	"github.com/nixpig/virtui/database"
)

func main() {

	fmt.Println("new db conn")
	db, err := database.NewConnection("virtui.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("new db mig")
	mig, err := database.NewMigration(db, database.Migrations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mig up")
	if err := mig.Up(); err != nil && !errors.Is(migrate.ErrNoChange, err) {
		log.Fatal(err)
	}

	fmt.Println("new impl")
	store := connection.NewConnectionStoreImpl(db)

	fmt.Println("insert connection")
	if err := store.InsertConnection(&connection.Connection{
		// URL: "qemu:///session",
		URL:         "qemu:///system",
		Autoconnect: true,
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("get connection by id")
	c, err := store.GetConnectionByID(1)
	if err != nil {
		log.Fatal(err)
	}

	var conn *libvirt.Libvirt
	if c.Autoconnect {
		fmt.Println("autoconnect")
		uri, err := url.Parse(c.URL)
		if err != nil {
			log.Fatal(err)
		}

		conn, err = libvirt.ConnectToURI(uri)
		defer conn.ConnectClose()
	}

	if conn != nil {
		fmt.Println("list all domains")
		x, _, err := conn.ConnectListAllDomains(1, 0)
		if err != nil {
			log.Fatal(err)
		}

		for _, d := range x {
			fmt.Println(d.Name)
		}
	}

}
