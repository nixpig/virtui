package service

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/nixpig/virtui/internal/entity"
	"github.com/nixpig/virtui/internal/mappers"
	"libvirt.org/go/libvirt"
)

type libvirtConnection interface {
	GetHostname() (string, error)
	GetLibVersion() (uint32, error)
	GetVersion() (uint32, error)
	GetType() (string, error)
	GetURI() (string, error)

	DomainEventLifecycleRegister(
		dom *libvirt.Domain,
		callback libvirt.DomainEventLifecycleCallback,
	) (int, error)
	DomainEventDeregister(callbackID int) error

	LookupDomainByUUIDString(uuid string) (*libvirt.Domain, error)
	ListAllDomains(
		flags libvirt.ConnectListAllDomainsFlags,
	) ([]libvirt.Domain, error)
	ListAllNetworks(
		flags libvirt.ConnectListAllNetworksFlags,
	) ([]libvirt.Network, error)
	ListAllStoragePools(
		flags libvirt.ConnectListAllStoragePoolsFlags,
	) ([]libvirt.StoragePool, error)

	Close() (int, error)
}

// libvirtService is the implementation of the Service interface for libvirt.
type libvirtService struct {
	conn libvirtConnection
}

// NewLibvirtService creates a new libvirtService.
func NewLibvirtService(conn libvirtConnection) Service {
	return &libvirtService{conn: conn}
}

func (s *libvirtService) GetConnectionDetails() (entity.ConnectionDetails, error) {
	hostname, err := s.conn.GetHostname()
	if err != nil {
		return entity.ConnectionDetails{}, err
	}

	lvVersion, err := s.conn.GetLibVersion()
	if err != nil {
		return entity.ConnectionDetails{}, err
	}

	hvVersion, err := s.conn.GetVersion()
	if err != nil {
		return entity.ConnectionDetails{}, err
	}

	connType, err := s.conn.GetType()
	if err != nil {
		return entity.ConnectionDetails{}, err
	}

	connURI, err := s.conn.GetURI()
	if err != nil {
		return entity.ConnectionDetails{}, err
	}

	return entity.ConnectionDetails{
		Hostname:  hostname,
		URI:       connURI,
		ConnType:  connType,
		HvVersion: mappers.Version(hvVersion),
		LvVersion: mappers.Version(lvVersion),
	}, nil
}

func (s *libvirtService) RegisterDomainEventCallback(events chan *libvirt.DomainEventLifecycle) (int, error) {
	callbackID, err := s.conn.DomainEventLifecycleRegister(
		nil,
		func(c *libvirt.Connect, d *libvirt.Domain, event *libvirt.DomainEventLifecycle) {
			events <- event
		},
	)
	if err != nil {
		return 0, fmt.Errorf("failed to register domain event handler: %w", err)
	}

	return callbackID, nil
}

func (s *libvirtService) DeregisterDomainEventCallback(callbackID int) error {
	return s.conn.DomainEventDeregister(callbackID)
}

func (s *libvirtService) EventLoop(ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := libvirt.EventRunDefaultImpl(); err != nil {
					log.Error("failed to run event loop", "err", err)
				}
			}
		}
	}()

	return nil
}

func (s *libvirtService) ListAllDomains() ([]entity.DomainWithState, error) {
	domains, err := s.conn.ListAllDomains(0)
	if err != nil {
		return nil, err
	}

	entities := make([]entity.DomainWithState, len(domains))
	for i, domain := range domains {
		d, err := entity.ToDomainStruct(&domain)
		if err != nil {
			return nil, err
		}

		state, _, err := domain.GetState()
		if err != nil {
			return nil, err
		}

		entities[i] = entity.DomainWithState{
			Domain: d,
			State:  state,
		}
	}

	return entities, nil
}

func (s *libvirtService) LookupDomainByUUIDString(uuid string) (entity.Domain, error) {
	domain, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return entity.Domain{}, err
	}

	return entity.ToDomainStruct(domain)
}

func (s *libvirtService) StartDomain(uuid string) error {
	d, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return err
	}
	return d.Create()
}

func (s *libvirtService) PauseResumeDomain(uuid string) error {
	d, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return err
	}

	state, _, err := d.GetState()
	if err != nil {
		return err
	}

	if state == libvirt.DOMAIN_PAUSED {
		return d.Resume()
	} else if state == libvirt.DOMAIN_RUNNING {
		return d.Suspend()
	}

	return nil
}

func (s *libvirtService) ShutdownDomain(uuid string) error {
	d, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return err
	}
	return d.Shutdown()
}

func (s *libvirtService) RebootDomain(uuid string) error {
	d, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return err
	}
	return d.Reboot(0)
}

func (s *libvirtService) ForceResetDomain(uuid string) error {
	d, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return err
	}
	return d.Reset(0)
}

func (s *libvirtService) ForceOffDomain(uuid string) error {
	d, err := s.conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		return err
	}
	return d.Destroy()
}

func (s *libvirtService) ListAllNetworks() ([]entity.Network, error) {
	networks, err := s.conn.ListAllNetworks(0)
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Network, len(networks))
	for i, network := range networks {
		n, err := entity.ToNetworkStruct(&network)
		if err != nil {
			return nil, err
		}
		entities[i] = n
	}

	return entities, nil
}

func (s *libvirtService) ListAllStoragePools() ([]entity.StoragePool, error) {
	pools, err := s.conn.ListAllStoragePools(0)
	if err != nil {
		return nil, err
	}

	entities := make([]entity.StoragePool, len(pools))
	for i, pool := range pools {
		p, err := entity.ToStoragePoolStruct(&pool)
		if err != nil {
			return nil, err
		}
		entities[i] = p
	}

	return entities, nil
}

func (s *libvirtService) Close() (int, error) {
	return s.conn.Close()
}
