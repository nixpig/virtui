package libvirt

import "libvirt.org/go/libvirt"

type DomainEvent struct {
	DomainName string
	Event      int
	Detail     int
}

func RegisterDefaultEventLoop() error {
	return libvirt.EventRegisterDefaultImpl()
}

func RunDefaultEventLoop() error {
	return libvirt.EventRunDefaultImpl()
}