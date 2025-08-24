package entity

import "libvirt.org/go/libvirt"

// DomainWithState combines a domain with its live state.
type DomainWithState struct {
	Domain Domain
	State  libvirt.DomainState
}
