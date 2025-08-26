package libvirtui

import (
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

type Domain struct {
	*libvirt.Domain
	xml libvirtxml.Domain
}

type DomainState uint32

const (
	DomainStateRunning  = DomainState(libvirt.DOMAIN_RUNNING)
	DomainStateBlocked  = DomainState(libvirt.DOMAIN_BLOCKED)
	DomainStatePaused   = DomainState(libvirt.DOMAIN_PAUSED)
	DomainStateShutdown = DomainState(libvirt.DOMAIN_SHUTDOWN)
	DomainStateShutoff  = DomainState(libvirt.DOMAIN_SHUTOFF)
)

func (d *Domain) Name() string {
	return d.xml.Name
}

func (d *Domain) State() string {
	state, _, err := d.GetState()
	if err != nil {
		return ""
	}

	return FromState(state)
}

func (d *Domain) Memory() uint64 {
	info, err := d.GetInfo()
	if err != nil {
		return 0
	}

	return info.MaxMem
}

func (d *Domain) VCPU() uint32 {
	info, err := d.GetInfo()
	if err != nil {
		return 0
	}

	return uint32(info.NrVirtCpu)
}

func ToDomainStruct(domain *libvirt.Domain) (Domain, error) {
	var d Domain
	d.Domain = domain

	doc, err := domain.GetXMLDesc(0)
	if err != nil {
		return d, err
	}

	if err := d.xml.Unmarshal(doc); err != nil {
		return d, err
	}

	return d, nil
}
