package event

import "libvirt.org/go/libvirt"

type VM struct {
	ID    int
	Event libvirt.DomainEventType
}
