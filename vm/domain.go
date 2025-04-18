package vm

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/google/uuid"
)

const (
	defaultEmulator = ""
)

func NewFromXML(xml string) (*DomainConfig, error) {
	return nil, nil
}

func NewFromFile(r io.Reader) (*DomainConfig, error) {
	return nil, nil
}

func New(name string) (*DomainConfig, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("generate uuid: %w", err)
	}

	return &DomainConfig{
		Name:     name,
		UUID:     id.String(),
		Metadata: &Metadata{LibOSInfo: &LibOSInfo{LibOSInfo: "", OS: &OS{}}},
		OS:       &OS{},
		Features: &Features{},
		Clock:    &Clock{},
		PM:       &PM{SuspendToDisk: &SuspendToDisk{}, SuspendToMem: &SuspendToMem{}},
		Devices:  &Devices{},
	}, nil

}

func (d *DomainConfig) ToXML() ([]byte, error) {
	return xml.Marshal(d)
}

func (d *DomainConfig) ToXMLFormatted() ([]byte, error) {
	return xml.MarshalIndent(d, "", "  ")
}

func (d *DomainConfig) SetName(n string) {
	d.Name = n
}

func (d *DomainConfig) SetUUID(u string) {
	d.UUID = u
}

func (d *DomainConfig) SetMetaData(m *Metadata) {
	d.Metadata = m
}

func (d *DomainConfig) SetMemory(m *Memory) {
	d.Memory = m
}

func (d *DomainConfig) SetCurrentMemory(c *CurrentMemory) {
	d.CurrentMemory = c
}

func (d *DomainConfig) SetVCPU(v *VCPU) {
	d.VCPU = v
}

func (d *DomainConfig) SetOSType(t *Type) {
	d.OS.Type = t
}

func (d *DomainConfig) SetBootMenu(b *BootMenu) {
	d.OS.BootMenu = b
}

func (d *DomainConfig) EnableACPI() {
	d.Features.ACPI = &ACPI{}
}

func (d *DomainConfig) DisableACPI() {
	d.Features.ACPI = nil
}

func (d *DomainConfig) EnableAPIC() {
	d.Features.APIC = &APIC{}
}

func (d *DomainConfig) DisableAPIC() {
	d.Features.APIC = nil
}

func (d *DomainConfig) SetCPU(c *CPU) {
	d.CPU = c
}

func (d *DomainConfig) SetClockOffset(o string) {
	d.Clock.Offset = o
}

func (d *DomainConfig) AddClockTimer(t Timer) {
	d.Clock.Timers = append(d.Clock.Timers, t)
}

func (d *DomainConfig) AddClockTimers(t ...Timer) {
	for _, x := range t {
		d.Clock.Timers = append(d.Clock.Timers, x)
	}
}

func (d *DomainConfig) EnableSuspendToMemory() {
	d.PM.SuspendToMem.Enabled = FLAG_ENABLED_YES
}

func (d *DomainConfig) DisableSuspendToMemory() {
	d.PM.SuspendToMem.Enabled = FLAG_ENABLED_NO
}

func (d *DomainConfig) EnableSuspendToDisk() {
	d.PM.SuspendToDisk.Enabled = FLAG_ENABLED_YES
}

func (d *DomainConfig) DisableSuspendToDisk() {
	d.PM.SuspendToDisk.Enabled = FLAG_ENABLED_NO
}

func (d *DomainConfig) SetEmulator(e string) {
	d.Devices.Emulator = e
}

func (d *DomainConfig) AddDisk(s Disk) {
	d.Devices.Disks = append(d.Devices.Disks, s)
}

func (d *DomainConfig) AddDisks(s ...Disk) {
	for _, x := range s {
		d.Devices.Disks = append(d.Devices.Disks, x)
	}
}

func (d *DomainConfig) AddController(c Controller) {
	d.Devices.Controllers = append(d.Devices.Controllers, c)
}

func (d *DomainConfig) AddControllers(c ...Controller) {
	for _, x := range c {
		d.Devices.Controllers = append(d.Devices.Controllers, x)
	}
}

func (d *DomainConfig) AddInterface(i Interface) {
	d.Devices.Interfaces = append(d.Devices.Interfaces, i)
}

func (d *DomainConfig) AddInterfaces(i ...Interface) {
	for _, x := range i {
		d.Devices.Interfaces = append(d.Devices.Interfaces, x)
	}
}

func (d *DomainConfig) AddSerial(s Serial) {
	d.Devices.Serials = append(d.Devices.Serials, s)
}

func (d *DomainConfig) AddConsole(c Console) {
	d.Devices.Consoles = append(d.Devices.Consoles, c)
}

func (d *DomainConfig) AddChannel(c Channel) {
	d.Devices.Channels = append(d.Devices.Channels, c)
}

