package domain

import (
	"encoding/xml"
)

type Domain struct {
	XMLName       xml.Name  `xml:"domain"`
	Type          string    `xml:"type,attr"`
	Name          string    `xml:"name"`
	UUID          string    `xml:"uuid"`
	Metadata      *Metadata `xml:"metadata"`
	Memory        *Memory   `xml:"memory"`
	CurrentMemory *Memory   `xml:"currentMemory"`
	VCPU          *VCPU     `xml:"vcpu"`
	OS            *OS       `xml:"os"`
	Features      *Features `xml:"features"`
	CPU           *CPU      `xml:"cpu"`
	Clock         *Clock    `xml:"clock"`
	PM            *PM       `xml:"pm"`
	Devices       *Devices  `xml:"devices"`
}

type Metadata struct {
	XMLName xml.Name `xml:"metadata"`
	// TODO: namespace stuff
}

type Memory int64

type VCPU int32

type OS struct {
	XMLName xml.Name      `xml:"os"`
	Type    *OSType       `xml:"type"`
	Kernel  string        `xml:"kernel"`
	Initrd  string        `xml:"initrd"`
	Cmdline *KernelOption `xml:"cmdline"`
}

type OSType struct {
	XMLName xml.Name `xml:"type"`
	Arch    string   `xml:"arch,attr"`
	Machine string   `xml:"machine,attr"`
	Value   string   `xml:",chardata"`
}

type KernelOption string

type Features struct {
	XMLName xml.Name `xml:"features"`
	ACPI    string   `xml:"acpi"`
	APIC    string   `xml:"apic"`
}

type CPU struct {
	XMLName xml.Name `xml:"cpu"`
	Mode    string   `xml:"mode,attr"`
}

type Clock struct {
	XMLName xml.Name `xml:"clock"`
	Offset  string   `xml:"offset,attr"`
	Timers  []Timer  `xml:"timer"`
}

type Timer struct {
	XMLName    xml.Name `xml:"timer"`
	Name       string   `xml:"name,attr"`
	TickPolicy string   `xml:"tickpolicy,attr,omitempty"`
	Present    string   `xml:"present,attr,omitempty"`
}

type PM struct {
	XMLName       xml.Name `xml:"pm"`
	SuspendToMem  *Suspend `xml:"suspend-to-mem"`
	SuspendToDisk *Suspend `xml:"suspend-to-disk"`
}

type Suspend struct {
	Enabled string `xml:"enabled,attr,omitempty"`
}

type Devices struct {
	XMLName  xml.Name `xml:"devices"`
	Emulator string   `xml:"emulator"`
	Disks    []Disk   `xml:"disk,omitempty"`
}

type Disk struct {
	XMLName xml.Name `xml:"disk"`
	Type    string   `xml:"type,attr"`
	Device  string   `xml:"device,attr"`
	Driver  *Driver  `xml:"driver"`
	Source  *Source  `xml:"source"`
	Target  *Target  `xml:"target"`
}

type Driver struct {
	XMLName xml.Name `xml:"driver"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

type Source struct {
	XMLName xml.Name `xml:"source"`
	File    string   `xml:"file,attr"`
}

type Target struct {
	XMLName xml.Name `xml:"target"`
	Dev     string   `xml:"dev,attr"`
	Bus     string   `xml:"bus,attr"`
}
