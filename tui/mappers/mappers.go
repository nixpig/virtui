package mappers

import (
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
