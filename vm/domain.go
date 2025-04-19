package vm

import (
	"encoding/xml"
	"io"

	"github.com/google/uuid"
)

type Readonly struct{}

const (
	MACHINE_TYPE_Q35 = "q35"
	ARCH_X86_64      = "x86_64"
	DOMAIN_TYPE_KVM  = "kvm"
	OS_TYPE_HVM      = "hvm"

	ON_EVENT_ACTION_DESTROY        = "destroy"
	ON_EVENT_ACTION_RESTART        = "restart"
	ON_EVENT_ACTION_PRESERVE       = "preserve"
	ON_EVENT_ACTION_RENAME_RESTART = "rename-restart"

	FLAG_ENABLED_YES = "yes"
	FLAG_ENABLED_NO  = "no"

	DISK_DEVICE_FLOPPY = "floppy"
	DISK_DEVICE_DISK   = "disk"
	DISK_DEVICE_CDROM  = "cdrom"
	DISK_DEVICE_LUN    = "lun"

	BOOT_DEVICE_FLOPPY  = "fd"
	BOOT_DEVICE_DISK    = "hd"
	BOOT_DEVICE_CDROM   = "cdrom"
	BOOT_DEVICE_NETWORK = "network"
)

type ACPI struct{}

type APIC struct{}

type Address struct {
	Base string `xml:"base,attr,omitempty"`
	// TODO: this should be hex, e.g. 0x0000
	Bus        string `xml:"bus,attr,omitempty"`
	Controller string `xml:"controller,attr,omitempty"`
	// TODO: this should be hex, e.g. 0x0000
	Domain string `xml:"domain,attr,omitempty"`
	// TODO: this should be hex, e.g. 0x0000
	Function      string `xml:"function,attr,omitempty"`
	Multifunction string `xml:"multifunction,attr,omitempty"`
	Port          *int   `xml:"port,attr"`
	// TODO: this should be hex, e.g. 0x0000
	Slot   string `xml:"slot,attr,omitempty"`
	Target string `xml:"target,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"`
	Unit   string `xml:"unit,attr,omitempty"`
}

type Alias struct {
	Name string `xml:"name,attr,omitempty"`
}

type Backend struct {
	CharData string `xml:",chardata"`
	Model    string `xml:"model,attr,omitempty"`
}

type BIOS struct {
	RebootTimeout int    `xml:"rebootTimeout,attr"`
	UseSerial     string `xml:"useserial,attr,omitempty"`
}

type Boot struct {
	Dev   string `xml:"dev,attr,omitempty"`
	Order int    `xml:"order,attr"`
}

type BootMenu struct {
	Enable  string `xml:"enable,attr,omitempty"`
	Timeout *int   `xml:"timeout,attr"`
}

type Channel struct {
	Type   string  `xml:"type,attr,omitempty"`
	Source *Source `xml:"source"`
	Target *Target `xml:"target"`
}

type Clock struct {
	Offset string  `xml:"offset,attr,omitempty"`
	Sync   string  `xml:"sync,attr,omitempty"`
	Timers []Timer `xml:"timer"`
}

type Console struct {
	Type   string  `xml:"type,attr"`
	Source *Source `xml:"source"`
	Target *Target `xml:"target"`
}

type Controller struct {
	Index   int      `xml:"index,attr"`
	Model   string   `xml:"model,attr,omitempty"`
	Ports   *int     `xml:"ports,attr"`
	Type    string   `xml:"type,attr,omitempty"`
	Address *Address `xml:"address"`
	Target  *Target  `xml:"target"`
}

type CPU struct {
	Migratable string `xml:"migratable,attr,omitempty"`
	Mode       string `xml:"mode,attr,omitempty"`
	Check      string `xml:"check,attr,omitempty"`
}

type CurrentMemory struct {
	CharData string `xml:",chardata"`
	Unit     string `xml:"unit,attr,omitempty"`
}

