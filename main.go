package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/digitalocean/go-libvirt"
	"github.com/nixpig/virtui/vm/domain"
	"github.com/nixpig/virtui/vm/network"
	"github.com/nixpig/virtui/vm/volume"
)

func main() {

	uri, err := url.Parse("qemu:///system")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := libvirt.ConnectToURI(uri)
	defer conn.ConnectClose()

	// --- CREATE VOLUME

	vol := volume.NewWithDefaults("default-vm.qcow2")
	volXML, err := vol.ToXML()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := conn.StoragePoolLookupByName("default")
	if err != nil {
		log.Fatal("get storage pool: ", err.Error())
	}

	if _, err := conn.StorageVolCreateXML(pool, string(volXML), 0); err != nil {
		log.Fatal("create storage volume: " + err.Error())
	}

	// --- CREATE NETWORK

	vnet := network.NewWithDefaults("default1")
	vnetXML, err := vnet.ToXML()
	if err != nil {
		log.Fatal("network to xml: " + err.Error())
	}

	if n, err := conn.NetworkDefineXML(string(vnetXML)); err != nil {
		log.Fatal("define network: " + err.Error())
	} else {
		if err := conn.NetworkCreate(n); err != nil {
			log.Fatal("start network: " + err.Error())
		}
	}

	// ------------------------------------------------

	// --- CREATE DOMAIN

	dom := domain.NewWithDefaults("network-test-vm")

	domXML, err := dom.ToXML()
	if err != nil {
		log.Fatal("domain to xml: " + err.Error())
	}

	if d, err := conn.DomainDefineXML(string(domXML)); err != nil {
		log.Fatal("define domain: " + err.Error())
	} else {
		if err := conn.DomainCreate(d); err != nil {
			log.Fatal("start domain: " + err.Error())
		}
	}

	fmt.Println("done!")

	// ------------------------------------------------

	// db, err := database.NewConnection("virtui.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// mig, err := database.NewMigration(db, database.Migrations)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// if err := mig.Up(); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// store := connection.NewConnectionStoreImpl(db)
	//
	// if err := store.InsertConnection(&connection.Connection{
	// 	URL: "qemu:///session",
	// 	// URL:         "qemu:///system",
	// 	Autoconnect: true,
	// }); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// c, err := store.GetConnectionByID(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// conn, err := libvirt.NewConnect(c.URL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()
	//
	// flags := libvirt.ConnectListAllvmsFlags(libvirt.CONNECT_LIST_vmS_ACTIVE | libvirt.CONNECT_LIST_vmS_INACTIVE)
	//
	// vms, err := conn.ListAllvms(flags)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, d := range vms {
	// 	id, err := d.GetUUIDString()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	name, err := d.GetName()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	fmt.Printf("%s (%s)\n", name, id)
	// }
	//
	// // STEPS
	// // 1. Get an ISO
	// // 2. Create logical volume pool (or use /var/lib/libvirt/images)
	// // 3. Create a .qcow2 disk image in the pool
	// // ...
	// // n. Create XML definition
	// // n. conn.vmDefineXML() // conn.vmCreateXML()
	//
	// x, err := os.ReadFile("test-vm.xml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// d, err := conn.vmDefineXML(string(x))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// name, err := d.GetName()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println("Name: ", name)
	//
	// // run it
	// fmt.Println("Run...")
	// if err := d.Create(); err != nil {
	// 	log.Fatal(err)
	// }

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
