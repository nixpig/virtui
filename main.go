package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/nixpig/virtui/vm"
	"libvirt.org/go/libvirt"
)

func main() {

	metadata := &vm.Metadata{
		LibOSInfo: &vm.LibOSInfo{
			LibOSInfo: "http://libosinfo.org/xmlns/libvirt/vm/1.0",
			OS:        &vm.OS{ID: "http://ubuntu.com/ubuntu/24.04"},
		},
	}

	d := &vm.Domain{
		Type:          "kvm",
		Name:          "another-test-vm",
		UUID:          "bdf08434-d11e-4953-90fd-5728b75224cd",
		Metadata:      metadata,
		Memory:        &vm.Memory{CharData: "2907152", Unit: "KiB"},
		CurrentMemory: &vm.CurrentMemory{CharData: "2907152", Unit: "KiB"},
		VCPU:          &vm.VCPU{CharData: "2", Placement: "static"},
		OS: &vm.OS{
			Type:     &vm.Type{Arch: "x86_64", Machine: "q35", CharData: vm.OS_TYPE_HVM},
			BootMenu: &vm.BootMenu{Enable: vm.FLAG_ENABLED_YES},
			// Kernel:  "/var/lib/libvirt/boot/virtinst-c6kdm5b8-vmlinuz",
			// InitRD:  "/var/lib/libvirt/boot/virtinst-ky336s4a-initrd",
			// Cmdline: "console=ttys0",
		},
		Features: &vm.Features{
			ACPI: &vm.ACPI{},
			APIC: &vm.APIC{},
		},
		Cpu: &vm.CPU{Mode: "host-passthrough", Check: "none", Migratable: "on"},
		Clock: &vm.Clock{
			Offset: "utc",
			Timers: []vm.Timer{
				{Name: "rtc", Tickpolicy: "catchup"},
				{Name: "pit", Tickpolicy: "delay"},
				{Name: "hpet", Present: "no"},
			},
		},
		// TODO: figure out why it gets destroyed on shutdown from virt-manager
		//        even though 'destroy' isn't set. Is it the default behaviour?
		// OnPoweroff: "destroy",
		OnReboot: vm.ON_EVENT_ACTION_RESTART,
		// OnCrash:    "destroy",
		PM: &vm.PM{
			SuspendToMem:  &vm.SuspendToMem{Enabled: vm.FLAG_ENABLED_NO},
			SuspendToDisk: &vm.SuspendToDisk{Enabled: vm.FLAG_ENABLED_NO},
		},
		Devices: &vm.Devices{
			Emulator: "/usr/bin/qemu-system-x86_64",
			Disks: []vm.Disk{
				{
					Type:   "file",
					Device: "disk",
					Boot:   &vm.Boot{Order: 1},
					Driver: &vm.Driver{Name: "qemu", Type: "qcow2", Discard: "unmap"},
					Source: &vm.Source{File: "/var/lib/libvirt/images/test-vm.qcow2"},
					Target: &vm.Target{Dev: "vda", Bus: "virtio"},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x03",
						Slot:     "0x00",
						Function: "0x0",
					},
				},
				{
					Type:     "file",
					Device:   "cdrom",
					Boot:     &vm.Boot{Order: 2},
					Driver:   &vm.Driver{Name: "qemu", Type: "raw"},
					Source:   &vm.Source{File: "/var/lib/libvirt/images/ubuntu-24.04.2-live-server-amd64.iso"},
					Target:   &vm.Target{Dev: "sda", Bus: "sata"},
					Readonly: &vm.Readonly{},
					Address: &vm.Address{
						Type:       "drive",
						Controller: "0",
						Bus:        "0",
						Target:     "0",
						Unit:       "0",
					},
				},
			},
			Controllers: []vm.Controller{
				{
					Type:  "usb",
					Model: "qemu-xhci",
					Ports: &[]int{15}[0],
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x02",
						Slot:     "0x00",
						Function: "0x0",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root",
					Index: 0,
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 1,
					Target: &vm.Target{
						Chassis: 1,
						Port:    "0x10",
					},
					Address: &vm.Address{
						Type:          "pci",
						Domain:        "0x0000",
						Bus:           "0x00",
						Slot:          "0x02",
						Function:      "0x0",
						Multifunction: "on",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 2,
					Target: &vm.Target{
						Chassis: 2,
						Port:    "0x11",
					},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x00",
						Slot:     "0x02",
						Function: "0x1",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 3,
					Target: &vm.Target{
						Chassis: 3,
						Port:    "0x12",
					},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x00",
						Slot:     "0x02",
						Function: "0x2",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 4,
					Target: &vm.Target{
						Chassis: 4,
						Port:    "0x13",
					},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x00",
						Slot:     "0x02",
						Function: "0x3",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 5,
					Target: &vm.Target{
						Chassis: 5,
						Port:    "0x14",
					},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x00",
						Slot:     "0x02",
						Function: "0x4",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 6,
					Target: &vm.Target{
						Chassis: 6,
						Port:    "0x15",
					},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x00",
						Slot:     "0x02",
						Function: "0x5",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 7,
					Target: &vm.Target{
						Chassis: 7,
						Port:    "0x16",
					},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x00",
						Slot:     "0x02",
						Function: "0x6",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 8,
					Target: &vm.Target{
						Chassis: 8,
						Port:    "0x17",
					},
					Address: &vm.Address{
						Type:     "pci",
						Domain:   "0x0000",
						Bus:      "0x00",
						Slot:     "0x02",
						Function: "0x7",
					},
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 9,
					Target: &vm.Target{
						Chassis: 9,
						Port:    "0x18",
					},
					Address: &vm.Address{
						Type:          "pci",
						Domain:        "0x0000",
						Bus:           "0x00",
						Slot:          "0x03",
						Function:      "0x0",
						Multifunction: "on",
					},
				},
			},
			Interfaces: []vm.Interface{
				{
					Source: &vm.Source{Network: "default"},
					Mac:    &vm.Mac{Address: "52:54:00:66:21:bf"},
					Model: &vm.Model{
						Type: "virtio"},
					Type: "network",
					Address: &vm.Address{Type: "pci",
						Domain:   "0x0000",
						Bus:      "0x01",
						Slot:     "0x00",
						Function: "0x0",
					},
				},
			},
			Serials: []vm.Serial{
				{
					Target: &vm.Target{Type: "isa-serial", Port: "0", Model: &vm.Model{
						Name: "isa-serial",
					}}},
			},
			Consoles: []vm.Console{{Type: "pty", Target: &vm.Target{Type: "serial", Port: "0"}}},
			Channels: []vm.Channel{{
				Type: "unix",
				Source: &vm.Source{
					Mode: "bind",
				},
				Target: &vm.Target{
					Type: "virtio",
					Name: "org.qemu.guest_agent.0"},
			}},
			Graphics: []vm.Graphics{
				{
					Type:     "vnc",
					Port:     -1,
					Autoport: "yes",
					Listen: &vm.Listen{
						Type: "address",
					},
				},
			},
			Video: &vm.Video{
				Model: &vm.Model{Type: "vga", VRAM: 16384, Heads: 1, Primary: "yes"},
				Alias: &vm.Alias{Name: "video0"},
				Address: &vm.Address{
					Type:     "pci",
					Domain:   "0x0000",
					Bus:      "0x00",
					Slot:     "0x01",
					Function: "0x0",
				},
			},
			Watchdog: &vm.Watchdog{Model: "itco", Action: "reset"},
			MemBalloon: &vm.MemBalloon{Model: "virtio", Address: &vm.Address{
				Type:     "pci",
				Domain:   "0x0000",
				Bus:      "0x05",
				Slot:     "0x00",
				Function: "0x0",
			}},
			RNG: &vm.RNG{
				Model: "virtio",
				Backend: &vm.Backend{
					Model:    "random",
					CharData: "/dev/urandom",
				},
				Address: &vm.Address{
					Type:     "pci",
					Domain:   "0x0000",
					Bus:      "0x06",
					Slot:     "0x00",
					Function: "0x0",
				},
			},
		},
	}

	o, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal("out: " + err.Error())
	}

	// fmt.Println(string(o))

	// o, err := os.ReadFile("orig.xml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dom, err := conn.DomainCreateXML(string(o), libvirt.DomainCreateFlags(0))
	if err != nil {
		log.Fatal(err)
	}

	name, _ := dom.GetName()
	fmt.Println("Name: ", name)

	// create but not start
	// conn.vmDefineXML()

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
