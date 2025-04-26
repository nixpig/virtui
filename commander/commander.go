package commander

import (
	"github.com/digitalocean/go-libvirt"
)

type Commander struct {
	conn *libvirt.Libvirt
}

type connection interface {
	DomainCreate()
}

func NewCommander(conn *libvirt.Libvirt) *Commander {
	return &Commander{conn}
}

func (c *Commander) StartDomain(d *libvirt.Domain) error {
	if err := c.conn.DomainCreate(*d); err != nil {
		return err
	}

	return nil
}

func (c *Commander) PauseDomain(d *libvirt.Domain) error {
	return nil
}

func (c *Commander) ResumeDomain(d *libvirt.Domain) error {
	if err := c.conn.DomainResume(*d); err != nil {
		return err
	}

	return nil
}

func (c *Commander) ShutdownDomain(d *libvirt.Domain) error {
	if err := c.conn.DomainShutdown(*d); err != nil {
		return err
	}

	return nil
}

func (c *Commander) RebootDomain(d *libvirt.Domain) error {
	if err := c.conn.DomainReboot(*d, 0); err != nil {
		return err
	}

	return nil
}

func (c *Commander) ResetDomain(d *libvirt.Domain) error {
	if err := c.conn.DomainReset(*d, 0); err != nil {
		return err
	}

	return nil
}

func (c *Commander) PoweroffDomain(d *libvirt.Domain) error {
	if err := c.conn.DomainDestroy(*d); err != nil {
		return err
	}

	return nil
}

func (c *Commander) SaveDomain(d *libvirt.Domain, to string) error {
	if err := c.conn.DomainSave(*d, to); err != nil {
		return err
	}

	return nil
}

func (c *Commander) RestoreDomain(from string) error {
	if err := c.conn.DomainRestore(from); err != nil {
		return err
	}

	return nil
}

func (c *Commander) DeleteDomain(d *libvirt.Domain) error {
	if err := c.conn.DomainUndefine(*d); err != nil {
		return err
	}

	return nil
}
