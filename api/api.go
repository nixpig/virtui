package api

import (
	"github.com/charmbracelet/log"
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

func NewAPI(conn *libvirt.Libvirt) *API {
	return &API{conn}
}

func (a *API) CreateNetwork(n *network.Network) (*network.Network, error) {
	return &network.Network{}, nil
}

func (a *API) GetNetworks() []network.Network {
	// n, _, err := c.ConnectListAllNetworks(1, 0)
	// if err != nil {
	// 	log.Error("failed to list networks", "err", err)
	// 	continue
	// }
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
	// v, err := c.StoragePoolListVolumes(p, 1024)
	// if err != nil {
	// 	log.Error("failed to list volumes", "err", err)
	// 	continue
	// }
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
	// domains, _, err := c.ConnectListAllDomains(1, 0)
	// if err != nil {
	// 	continue
	// }
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

func (a *API) GetPools() ([]*pool.Pool, error) {
	pools, _, err := a.conn.ConnectListAllStoragePools(1, 0)
	if err != nil {
		return nil, err
	}

	r := []*pool.Pool{}

	for _, p := range pools {
		x, err := a.conn.StoragePoolGetXMLDesc(p, 0)
		if err != nil {
			log.Error(err)
			continue
		}

		d, err := pool.NewFromXML([]byte(x))
		if err != nil {
			log.Error(err)
			continue
		}

		r = append(r, d)

	}

	return r, nil
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

// for _, d := range domains {
// 	cpuStats, _, err := c.DomainGetCPUStats(d, 1, -1, 1, 0)
// 	if err != nil {
// 		log.Error("failed to get cpu stats", "domain", d.Name, "err", err)
// 	} else {
// 		for i, s := range cpuStats {
// 			log.Info("cpustats", "iter", i, "domain", d.Name, "stat", s)
// 		}
// 	}
//
// 	memStats, err := c.DomainMemoryStats(d, uint32(libvirt.DomainMemoryStatNr), 0)
// 	if err != nil {
// 		log.Error("failed to get mem stats", "domain", d.Name, "err", err)
// 	} else {
// 		for _, s := range memStats {
// 			tag := libvirt.DomainMemoryStatTags(s.Tag)
//
// 			// a'la https://github.com/free4inno/prometheus-libvirt-exporter/blob/9da210267ae14300fdd4d2036294e66bbecaa03b/collector/memory.go#L183
// 			switch tag {
// 			case libvirt.DomainMemoryStatAvailable:
// 				log.Info("available memory", "domain", d.Name, "mem", s.Val)
// 			}
//
// 		}
// 	}
//
// 	state, _, _ := c.DomainGetState(d, 0)
// 	uuid, _ := uuid.FromBytes(d.UUID[:])
//
// 	data = append(data, dashboardDomain{
// 		lvdom:  &d,
// 		conn:   c,
// 		name:   d.Name,
// 		uuid:   uuid.String(),
// 		id:     int(d.ID),
// 		status: domain.PresentableState(libvirt.DomainState(state)),
// 		host:   u.Host + u.Path,
// 	})

// ---

// if _, err := model.store.GetConnectionByURI(string(libvirt.QEMUSystem)); err != nil {
// 	log.Debug("system connection not found; insert new")
// 	if err := model.store.InsertConnection(&connection.Connection{
// 		URI: string(libvirt.QEMUSystem),
// 	}); err != nil {
// 		log.Error("failed to insert system connection")
// 	}
// }

// lv, err := libvirt.ConnectToURI(uri)
// if err != nil {
// 	log.Error("failed to connect to uri", "uri", uri, "err", err)
// 	continue
// }
