package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/nixpig/virtui/domain"
	"libvirt.org/go/libvirt"
)

func main() {

	metadata := &domain.Metadata{
		LibOSInfo: &domain.LibOSInfo{
			LibOSInfo: "http://libosinfo.org/xmlns/libvirt/domain/1.0",
			OS:        &domain.OS{ID: "http://ubuntu.com/ubuntu/24.04"},
		},
	}

	d := &domain.Domain{
		Type:          "kvm",
		Name:          "another-test-vm",
		UUID:          "bdf08434-d11e-4953-90fd-5728b75224cd",
		Metadata:      metadata,
		Memory:        &domain.Memory{CharData: "2907152", Unit: "KiB"},
		CurrentMemory: &domain.CurrentMemory{CharData: "2907152", Unit: "KiB"},
		VCPU:          &domain.VCPU{CharData: "2", Placement: "static"},
		OS: &domain.OS{
			Type: &domain.Type{Arch: "x86_64", Machine: "q35", CharData: "hvm"},
			Boot: &domain.Boot{Dev: "hd"},
			// Kernel:  "/var/lib/libvirt/boot/virtinst-c6kdm5b8-vmlinuz",
			// InitRD:  "/var/lib/libvirt/boot/virtinst-ky336s4a-initrd",
			Cmdline: "console=ttys0",
		},
		Features: &domain.Features{
			ACPI: &domain.ACPI{},
			APIC: &domain.APIC{},
		},
		Cpu: &domain.CPU{Mode: "host-passthrough", Check: "none", Migratable: "on"},
		Clock: &domain.Clock{
			Offset: "utc",
			Timers: []domain.Timer{
				{Name: "rtc", Tickpolicy: "catchup"},
				{Name: "pit", Tickpolicy: "delay"},
				{Name: "hpet", Present: "no"},
			},
		},
		OnPoweroff: "destroy",
		OnReboot:   "restart",
		OnCrash:    "destroy",
		PM: &domain.PM{
			SuspendToMem:  &domain.SuspendToMem{Enabled: "no"},
			SuspendToDisk: &domain.SuspendToDisk{Enabled: "no"},
		},
		Devices: &domain.Devices{
			Emulator: "/usr/bin/qemu-system-x86_64",
			Disks: []domain.Disk{
				{
					Type:   "file",
					Device: "disk",
					Driver: &domain.Driver{Name: "qemu", Type: "qcow2", Discard: "unmap"},
					Source: &domain.Source{File: "/var/lib/libvirt/images/test-vm.qcow2"},
					Target: &domain.Target{Dev: "vda", Bus: "virtio"},
					Address: &domain.Address{Type: "pci",
						Domain:   "0x0000",
						Bus:      "0x03",
						Slot:     "0x00",
						Function: "0x0",
					},
				},
				{
					Type:     "file",
					Device:   "cdrom",
					Driver:   &domain.Driver{Name: "qemu", Type: "raw"},
					Source:   &domain.Source{File: "/var/lib/libvirt/images/ubuntu-24.04.2-live-server-amd64.iso"},
					Target:   &domain.Target{Dev: "sda", Bus: "sata"},
					Readonly: &domain.Readonly{},
					Address:  &domain.Address{Type: "drive", Controller: "0", Bus: "0", Target: "0", Unit: "0"},
				},
			},
			Controllers: []domain.Controller{
				{Type: "usb", Model: "qemu-xhci", Ports: &[]int{15}[0], Address: &domain.Address{
					Type:   "pci",
					Domain: "0x0000", Bus: "0x02", Slot: "0x00", Function: "0x0",
				}},
				{
					Type:  "pci",
					Model: "pcie-root",
					Index: 0,
				},
				{
					Type:  "pci",
					Model: "pcie-root-port",
					Index: 1,
					Target: &domain.Target{
						Chassis: 1,
						Port:    "0x10",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 2,
						Port:    "0x11",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 3,
						Port:    "0x12",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 4,
						Port:    "0x13",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 5,
						Port:    "0x14",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 6,
						Port:    "0x15",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 7,
						Port:    "0x16",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 8,
						Port:    "0x17",
					},
					Address: &domain.Address{
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
					Target: &domain.Target{
						Chassis: 9,
						Port:    "0x18",
					},
					Address: &domain.Address{
						Type:          "pci",
						Domain:        "0x0000",
						Bus:           "0x00",
						Slot:          "0x03",
						Function:      "0x0",
						Multifunction: "on",
					},
				},
			},
			Interfaces: []domain.Interface{
				{
					Source: &domain.Source{Network: "default"},
					Mac:    &domain.Mac{Address: "52:54:00:66:21:bf"},
					Model:  &domain.Model{Type: "virtio"},
					Type:   "network",
					Address: &domain.Address{Type: "pci",
						Domain:   "0x0000",
						Bus:      "0x01",
						Slot:     "0x00",
						Function: "0x0",
					},
				},
			},
			Serials: []domain.Serial{
				{Target: &domain.Target{Type: "isa-serial", Port: "0", Model: &domain.Model{
					Name: "isa-serial",
				}}},
			},
			Consoles: []domain.Console{{Type: "pty", Target: &domain.Target{Type: "serial", Port: "0"}}},
			Channels: []domain.Channel{{
				Type:   "unix",
				Source: &domain.Source{Mode: "bind"},
				Target: &domain.Target{Type: "virtio", Name: "org.qemu.guest_agent.0"},
			}},
			Graphics: []domain.Graphics{
				{Type: "vnc", Port: -1, Autoport: "yes", Listen: &domain.Listen{Type: "address"}},
			},
			Audio: []domain.Audio{{ID: 1, Type: "none"}},
			Video: &domain.Video{
				Model: &domain.Model{Type: "vga", VRAM: 16384, Heads: 1, Primary: "yes"},
				Alias: &domain.Alias{Name: "video0"},
				Address: &domain.Address{Type: "pci",
					Domain:   "0x0000",
					Bus:      "0x00",
					Slot:     "0x01",
					Function: "0x0",
				},
			},
			Watchdog: &domain.Watchdog{Model: "itco", Action: "reset"},
			MemBalloon: &domain.MemBalloon{Model: "virtio", Address: &domain.Address{
				Type:     "pci",
				Domain:   "0x0000",
				Bus:      "0x05",
				Slot:     "0x00",
				Function: "0x0",
			}},
			RNG: &domain.RNG{
				Model:   "virtio",
				Backend: &domain.Backend{Model: "random", CharData: "/dev/urandom"},
				Address: &domain.Address{
					Type:     "pci",
					Domain:   "0x0000",
					Bus:      "0x06",
					Slot:     "0x00",
					Function: "0x0",
				},
			},
			Crypto:      &domain.Crypto{},
			Filesystem:  []domain.Filesystem{},
			HostDev:     []domain.HostDev{},
			Hub:         &domain.Hub{},
			Input:       []domain.Input{},
			IOMMU:       &domain.IOMMU{},
			Lease:       &domain.Lease{},
			Memory:      []domain.Memory{},
			NVRam:       &domain.NVRAM{},
			Panic:       []domain.Panic{},
			Parallel:    []domain.Parallel{},
			Pstore:      &domain.PStore{},
			RedirDev:    []domain.RedirDev{},
			RedirFilter: &domain.RedirFilter{},
			SHMem:       []domain.SHMem{},
			Smartcard:   []domain.Smartcard{},
			Sound:       []domain.Sound{},
			TPM:         []domain.Tpm{},
			VSock:       &domain.VSock{},
		},
		XMLName:         xml.Name{},
		BlkioTune:       &domain.BlkioTune{},
		CPUTune:         &domain.CPUTune{},
		DefaultIOThread: &domain.DefaultIOThread{},
		Description:     new(string),
		GenID:           new(string),
		IOThreadIDs:     &domain.IOThreadIDs{},
		IOThreads:       &domain.IOThread{},
		Keywrap:         &domain.Keywrap{},
		LaunchSecurity:  &domain.LaunchSecurity{},
		MaxMemory:       &domain.MaxMemory{},
		MemoryBacking:   &domain.MemoryBacking{},
		MemTune:         &domain.Memtune{},
		NumaTune:        &domain.Numatune{},
		OnLockFailure:   new(string),
		Override:        &domain.Override{},
		Perf:            &domain.Perf{},
		Resources:       []domain.Resource{},
		SecLabels:       []domain.SecLabel{},
		Sysinfo:         []domain.Sysinfo{},
		ThrottleGroups:  &domain.ThrottleGroups{},
		Title:           "",
		VCPUs:           &domain.VCPUs{},
	}

	o, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal("out: " + err.Error())
	}

	fmt.Println(string(o))

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
	// conn.DomainDefineXML()

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
