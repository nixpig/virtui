package libvirtui

import (
	"libvirt.org/go/libvirt"
	"libvirt.org/go/libvirtxml"
)

// StorageVolume encapsulates the libvirt storage volume and XML representation.
type StorageVolume struct {
	*libvirt.StorageVol
	xml libvirtxml.StorageVolume
}

func (s StorageVolume) Name() string {
	return s.xml.Name
}

func (s StorageVolume) Capacity() (uint64, string) {
	return s.xml.Capacity.Value, s.xml.Capacity.Unit
}

func (s StorageVolume) TargetFormatType() string {
	return s.xml.Target.Format.Type
}

// StoragePool encapsulates the libvirt storage pool and xml representation.
type StoragePool struct {
	*libvirt.StoragePool
	xml libvirtxml.StoragePool
}

func (s StoragePool) Name() string {
	return s.xml.Name
}

func (s StoragePool) UUID() string {
	return s.xml.UUID
}

func (s StoragePool) Type() string {
	return s.xml.Type
}

func (s StoragePool) Capacity() (uint64, string) {
	return s.xml.Capacity.Value, s.xml.Capacity.Unit
}

func (s StoragePool) Available() (uint64, string) {
	return s.xml.Available.Value, s.xml.Available.Unit
}

func (s StoragePool) TargetPath() string {
	return s.xml.Target.Path
}

func ToStorageVolumeStruct(vol *libvirt.StorageVol) (StorageVolume, error) {
	var v StorageVolume
	v.StorageVol = vol

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
	p.StoragePool = pool

	doc, err := pool.GetXMLDesc(0)
	if err != nil {
		return p, err
	}

	if err := p.xml.Unmarshal(doc); err != nil {
		return p, err
	}

	return p, nil
}
