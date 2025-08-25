package libvirtui

import (
	"context"

	"libvirt.org/go/libvirt"
)

type Connection interface {
	GetHostname() (string, error)
	GetLibVersion() (uint32, error)
	GetVersion() (uint32, error)
	GetType() (string, error)
	GetURI() (string, error)

	Close() (int, error)

	DomainEventLifecycleRegister(func(DomainEvent)) (int, error)
	DomainEventLifecycleDeregister(callbackID int) error

	LookupDomainByUUIDString(uuid string) (*libvirt.Domain, error)
	ListAllDomains(
		flags libvirt.ConnectListAllDomainsFlags,
	) ([]libvirt.Domain, error)
	DefineDomainFlags(
		xml string,
		flags libvirt.DomainDefineFlags,
	) (*libvirt.Domain, error)
}

type connection struct {
	*libvirt.Connect
}

func NewConnection(ctx context.Context, qemuURI string) (Connection, error) {
	c, err := libvirt.NewConnect(qemuURI)
	if err != nil {
		return nil, err
	}

	return connection{c}, nil
}

func (c connection) DefineDomainFlags(
	xml string,
	flags libvirt.DomainDefineFlags,
) (*libvirt.Domain, error) {
	return c.Connect.DomainDefineXMLFlags(xml, flags)
}

func (c connection) DomainEventLifecycleRegister(
	cb func(DomainEvent),
) (int, error) {
	return c.Connect.DomainEventLifecycleRegister(
		nil,
		func(_ *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventLifecycle) {
			cb(DomainEvent{event})
		},
	)
}

func (c connection) DomainEventLifecycleDeregister(callbackID int) error {
	return c.Connect.DomainEventDeregister(callbackID)
}
