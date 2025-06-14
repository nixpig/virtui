package entity

import (
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

type Domain struct {
	libvirtxml.Domain
}

type Network struct {
	libvirtxml.Network
}

type StorageVolume struct {
	libvirtxml.StorageVolume
}

type StoragePool struct {
	libvirtxml.StoragePool
}

func ToStorageVolume(vol *libvirt.StorageVol) (*StorageVolume, error) {
	doc, err := vol.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	v := &StorageVolume{}

	if err := v.Unmarshal(doc); err != nil {
		return nil, err
	}

	return v, nil
}

func ToStoragePoolStruct(pool *libvirt.StoragePool) (*StoragePool, error) {
	doc, err := pool.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	p := &StoragePool{}

	if err := p.Unmarshal(doc); err != nil {
		return nil, err
	}

	return p, nil
}

func ToNetworkStruct(network *libvirt.Network) (*Network, error) {
	doc, err := network.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	n := &Network{}

	if err := n.Unmarshal(doc); err != nil {
		return nil, err
	}

	return n, nil
}

func ToDomainStruct(domain *libvirt.Domain) (*Domain, error) {
	doc, err := domain.GetXMLDesc(0)
	if err != nil {
		return nil, err
	}

	dom := &Domain{}

	if err := dom.Unmarshal(doc); err != nil {
		return nil, err
	}

	return dom, nil
}

func misc() {
	// dom := &libvirtxml.Domain{
	// 	Type: "kvm",
	// 	Name: "demo",
	// 	UUID: uuid.NewString(),
	// }
	//
	// domXML, err := dom.Marshal()
	// if err != nil {
	// 	slog.Error("failed to marshal domain xml", "err", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(domXML)

	// conn, err := libvirt.NewConnect(qemuURL)
	// if err != nil {
	// 	slog.Error("failed to connect", "qemuURL", qemuURL, "err", err)
	// 	os.Exit(1)
	// }
	//
	// fmt.Println("DOMAINS: ")
	// doms, err := conn.ListAllDomains(0)
	// if err != nil {
	// 	slog.Error("failed to list all active domains", "err", err)
	// 	os.Exit(1)
	// }
	//
	// for _, d := range doms {
	// 	defer d.Free()
	//
	// 	s, err := mappers.ToStructXML(&d)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	//
	// 	fmt.Printf("%+v\n", s)
	// }
}
