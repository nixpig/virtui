package api

import (
	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
	"github.com/nixpig/virtui/vm/domain"
	"github.com/nixpig/virtui/vm/network"
	"github.com/nixpig/virtui/vm/pool"
	"github.com/nixpig/virtui/vm/volume"
)

type API struct {
	conn *libvirt.Libvirt
}

func (a *API) NewAPI(conn *libvirt.Libvirt) *API {
	return &API{conn}
}

func (a *API) CreateNetwork(n *network.Network) (*network.Network, error) {
	return &network.Network{}, nil
}

func (a *API) GetNetworks() []network.Network {
	return []network.Network{}
}

func (a *API) GetNetworkByUUID(u string) (*network.Network, error) {
	return &network.Network{}, nil
}

func (a *API) UpdateNetwork(n *network.Network) (*network.Network, error) {
	return &network.Network{}, nil
}

func (a *API) DeleteNetworkByUUID(u string) error {
	return nil
}

func (a *API) CreateVolume(v *volume.Volume) (*volume.Volume, error) {
	return &volume.Volume{}, nil
}

func (a *API) GetVolumes() []volume.Volume {
	return []volume.Volume{}
}

func (a *API) GetVolumeByPath(p string) (*volume.Volume, error) {
	return &volume.Volume{}, nil
}

func (a *API) UpdateVolume(v *volume.Volume) (*volume.Volume, error) {
	return &volume.Volume{}, nil
}

func (a *API) DeleteVolumeByPath(p string) error {
	return nil
}

func (a *API) CreateDomain(d *domain.Domain) (*domain.Domain, error) {
	return &domain.Domain{}, nil
}

func (a *API) GetDomains() []domain.Domain {
	return []domain.Domain{}
}

func (a *API) GetDomainByUUID(u string) (*domain.Domain, error) {
	b, _ := uuid.FromBytes([]byte(u))
	d, err := a.conn.DomainLookupByUUID(libvirt.UUID(b))
	if err != nil {
		return nil, err
	}

	return domain.FromLibvirt(&d), nil
}

func (a *API) UpdateDomain(d *domain.Domain) (*domain.Domain, error) {
	return &domain.Domain{}, nil
}

func (a *API) DeleteDomainByUUID(u string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainUndefine(*domain.ToLibvirt(d)); err != nil {
		return err
	}

	return nil
}

func (a *API) CreatePool(p *pool.Pool) (*pool.Pool, error) {
	return &pool.Pool{}, nil
}

func (a *API) GetPools() []pool.Pool {
	return []pool.Pool{}
}

func (a *API) GetPoolByUUID(u string) (*pool.Pool, error) {
	return &pool.Pool{}, nil
}

func (a *API) UpdatePool(p *pool.Pool) (*pool.Pool, error) {
	return &pool.Pool{}, nil
}

func (a *API) DeletePoolByUUID(u string) error {
	return nil
}

func (a *API) StartDomainByUUID(u string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainCreate(*domain.ToLibvirt(d)); err != nil {
		return err
	}

	return nil
}

func (a *API) PauseDomainByUUID(u string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainResume(*domain.ToLibvirt(d)); err != nil {
		return err
	}

	return nil
}

func (a *API) ResumeDomainByUUID(u string) error {
	return nil
}

func (a *API) ShutdownDomainByUUID(u string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainShutdown(*domain.ToLibvirt(d)); err != nil {
		return err
	}

	return nil
}

func (a *API) RebootDomainByUUID(u string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainReboot(*domain.ToLibvirt(d), 0); err != nil {
		return err
	}

	return nil
}

func (a *API) ResetDomainByUUID(u string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainReset(*domain.ToLibvirt(d), 0); err != nil {
		return err
	}

	return nil
}

func (a *API) PoweroffDomainByUUID(u string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainDestroy(*domain.ToLibvirt(d)); err != nil {
		return err
	}

	return nil
}

func (a *API) SaveDomainByUUID(u string, to string) error {
	d, err := a.GetDomainByUUID(u)
	if err != nil {
		return err
	}

	if err := a.conn.DomainSave(*domain.ToLibvirt(d), to); err != nil {
		return err
	}

	return nil
}

func (a *API) RestoreDomainFrom(from string) error {
	if err := a.conn.DomainRestore(from); err != nil {
		return err
	}

	return nil
}
