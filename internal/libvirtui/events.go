package libvirtui

import "libvirt.org/go/libvirt"

// DomainEvent represents an emitted libvirt domain event to handle.
type DomainEvent struct {
	*libvirt.DomainEventLifecycle
}

// RegisterDefaultEventLoop registers the default implementation of the libvirt
// event loop.
func RegisterDefaultEventLoop() error {
	return libvirt.EventRegisterDefaultImpl()
}

// RunDefaultEventLoop runs the default implementation of the libvirt event loop.
func RunDefaultEventLoop() error {
	return libvirt.EventRunDefaultImpl()
}
