package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"runtime"

	"github.com/digitalocean/go-libvirt"
	"github.com/nixpig/virtui/vm"
)

func main() {

	d := vm.NewDefaultDomainConfig()

	o, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal("out: " + err.Error())
	}

	uri, err := url.Parse("qemu:///system")
	if err != nil {
		log.Fatal(err)
	}
	l, err := libvirt.ConnectToURI(uri)
	defer l.ConnectClose()

	runtime.NumCPU()

	// conn, err := libvirt.NewConnect("qemu:///system")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	dom, err := l.DomainDefineXML(string(o)) // create persistent
	if err != nil {
		log.Fatal(err)
	}

	// dom, err := conn.LookupDomainByName("another-test-vm")
	// if err != nil {
	// 	log.Fatal("lookup: ", err)
	// }

	fmt.Println("Name: ", dom.Name)

	if err := l.DomainCreate(dom); err != nil {
		log.Fatal("start domain: ", err)
	}

	// time.Sleep(10 * time.Second)
	//
	// w := bytes.Buffer{}
	//
	// opt := libvirt.OptString{}
	//
	// go func() {
	// 	fmt.Println("opening console...")
	// 	if err := l.DomainOpenConsole(dom, opt, &w, 0); err != nil {
	// 		log.Fatal("open console: ", err)
	// 	}
	// }()
	//
	// time.Sleep(3 * time.Second)
	//
	// fmt.Println("apparently opened??")
	//
	// b := make([]byte, 1024)
	// w.Read(b)
	// fmt.Println("read: ", string(b))
	//
	// w.WriteString("foobarbaz")

	// create but not start
	// conn.DomainCreateXML() // create transient

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
