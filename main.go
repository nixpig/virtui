package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/digitalocean/go-libvirt"
	"github.com/nixpig/virtui/vm"
)

func main() {

	uri, err := url.Parse("qemu:///system")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := libvirt.ConnectToURI(uri)
	defer conn.ConnectClose()

	// --- CREATE VOLUME

	volume := vm.NewVolumeWithDefaults("default-vm.qcow2")
	volXML, err := volume.ToXMLFormatted()
	if err != nil {
		log.Fatal("volume to xml: " + err.Error())
	}

	fmt.Println(string(volXML))

	pool, err := conn.StoragePoolLookupByName("default")
	if err != nil {
		log.Fatal("get storage pool: ", err.Error())
	}

	v, err := conn.StorageVolCreateXML(pool, string(volXML), 0)
	if err != nil {
		log.Fatal("create storage volume: " + err.Error())
	}

	fmt.Println("created volume: ", v.Name)

	// --- CREATE NETWORK

	network := vm.NewNetworkWithDefaults("default1")
	netXML, err := network.ToXML()
	if err != nil {
		log.Fatal("network to xml: " + err.Error())
	}

	n, err := conn.NetworkDefineXML(string(netXML))
	if err != nil {
		log.Fatal("define network: " + err.Error())
	}

	if err := conn.NetworkCreate(n); err != nil {
		log.Fatal("start network: " + err.Error())
	}

	fmt.Println("created network: ", n.Name)

	// ------------------------------------------------

	// --- CREATE DOMAIN

	domain := vm.NewDomainWithDefaults("network-test-vm")

	output, err := domain.ToXML()
	if err != nil {
		log.Fatal("domain to xml: " + err.Error())
	}

	dom, err := conn.DomainDefineXML(string(output)) // create persistent
	if err != nil {
		log.Fatal("define domain: " + err.Error())
	}

	if err := conn.DomainCreate(dom); err != nil {
		log.Fatal("start domain: " + err.Error())
	}

	fmt.Println("created domain: ", dom.Name)

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
