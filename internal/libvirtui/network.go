package libvirtui

import (
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

// Network encapsulates a libvirt network and xml representation.
type Network struct {
	*libvirt.Network
	xml libvirtxml.Network
}

func (n Network) Name() string {
	return n.xml.Name
}

func (n Network) UUID() string {
	return n.xml.UUID
}

func (n Network) Bridge() string {
	return n.xml.Bridge.Name
}

func (n Network) Active() bool {
	state, err := n.IsActive()
	if err != nil {
		return false
	}

	return state
}

func ToNetworkStruct(network *libvirt.Network) (Network, error) {
	var n Network
	n.Network = network

	doc, err := network.GetXMLDesc(0)
	if err != nil {
		return n, err
	}

	if err := n.xml.Unmarshal(doc); err != nil {
		return n, err
	}

	return n, nil
}
