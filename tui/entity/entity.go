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

func ToStorageVolume(vol *libvirt.StorageVol) (StorageVolume, error) {
	var v StorageVolume

	doc, err := vol.GetXMLDesc(0)
	if err != nil {
		return v, err
	}

	if err := v.Unmarshal(doc); err != nil {
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

	if err := p.Unmarshal(doc); err != nil {
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

	if err := n.Unmarshal(doc); err != nil {
		return n, err
	}

	return n, nil
}

func ToDomainStruct(domain *libvirt.Domain) (Domain, error) {
	var d Domain

	doc, err := domain.GetXMLDesc(0)
	if err != nil {
		return d, err
	}

	if err := d.Unmarshal(doc); err != nil {
		return d, err
	}

	return d, nil
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
