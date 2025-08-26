package libvirtui

import (
	"context"

	"libvirt.org/go/libvirt"
)

// Connection is the abstraction around the underlying libvirt API connection.
type Connection interface {
	GetHostname() (string, error)
	GetLibVersion() (uint32, error)
	GetVersion() (uint32, error)
	GetType() (string, error)
	GetURI() (string, error)

	Close() (int, error)

	DomainEventLifecycleRegister(cb func(DomainEvent)) (int, error)
	DomainEventLifecycleDeregister(callbackID int) error

	LookupDomainByUUIDString(uuid string) (*libvirt.Domain, error)

	ListAllDomains(
		flags libvirt.ConnectListAllDomainsFlags,
	) ([]libvirt.Domain, error)

	DefineDomainFlags(
		xml string,
		flags libvirt.DomainDefineFlags,
	) (*libvirt.Domain, error)

	ListAllStoragePools(
		flags libvirt.ConnectListAllStoragePoolsFlags,
	) ([]libvirt.StoragePool, error)

	ListAllNetworks(
		flags libvirt.ConnectListAllNetworksFlags,
	) ([]libvirt.Network, error)
}

type connection struct {
	*libvirt.Connect
}

// NewConnection returns a new Connection for the provided QEMU server URI.
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

func (c connection) ListAllStoragePools(
	flags libvirt.ConnectListAllStoragePoolsFlags,
) ([]libvirt.StoragePool, error) {
	return c.Connect.ListAllStoragePools(flags)
}

func (c connection) ListAllNetworks(
	flags libvirt.ConnectListAllNetworksFlags,
) ([]libvirt.Network, error) {
	return c.Connect.ListAllNetworks(flags)
}
