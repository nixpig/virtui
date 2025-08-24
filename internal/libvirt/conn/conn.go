package libvirtconn

import (
	"context"

	"libvirt.org/go/libvirt"
)

type LibvirtConnection interface {
	GetHostname() (string, error)
	GetLibVersion() (uint32, error)
	GetVersion() (uint32, error)
	GetType() (string, error)
	GetURI() (string, error)

	Close() (int, error)

	DomainEventLifecycleRegister(
		*libvirt.Domain,
		libvirt.DomainEventLifecycleCallback,
	) (int, error)
}

type connection struct {
	*libvirt.Connect
}

func New(ctx context.Context, qemuURI string) (LibvirtConnection, error) {
	c, err := libvirt.NewConnect(qemuURI)
	if err != nil {
		return nil, err
	}

	return connection{c}, nil
}
