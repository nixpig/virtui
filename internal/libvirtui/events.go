package libvirtui

import "libvirt.org/go/libvirt"

type DomainEvent struct {
	*libvirt.DomainEventLifecycle
}

func RegisterDefaultEventLoop() error {
	return libvirt.EventRegisterDefaultImpl()
}

func RunDefaultEventLoop() error {
	return libvirt.EventRunDefaultImpl()
}
