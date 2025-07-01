package mappers

import (
	"fmt"

	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

var states = map[libvirt.DomainState]string{
	libvirt.DOMAIN_NOSTATE:     "None",
	libvirt.DOMAIN_RUNNING:     "Running",
	libvirt.DOMAIN_BLOCKED:     "Blocked",
	libvirt.DOMAIN_PAUSED:      "Paused",
	libvirt.DOMAIN_SHUTDOWN:    "Shutdown",
	libvirt.DOMAIN_CRASHED:     "Crashed",
	libvirt.DOMAIN_PMSUSPENDED: "Suspended",
	libvirt.DOMAIN_SHUTOFF:     "Shutoff",
}

func FromState(domainState libvirt.DomainState) string {
	s, ok := states[domainState]
	if !ok {
		return ""
	}

	return s
}

func ToStructXML(domain *libvirt.Domain) (libvirtxml.Domain, error) {
	var dom libvirtxml.Domain

	doc, err := domain.GetXMLDesc(0)
	if err != nil {
		return dom, err
	}

	if err := dom.Unmarshal(doc); err != nil {
		return dom, err
	}

	return dom, nil
}

func Version(v uint32) string {
	major := v / 1_000_000
	minor := (v - (major * 1_000_000)) / 1000
	patch := v - (major * 1_000_000) - (minor * 1000)

	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
