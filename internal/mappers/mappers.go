package mappers

import (
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

func ToStructXML(domain *libvirt.Domain) (*libvirtxml.Domain, error) {
	doc, err := domain.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	dom := &libvirtxml.Domain{}

	if err := dom.Unmarshal(doc); err != nil {
		return nil, err
	}

	return dom, nil
}