func (d *DomainConfig) AddGraphics(g Graphics) {
	d.Devices.Graphics = append(d.Devices.Graphics, g)
}

func (d *DomainConfig) SetVideo(v *Video) {
	d.Devices.Video = v
}

func (d *DomainConfig) SetWatchdog(w *Watchdog) {
	d.Devices.Watchdog = w
}

func (d *DomainConfig) SetRNG(r *RNG) {
	d.Devices.RNG = r
}

func (d *DomainConfig) SetMemBalloon(m *MemBalloon) {
	d.Devices.MemBalloon = m
}

func (d *DomainConfig) SetOnReboot(o OnEventAction) {
	d.OnReboot = o
}

func (d *DomainConfig) SetOnPoweroff(o OnEventAction) {
	d.OnPoweroff = o
}

func (d *DomainConfig) SetOnCrash(o OnEventAction) {
	d.OnCrash = o
}

func (d *DomainConfig) SetLibOSInfo(l string) {
	d.Metadata.LibOSInfo.LibOSInfo = l
}

func (d *DomainConfig) SetLibOSID(l string) {
	d.Metadata.LibOSInfo.OS.ID = l
}

func NewWithDefaults(name string) (*DomainConfig, error) {
	c, err := New(name)
	if err != nil {
		return nil, err
	}

	c.Type = DOMAIN_TYPE_KVM

	c.SetLibOSInfo("http://libosinfo.org/xmlns/libvirt/vm/1.0")
	c.SetLibOSID("http://ubuntu.com/ubuntu/24.04")

	c.SetMemory(&Memory{CharData: "2048", Unit: "MiB"})

	c.SetCurrentMemory(&CurrentMemory{CharData: "2048", Unit: "MiB"})

	c.SetVCPU(&VCPU{CharData: "1", Placement: "static"})

	c.SetOSType(&Type{Arch: ARCH_X86_64, Machine: MACHINE_TYPE_Q35, CharData: OS_TYPE_HVM})

	c.SetBootMenu(&BootMenu{Enable: FLAG_ENABLED_YES})

	c.EnableACPI()
	c.EnableAPIC()

	c.SetCPU(&CPU{Mode: "host-passthrough", Check: "none", Migratable: "on"})

	c.SetClockOffset("utc")

	c.AddClockTimers(
		Timer{Name: "rtc", Tickpolicy: "catchup"},
		Timer{Name: "pit", Tickpolicy: "delay"},
		Timer{Name: "hpet", Present: "no"},
	)

	c.DisableSuspendToMemory()
	c.DisableSuspendToDisk()

	c.SetEmulator("/usr/bin/qemu-system-x86_64")

	c.AddDisks(
		Disk{
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
		Disk{
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
	)

	c.AddControllers(
		Controller{
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
		Controller{
			Type:  "pci",
			Model: "pcie-root",
			Index: 0,
		},
		Controller{
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
		Controller{
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
		Controller{
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
		Controller{
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
		Controller{
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
		Controller{
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
		Controller{
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
		Controller{
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
		Controller{
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
	)

	c.AddInterface(Interface{
		Source: &Source{Network: "default"},
		// Mac:    &Mac{Address: "52:54:00:66:21:bf"},
		Model:   &Model{Type: "virtio"},
		Type:    "network",
		Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x01", Slot: "0x00", Function: "0x0"},
	})

	c.AddConsole(Console{Type: "pty", Target: &Target{Type: "serial", Port: "0"}})

	c.AddSerial(Serial{Target: &Target{Type: "isa-serial", Port: "0", Model: &Model{Name: "isa-serial"}}})

	c.AddChannel(Channel{
		Type:   "unix",
		Source: &Source{Mode: "bind"},
		Target: &Target{Type: "virtio", Name: "org.qemu.guest_agent.0"},
	})

	c.AddGraphics(Graphics{
		Type:     "vnc",
		Port:     -1,
		Autoport: "yes",
		Listen:   &Listen{Type: "address"},
	})

	c.SetWatchdog(&Watchdog{Model: "itco", Action: "reset"})

	c.SetVideo(&Video{
		Model:   &Model{Type: "vga", VRAM: 16384, Heads: 1, Primary: "yes"},
		Alias:   &Alias{Name: "video0"},
		Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x00", Slot: "0x01", Function: "0x0"},
	})

	c.SetRNG(&RNG{
		Model:   "virtio",
		Backend: &Backend{Model: "random", CharData: "/dev/urandom"},
		Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x06", Slot: "0x00", Function: "0x0"},
	})

	c.SetMemBalloon(&MemBalloon{
		Model:   "virtio",
		Address: &Address{Type: "pci", Domain: "0x0000", Bus: "0x05", Slot: "0x00", Function: "0x0"},
	})

	c.SetOnReboot(ON_EVENT_ACTION_RESTART)

	return c, nil
}