type Devices struct {
	Emulator    string       `xml:"emulator,omitempty"`
	Disks       []Disk       `xml:"disk"`
	Controllers []Controller `xml:"controller"`
	Interfaces  []Interface  `xml:"interface"`
	Serials     []Serial     `xml:"serial"`
	Consoles    []Console    `xml:"console"`
	Channels    []Channel    `xml:"channel"`
	Video       *Video       `xml:"video"`
	Watchdog    *Watchdog    `xml:"watchdog"`
	MemBalloon  *MemBalloon  `xml:"memballoon"`
	RNG         *RNG         `xml:"rng"`
	Graphics    []Graphics   `xml:"graphics"`
}

type Disk struct {
	Type     string    `xml:"type,attr"`
	Device   string    `xml:"device,attr"`
	Driver   *Driver   `xml:"driver"`
	Source   *Source   `xml:"source"`
	Target   *Target   `xml:"target"`
	Address  *Address  `xml:"address"`
	Boot     *Boot     `xml:"boot"`
	Readonly *Readonly `xml:"readonly"`
}

type Domain struct {
	XMLName       xml.Name        `xml:"domain"`
	Title         string          `xml:"title,omitempty"`
	Description   string          `xml:"description,omitempty"`
	Type          string          `xml:"type,attr,omitempty"`
	Name          string          `xml:"name"`
	UUID          string          `xml:"uuid,omitempty"`
	Metadata      *DomainMetadata `xml:"metadata"`
	Memory        *Memory         `xml:"memory"`
	CurrentMemory *CurrentMemory  `xml:"currentMemory"`
	VCPU          *VCPU           `xml:"vcpu"`
	OS            *OS             `xml:"os"`
	Features      *Features       `xml:"features"`
	CPU           *CPU            `xml:"cpu"`
	Clock         *Clock          `xml:"clock"`
	Devices       *Devices        `xml:"devices"`
	OnCrash       string          `xml:"on_crash,omitempty"`
	OnPoweroff    string          `xml:"on_poweroff,omitempty"`
	OnReboot      string          `xml:"on_reboot,omitempty"`
	PM            *PM             `xml:"pm"`
}

type Driver struct {
	Discard string `xml:"discard,attr,omitempty"`
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
}

type Features struct {
	ACPI *ACPI `xml:"acpi"`
	APIC *APIC `xml:"apic"`
}

type Graphics struct {
	Type     string  `xml:"type,attr"`
	Port     int     `xml:"port,attr"`
	Autoport string  `xml:"autoport,attr,omitempty"`
	Listen   *Listen `xml:"listen"`
}

type Interface struct {
	Type    string   `xml:"type,attr,omitempty"`
	Address *Address `xml:"address"`
	Mac     *Mac     `xml:"mac"`
	Model   *Model   `xml:"model"`
	Source  *Source  `xml:"source"`
}

type LibOSInfo struct {
	LibOSInfo string `xml:"xmlns:libosinfo,attr"`
	OS        *OS    `xml:"libosinfo:os"`
}

type Listen struct {
	Network string `xml:"network,attr"`
	Type    string `xml:"type,attr"`
}

type Mac struct {
	Address string `xml:"address,attr"`
}

type MemBalloon struct {
	Model   string   `xml:"model,attr"`
	Address *Address `xml:"address"`
	Driver  *Driver  `xml:"driver"`
}

type Memory struct {
	CharData string `xml:",chardata"`
	Unit     string `xml:"unit,attr,omitempty"`
}

type DomainMetadata struct {
	LibOSInfo *LibOSInfo `xml:"libosinfo:libosinfo"`
}

type Model struct {
	Heads    int    `xml:"heads,attr"`
	Name     string `xml:"name,attr,omitempty"`
	Type     string `xml:"type,attr"`
	VRAM     int    `xml:"vram,attr"`
	CharData string `xml:",chardata"`
	Primary  string `xml:"primary,attr,omitempty"`
}

