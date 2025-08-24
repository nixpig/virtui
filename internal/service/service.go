package service

import (
	"context"

	"github.com/nixpig/virtui/internal/entity"
	"libvirt.org/go/libvirt"
)

// Service defines the interface for interacting with the hypervisor.
type Service interface {
	// Connection details
	GetConnectionDetails() (entity.ConnectionDetails, error)

	// Event handling
	RegisterDomainEventCallback(events chan *libvirt.DomainEventLifecycle) (int, error)
	DeregisterDomainEventCallback(callbackID int) error
	EventLoop(ctx context.Context) error

	// Domain management
	ListAllDomains() ([]entity.DomainWithState, error)
	LookupDomainByUUIDString(uuid string) (entity.Domain, error)
	StartDomain(uuid string) error
	PauseResumeDomain(uuid string) error
	ShutdownDomain(uuid string) error
	RebootDomain(uuid string) error
	ForceResetDomain(uuid string) error
	ForceOffDomain(uuid string) error

	// Network management
	ListAllNetworks() ([]entity.Network, error)

	// Storage management
	ListAllStoragePools() ([]entity.StoragePool, error)

	// Close the connection
	Close() (int, error)
}
