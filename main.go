package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/connection"
	"github.com/nixpig/virtui/database"
	"libvirt.org/go/libvirt"
)

func main() {
	db, err := database.NewConnection("virtui.db")
	if err != nil {
		log.Fatal(err)
	}

	mig, err := database.NewMigration(db, database.Migrations)
	if err != nil {
		log.Fatal(err)
	}

	if err := mig.Up(); err != nil {
		log.Fatal(err)
	}

	store := connection.NewConnectionStoreImpl(db)

	if err := store.InsertConnection(&connection.Connection{
		URL: "qemu:///session",
		// URL:         "qemu:///system",
		Autoconnect: true,
	}); err != nil {
		log.Fatal(err)
	}

	c, err := store.GetConnectionByID(1)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := libvirt.NewConnect(c.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	flags := libvirt.ConnectListAllDomainsFlags(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)

	domains, err := conn.ListAllDomains(flags)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range domains {
		id, err := d.GetUUIDString()
		if err != nil {
			log.Fatal(err)
		}

		name, err := d.GetName()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s (%s)\n", name, id)
	}

	// STEPS
	// 1. Get an ISO
	// 2. Create logical volume pool (or use /var/lib/libvirt/images)
	// 3. Create a .qcow2 disk image in the pool
	// ...
	// n. Create XML definition
	// n. conn.DomainDefineXML() // conn.DomainCreateXML()

	x, err := os.ReadFile("test-vm.xml")
	if err != nil {
		log.Fatal(err)
	}

	d, err := conn.DomainDefineXML(string(x))
	if err != nil {
		log.Fatal(err)
	}

	name, err := d.GetName()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name: ", name)

	// run it
	fmt.Println("Run...")
	if err := d.Create(); err != nil {
		log.Fatal(err)
	}

	// time.Sleep(5000)
	//
	// fmt.Println("Pause...")
	// if err := d.Suspend(); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// time.Sleep(5000)
	// fmt.Println("Resume...")
	// if err := d.Resume(); err != nil {
	// 	log.Fatal(err)
	// }

}
