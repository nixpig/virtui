package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/nixpig/virtui/domain"
)

func main() {

	d := &domain.Domain{
		Type: "kvm",
		Name: "test-vm",
		Uuid: "63cfcedf-3de1-433f-80a0-9b39bfaa9605",
		Metadata: &domain.Metadata{
			Libosinfo: &domain.Libosinfo{
				Libosinfo: "http://libosinfo.org/xmlns/libvirt/domain/1.0",
				Os: &domain.Os{
					ID: "http://ubuntu.com/ubuntu/24.04",
				},
			},
		},
		Memory: &domain.Memory{
			CharData: "2907152",
		},
		CurrentMemory: &domain.CurrentMemory{
			CharData: "2907152",
		},
		Vcpu: &domain.Vcpu{CharData: "2"},
		Os: &domain.Os{
			Type: &domain.Type{
				Arch:    "x86_64",
				Machine: "q35",
			},
			Kernel:  "/var/lib/libvirt/boot/virtinst-c6kdm5b8-vmlinuz",
			Initrd:  "/var/lib/libvirt/boot/virtinst-ky336s4a-initrd",
			Cmdline: "console=ttys0",
		},
		Features: &domain.Features{
			Acpi: &domain.Acpi{},
			Apic: &domain.Apic{},
		},
		Cpu: &domain.Cpu{
			Mode: "host-passthrough",
		},
		Clock: &domain.Clock{
			Offset: "utc",
			Timer: []domain.Timer{
				{Name: "rtc", Tickpolicy: "catchup"},
				{Name: "pit", Tickpolicy: "delay"},
				{Name: "hpet", Present: "no"},
			},
		},
		Pm: &domain.Pm{
			SuspendToMem: &domain.SuspendToMem{
				Enabled: "no",
			},
			SuspendToDisk: &domain.SuspendToDisk{
				Enabled: "no",
			},
		},
		Devices: &domain.Devices{
			Emulator: "/usr/bin/qemu-system-x86_64",
			Disk: []domain.Disk{
				{
					Type:   "file",
					Device: "disk",
					Driver: []domain.Driver{{Name: "qemu", Type: "qcow2"}},
					Source: &domain.Source{File: "/var/lib/libvirt/images/test-vm.qcow2"},
					Target: &domain.Target{Dev: "vda", Bus: "virtio"},
				},
				{
					Type:     "file",
					Device:   "cdrom",
					Driver:   []domain.Driver{{Name: "qemu", Type: "raw"}},
					Source:   &domain.Source{File: "/var/lib/libvirt/images/ubuntu-24.04.2-live-server-amd64.iso"},
					Target:   &domain.Target{Dev: "sda", Bus: "sata"},
					Readonly: &domain.Readonly{},
				},
			},
			Controller: []domain.Controller{
				{Type: "usb", Model: "qemu-xhci", Ports: &[]int{15}[0]},
				{Type: "pci", Model: "pcie-root"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
				{Type: "pci", Model: "pcie-root-port"},
			},
			Interface: []domain.Interface{
				{
					Source: []domain.Source{{Network: "default"}},
					Mac:    &domain.Mac{Address: "52:54:00:66:21:bf"},
					Model:  &domain.Model{Type: "virtio"},
					Type:   "network",
				},
			},
			Console: []domain.Console{{Type: "pty"}},
			Channel: []domain.Channel{{
				Type:   "unix",
				Source: &domain.Source{Mode: "bind"},
				Target: &domain.Target{Type: "virtio", Name: "org.qemu.guest_agent.0"},
			}},
			Memballoon: []domain.Memballoon{
				{Model: "virtio"},
			},
			Rng: &domain.Rng{
				Model: "virtio",
				Backend: []domain.Backend{
					{
						Model:    "random",
						CharData: "/dev/urandom",
					},
				},
			},
		},
		OnReboot: "destroy",
	}

	o, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal("out: " + err.Error())
	}

	fmt.Println(string(o))

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
	// flags := libvirt.ConnectListAllDomainsFlags(libvirt.CONNECT_LIST_DOMAINS_ACTIVE | libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	//
	// domains, err := conn.ListAllDomains(flags)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// for _, d := range domains {
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
	// // n. conn.DomainDefineXML() // conn.DomainCreateXML()
	//
	// x, err := os.ReadFile("test-vm.xml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// d, err := conn.DomainDefineXML(string(x))
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
