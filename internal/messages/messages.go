package messages

import (
	"github.com/nixpig/virtui/internal/libvirtui"
)

// DomainsMsg is a message to notify of domain list.
type DomainsMsg struct {
	// Domains contains the list of available domains.
	Domains []libvirtui.Domain
}

// ScreenSizeMsg is a message to notify a screen of a size change.
type ScreenSizeMsg struct {
	// Width is the screen size width in terminal characters.
	Width int
	// Height is the screen size height in terminal characters.
	Height int
}

// StoragePoolsMsg is a message to notify of storage pool and volume list.
type StoragePoolsMsg struct {
	// Storage contains the list of available storage pools and their volumes.
	Storage map[libvirtui.StoragePool][]libvirtui.StorageVolume
}

// NetworksMsg is a message to notify of network list.
type NetworksMsg struct {
	// Networks contains the list of available networks.
	Networks []libvirtui.Network
}
