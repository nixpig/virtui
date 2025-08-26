package messages

import (
	"github.com/nixpig/virtui/internal/libvirtui"
)

// DomainsMsg is a message to notify of domain list.
type DomainsMsg struct {
	Domains []libvirtui.Domain
}

// ScreenSizeMsg is a message to notify a screen of a size change.
type ScreenSizeMsg struct {
	Width  int
	Height int
}

// StoragePoolsMsg is a message to notify of storage pool list.
type StoragePoolsMsg struct {
	Pools []libvirtui.StoragePool
}

// StorageVolumesMsg is a message to notify of storage volume list.
type StorageVolumesMsg struct {
	Volumes []libvirtui.StorageVolume
}

// StorageVolumeDetailsMsg is a message to notify of storage volume details.
type StorageVolumeDetailsMsg struct {
	Details []string
}

// NetworksMsg is a message to notify of network list.
type NetworksMsg struct {
	Networks []libvirtui.Network
}

// ErrorMsg is a message to notify of an error.
type ErrorMsg struct {
	Err error
}
