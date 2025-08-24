package events

import "libvirt.org/go/libvirt"

func RegisterDefaultEventLoopImpl() error {
	return libvirt.EventRegisterDefaultImpl()
}

func RunDefaultEventLoopImpl() error {
	return libvirt.EventRunDefaultImpl()
}
