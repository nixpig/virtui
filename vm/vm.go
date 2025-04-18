package vm

import "encoding/xml"

func SerialiseHexInt(n any) error {
	return nil
}

type Readonly struct{}

type DomainType string

const (
	DOMAIN_TYPE_KVM = "kvm"
	// DOMAIN_TYPE_XEN = "xen"
	// DOMAIN_TYPE_QEMU = "qemu"
	// DOMAIN_TYPE_HVF = "hvf"
	// DOMAIN_TYPE_LXC = "lxc"
)

type OSType string

const (
	OS_TYPE_HVM   = "hvm"   // OS designed to run on bare metal
	OS_TYPE_LINUX = "linux" // OS designed to run on Xen 3
)

type OnEventAction string

const (
	ON_EVENT_ACTION_DESTROY        OnEventAction = "destroy"
	ON_EVENT_ACTION_RESTART        OnEventAction = "restart"
	ON_EVENT_ACTION_PRESERVE       OnEventAction = "preserve"
	ON_EVENT_ACTION_RENAME_RESTART OnEventAction = "rename-restart"
)

type EnabledFlag string

const (
	FLAG_ENABLED_YES EnabledFlag = "yes"
	FLAG_ENABLED_NO  EnabledFlag = "no"
)

type DiskType string

const (
	DISK_TYPE_FILE       DiskType = "file"
	DISK_TYPE_BLOCK      DiskType = "block"
	DISK_TYPE_NETWORK    DiskType = "network"
	DISK_TYPE_VOLUME     DiskType = "volume"
	DISK_TYPE_DIR        DiskType = "dir"
	DISK_TYPE_NVME       DiskType = "nvme"
	DISK_TYPE_VHOST_USER DiskType = "vhostuser"
	DISK_TYPE_VHOST_VDPA DiskType = "vhostvdpa"
)

type DiskDevice string

const (
	DISK_DEVICE_FLOPPY DiskDevice = "floppy"
	DISK_DEVICE_DISK   DiskDevice = "disk"
	DISK_DEVICE_CDROM  DiskDevice = "cdrom"
	DISK_DEVICE_LUN    DiskDevice = "lun"
)

type BootDevice string

const (
	BOOT_DEVICE_FLOPPY  BootDevice = "fd"
	BOOT_DEVICE_DISK    BootDevice = "hd"
	BOOT_DEVICE_CDROM   BootDevice = "cdrom"
	BOOT_DEVICE_NETWORK BootDevice = "network"
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
	RebootTimeout int         `xml:"rebootTimeout,attr"`
	UseSerial     EnabledFlag `xml:"useserial,attr,omitempty"`
}

type Boot struct {
	Dev   BootDevice `xml:"dev,attr,omitempty"`
	Order int        `xml:"order,attr"`
}

type BootMenu struct {
	Enable  EnabledFlag `xml:"enable,attr,omitempty"`
	Timeout *int        `xml:"timeout,attr"`
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
	Type     DiskType   `xml:"type,attr"`
	Device   DiskDevice `xml:"device,attr"`
	Driver   *Driver    `xml:"driver"`
	Source   *Source    `xml:"source"`
	Target   *Target    `xml:"target"`
	Address  *Address   `xml:"address"`
	Boot     *Boot      `xml:"boot"`
	Readonly *Readonly  `xml:"readonly"`
}

type Domain struct {
	XMLName       xml.Name       `xml:"domain"`
	Title         string         `xml:"title,omitempty"`
	Description   string         `xml:"description,omitempty"`
	Type          DomainType     `xml:"type,attr,omitempty"`
	Name          string         `xml:"name"`
	UUID          string         `xml:"uuid,omitempty"`
	Metadata      *Metadata      `xml:"metadata"`
	Memory        *Memory        `xml:"memory"`
	CurrentMemory *CurrentMemory `xml:"currentMemory"`
	VCPU          *VCPU          `xml:"vcpu"`
	OS            *OS            `xml:"os"`
	Features      *Features      `xml:"features"`
	Cpu           *CPU           `xml:"cpu"`
	Clock         *Clock         `xml:"clock"`
	Devices       *Devices       `xml:"devices"`
	OnCrash       OnEventAction  `xml:"on_crash,omitempty"`
	OnPoweroff    OnEventAction  `xml:"on_poweroff,omitempty"`
	OnReboot      OnEventAction  `xml:"on_reboot,omitempty"`
	PM            *PM            `xml:"pm"`
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

type Metadata struct {
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
	Enabled EnabledFlag `xml:"enabled,attr"`
}

type SuspendToMem struct {
	Enabled EnabledFlag `xml:"enabled,attr"`
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
	CharData OSType `xml:",chardata"`
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
