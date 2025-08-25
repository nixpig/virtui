package libvirtui

import (
	"fmt"

	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

type Domain struct {
	*libvirt.Domain
	xml    libvirtxml.Domain
	Name   string
	State  string
	Memory uint64
	VCPU   uint32
}

type DomainState uint32

const (
	DomainStateRunning  = DomainState(libvirt.DOMAIN_RUNNING)
	DomainStateBlocked  = DomainState(libvirt.DOMAIN_BLOCKED)
	DomainStatePaused   = DomainState(libvirt.DOMAIN_PAUSED)
	DomainStateShutdown = DomainState(libvirt.DOMAIN_SHUTDOWN)
	DomainStateShutoff  = DomainState(libvirt.DOMAIN_SHUTOFF)
)

type Network struct {
	libvirt.Network
	xml libvirtxml.Network
}

type StorageVolume struct {
	libvirt.StorageVol
	xml libvirtxml.StorageVolume
}

type StoragePool struct {
	libvirt.StoragePool
	xml libvirtxml.StoragePool
}

func ToStorageVolumeStruct(vol *libvirt.StorageVol) (StorageVolume, error) {
	var v StorageVolume

	doc, err := vol.GetXMLDesc(0)
	if err != nil {
		return v, err
	}

	if err := v.xml.Unmarshal(doc); err != nil {
		return v, err
	}

	return v, nil
}

func ToStoragePoolStruct(pool *libvirt.StoragePool) (StoragePool, error) {
	var p StoragePool

	doc, err := pool.GetXMLDesc(0)
	if err != nil {
		return p, err
	}

	if err := p.xml.Unmarshal(doc); err != nil {
		return p, err
	}

	return p, nil
}

func ToNetworkStruct(network *libvirt.Network) (Network, error) {
	var n Network

	doc, err := network.GetXMLDesc(0)
	if err != nil {
		return n, err
	}

	if err := n.xml.Unmarshal(doc); err != nil {
		return n, err
	}

	return n, nil
}

func ToDomainStruct(domain *libvirt.Domain) (Domain, error) {
	var d Domain
	d.Domain = domain

	name, err := domain.GetName()
	if err != nil {
		return d, fmt.Errorf("get domain name: %w", err)
	}
	d.Name = name

	state, _, err := domain.GetState()
	if err != nil {
		return d, fmt.Errorf("get domain state: %w", err)
	}
	d.State = FromState(state)

	info, err := domain.GetInfo()
	if err != nil {
		return d, fmt.Errorf("get domain info: %w", err)
	}
	d.Memory = info.MaxMem
	d.VCPU = uint32(info.NrVirtCpu)

	doc, err := domain.GetXMLDesc(0)
	if err != nil {
		return d, err
	}

	if err := d.xml.Unmarshal(doc); err != nil {
		return d, err
	}

	return d, nil
}