type OS struct {
	ID             string    `xml:"id,attr,omitempty"`
	ACPI           *ACPI     `xml:"acpi"`
	BIOS           *BIOS     `xml:"bios"`
	Boot           *Boot     `xml:"boot"`
	Bootloader     string    `xml:"bootloader,omitempty"`
	BootloaderArgs string    `xml:"bootloader_args,omitempty"`
	BootMenu       *BootMenu `xml:"bootmenu"`
	Cmdline        string    `xml:"cmdline,omitempty"`
	Init           string    `xml:"init,omitempty"`
	InitArgs       []string  `xml:"initarg"`
	InitDir        string    `xml:"initdir,omitempty"`
	InitGroup      *int      `xml:"initgroup"`
	InitRD         string    `xml:"initrd,omitempty"`
	InitUser       string    `xml:"inituser,omitempty"`
	Kernel         string    `xml:"kernel,omitempty"`
	Type           *Type     `xml:"type"`
}

type PM struct {
	SuspendToDisk *SuspendToDisk `xml:"suspend-to-disk"`
	SuspendToMem  *SuspendToMem  `xml:"suspend-to-mem"`
}

type RNG struct {
	Model   string   `xml:"model,attr"`
	Backend *Backend `xml:"backend"`
	Address *Address `xml:"address"`
}

type Serial struct {
	Target *Target `xml:"target"`
}

type Source struct {
	File    string `xml:"file,attr,omitempty"`
	Mode    string `xml:"mode,attr,omitempty"`
	Network string `xml:"network,attr,omitempty"`
}

type SuspendToDisk struct {
	Enabled string `xml:"enabled,attr"`
}

type SuspendToMem struct {
	Enabled string `xml:"enabled,attr"`
}

