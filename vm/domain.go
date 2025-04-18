package vm

import "github.com/google/uuid"

const (
	defaultEmulator = "/usr/bin/qemu-system-x86_64"
)

func NewDomainConfig() *DomainConfig {
	return &DomainConfig{}
}

func NewDefaultDomainConfig(name string) DomainConfig {
	id, _ := uuid.NewUUID()

	return DomainConfig{
		Type: DOMAIN_TYPE_KVM,
		Name: name,
		UUID: id.String(),
		Metadata: &Metadata{
			LibOSInfo: &LibOSInfo{
				LibOSInfo: "http://libosinfo.org/xmlns/libvirt/vm/1.0",
				OS:        &OS{ID: "http://ubuntu.com/ubuntu/24.04"},
			},
		},
		Memory:        &Memory{CharData: "2048", Unit: "MiB"},
		CurrentMemory: &CurrentMemory{CharData: "2048", Unit: "MiB"},
		VCPU:          &VCPU{CharData: "1", Placement: "static"},
		OS: &OS{
			Type:     &Type{Arch: ARCH_X86_64, Machine: MACHINE_TYPE_Q35, CharData: OS_TYPE_HVM},
			BootMenu: &BootMenu{Enable: FLAG_ENABLED_YES},
		},
		Features: &Features{ACPI: &ACPI{}, APIC: &APIC{}},
		Cpu:      &CPU{Mode: "host-passthrough", Check: "none", Migratable: "on"},
		Clock: &Clock{
			Offset: "utc",
			Timers: []Timer{
				{Name: "rtc", Tickpolicy: "catchup"},
				{Name: "pit", Tickpolicy: "delay"},
				{Name: "hpet", Present: "no"},
			},
		},
		OnReboot: ON_EVENT_ACTION_RESTART,
		PM: &PM{
			SuspendToMem:  &SuspendToMem{Enabled: FLAG_ENABLED_NO},
			SuspendToDisk: &SuspendToDisk{Enabled: FLAG_ENABLED_NO},
		},
		Devices: &Devices{
			Emulator: defaultEmulator,
			Disks: []Disk{
				{
					Type:   "file",
					Device: "disk",
					Boot:   &Boot{Order: 1},
					Driver: &Driver{Name: "qemu", Type: "qcow2", Discard: "unmap"},
					Source: &Source{File: "/var/lib/libvirt/images/test-vm.qcow2"},
					Target: &Target{Dev: "vda", Bus: "virtio"},
					Address: &Address{
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
					Boot:     &Boot{Order: 2},
					Driver:   &Driver{Name: "qemu", Type: "raw"},
					Source:   &Source{File: "/var/lib/libvirt/images/ubuntu-24.04.2-live-server-amd64.iso"},
					Target:   &Target{Dev: "sda", Bus: "sata"},
					Readonly: &Readonly{},
					Address: &Address{
						Type:       "drive",
						Controller: "0",
						Bus:        "0",
						Target:     "0",
						Unit:       "0",
					},
				},
			},
			Controllers: []Controller{
				{
					Type:  "usb",
					Model: "qemu-xhci",
					Ports: &[]int{15}[0],
					Address: &Address{
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
					Target: &Target{
						Chassis: 1,
						Port:    "0x10",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 2,
						Port:    "0x11",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 3,
						Port:    "0x12",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 4,
						Port:    "0x13",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 5,
						Port:    "0x14",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 6,
						Port:    "0x15",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 7,
						Port:    "0x16",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 8,
						Port:    "0x17",
					},
					Address: &Address{
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
					Target: &Target{
						Chassis: 9,
						Port:    "0x18",
					},
					Address: &Address{
						Type:          "pci",
						Domain:        "0x0000",
						Bus:           "0x00",
						Slot:          "0x03",
						Function:      "0x0",
						Multifunction: "on",
					},
				},
			},
			Interfaces: []Interface{
				{
					Source: &Source{Network: "default"},
					// Mac:    &Mac{Address: "52:54:00:66:21:bf"},
					Model:   &Model{Type: "virtio"},
					Type:    "network",
					Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x01", Slot: "0x00", Function: "0x0"},
				},
			},
			Serials: []Serial{
				{
					Target: &Target{Type: "isa-serial", Port: "0", Model: &Model{Name: "isa-serial"}},
				},
			},
			Consoles: []Console{
				{
					Type: "pty", Target: &Target{Type: "serial", Port: "0"},
				},
			},
			Channels: []Channel{
				{
					Type:   "unix",
					Source: &Source{Mode: "bind"},
					Target: &Target{Type: "virtio", Name: "org.qemu.guest_agent.0"},
				},
			},
			Graphics: []Graphics{
				{
					Type:     "vnc",
					Port:     -1,
					Autoport: "yes",
					Listen:   &Listen{Type: "address"},
				},
			},
			Video: &Video{
				Model:   &Model{Type: "vga", VRAM: 16384, Heads: 1, Primary: "yes"},
				Alias:   &Alias{Name: "video0"},
				Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x00", Slot: "0x01", Function: "0x0"},
			},
			Watchdog: &Watchdog{Model: "itco", Action: "reset"},
			MemBalloon: &MemBalloon{
				Model:   "virtio",
				Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x05", Slot: "0x00", Function: "0x0"},
			},
			RNG: &RNG{
				Model:   "virtio",
				Backend: &Backend{Model: "random", CharData: "/dev/urandom"},
				Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x06", Slot: "0x00", Function: "0x0"},
			},
		},
	}
}