type Target struct {
	Bus  string `xml:"bus,attr,omitempty"`
	Dev  string `xml:"dev,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	// TODO: this should be hex, e.g. 0x0000
	Port    string `xml:"port,attr,omitempty"`
	Type    string `xml:"type,attr,omitempty"`
	Model   *Model `xml:"model"`
	Chassis int    `xml:"chassis,attr,omitzero"`
}

type Timer struct {
	Name       string `xml:"name,attr"`
	Present    string `xml:"present,attr,omitempty"`
	Tickpolicy string `xml:"tickpolicy,attr,omitempty"`
}

type Type struct {
	Arch     string `xml:"arch,attr,omitempty"`
	Machine  string `xml:"machine,attr,omitempty"`
	CharData string `xml:",chardata"`
}

type VCPU struct {
	CharData  string `xml:",chardata"`
	Placement string `xml:"placement,attr,omitempty"`
}

type Video struct {
	Alias   *Alias   `xml:"alias"`
	Model   *Model   `xml:"model"`
	Address *Address `xml:"address"`
}

type Watchdog struct {
	Action string `xml:"action,attr,omitempty"`
	Model  string `xml:"model,attr,omitempty"`
}

func NewDomainFromXML(xml string) (*Domain, error) {
	return nil, nil
}

func NewDomainFromFile(r io.Reader) (*Domain, error) {
	return nil, nil
}

func NewDomain(name string) *Domain {
	return &Domain{
		Name:     name,
		UUID:     uuid.NewString(),
		Metadata: &DomainMetadata{LibOSInfo: &LibOSInfo{LibOSInfo: "", OS: &OS{}}},
		OS:       &OS{},
		Features: &Features{},
		Clock:    &Clock{},
		PM:       &PM{SuspendToDisk: &SuspendToDisk{}, SuspendToMem: &SuspendToMem{}},
		Devices:  &Devices{},
	}
}

func (d *Domain) ToXML() ([]byte, error) {
	return xml.Marshal(d)
}

func (d *Domain) ToXMLFormatted() ([]byte, error) {
	return xml.MarshalIndent(d, "", "  ")
}

func (d *Domain) SetName(n string) {
	d.Name = n
}

func (d *Domain) SetUUID(u string) {
	d.UUID = u
}

func (d *Domain) SetMetaData(m *DomainMetadata) {
	d.Metadata = m
}

func (d *Domain) SetMemory(m *Memory) {
	d.Memory = m
}

func (d *Domain) SetCurrentMemory(c *CurrentMemory) {
	d.CurrentMemory = c
}

func (d *Domain) SetVCPU(v *VCPU) {
	d.VCPU = v
}

func (d *Domain) SetOSType(t *Type) {
	d.OS.Type = t
}

func (d *Domain) SetBootMenu(b *BootMenu) {
	d.OS.BootMenu = b
}

func (d *Domain) EnableACPI() {
	d.Features.ACPI = &ACPI{}
}

func (d *Domain) DisableACPI() {
	d.Features.ACPI = nil
}

func (d *Domain) EnableAPIC() {
	d.Features.APIC = &APIC{}
}

func (d *Domain) DisableAPIC() {
	d.Features.APIC = nil
}

func (d *Domain) SetCPU(c *CPU) {
	d.CPU = c
}

func (d *Domain) SetClockOffset(o string) {
	d.Clock.Offset = o
}

func (d *Domain) AddClockTimer(t Timer) {
	d.Clock.Timers = append(d.Clock.Timers, t)
}

func (d *Domain) AddClockTimers(t ...Timer) {
	for _, x := range t {
		d.Clock.Timers = append(d.Clock.Timers, x)
	}
}

func (d *Domain) EnableSuspendToMemory() {
	d.PM.SuspendToMem.Enabled = FLAG_ENABLED_YES
}

func (d *Domain) DisableSuspendToMemory() {
	d.PM.SuspendToMem.Enabled = FLAG_ENABLED_NO
}

func (d *Domain) EnableSuspendToDisk() {
	d.PM.SuspendToDisk.Enabled = FLAG_ENABLED_YES
}

func (d *Domain) DisableSuspendToDisk() {
	d.PM.SuspendToDisk.Enabled = FLAG_ENABLED_NO
}

func (d *Domain) SetEmulator(e string) {
	d.Devices.Emulator = e
}

func (d *Domain) AddDisk(s Disk) {
	d.Devices.Disks = append(d.Devices.Disks, s)
}

func (d *Domain) AddDisks(s ...Disk) {
	for _, x := range s {
		d.Devices.Disks = append(d.Devices.Disks, x)
	}
}

func (d *Domain) AddController(c Controller) {
	d.Devices.Controllers = append(d.Devices.Controllers, c)
}

func (d *Domain) AddControllers(c ...Controller) {
	for _, x := range c {
		d.Devices.Controllers = append(d.Devices.Controllers, x)
	}
}

func (d *Domain) AddInterface(i Interface) {
	d.Devices.Interfaces = append(d.Devices.Interfaces, i)
}

func (d *Domain) AddInterfaces(i ...Interface) {
	for _, x := range i {
		d.Devices.Interfaces = append(d.Devices.Interfaces, x)
	}
}

func (d *Domain) AddSerial(s Serial) {
	d.Devices.Serials = append(d.Devices.Serials, s)
}

func (d *Domain) AddConsole(c Console) {
	d.Devices.Consoles = append(d.Devices.Consoles, c)
}

func (d *Domain) AddChannel(c Channel) {
	d.Devices.Channels = append(d.Devices.Channels, c)
}

func (d *Domain) AddGraphics(g Graphics) {
	d.Devices.Graphics = append(d.Devices.Graphics, g)
}

func (d *Domain) SetVideo(v *Video) {
	d.Devices.Video = v
}

func (d *Domain) SetWatchdog(w *Watchdog) {
	d.Devices.Watchdog = w
}

func (d *Domain) SetRNG(r *RNG) {
	d.Devices.RNG = r
}

func (d *Domain) SetMemBalloon(m *MemBalloon) {
	d.Devices.MemBalloon = m
}

func (d *Domain) SetOnReboot(o string) {
	d.OnReboot = o
}

func (d *Domain) SetOnPoweroff(o string) {
	d.OnPoweroff = o
}

func (d *Domain) SetOnCrash(o string) {
	d.OnCrash = o
}

func (d *Domain) SetLibOSInfo(l string) {
	d.Metadata.LibOSInfo.LibOSInfo = l
}

func (d *Domain) SetLibOSID(l string) {
	d.Metadata.LibOSInfo.OS.ID = l
}

func NewDomainWithDefaults(name string) *Domain {
	c := NewDomain(name)

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
			Source: &Source{File: "/var/lib/libvirt/images/default-vm.qcow2"},
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
		Source:  &Source{Network: "default1"},
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

	return c
}
