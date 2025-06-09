package pool

import (
	"encoding/xml"

	"github.com/charmbracelet/log"
	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
	"github.com/nixpig/virtui/vm/volume"
)

type Pool struct {
	XMLName    xml.Name    `xml:"pool"`
	Type       string      `xml:"type,attr"`
	Name       string      `xml:"name"`
	UUID       string      `xml:"uuid"`
	Source     Source      `xml:"source"`
	Capacity   *Capacity   `xml:"capacity,omitempty"`
	Allocation *Allocation `xml:"allocation,omitempty"`
	Available  *Available  `xml:"available,omitempty"`
	Target     *Target     `xml:"target,omitempty"`
}

type Source struct {
	Value string `xml:",innerxml"`
}

type Target struct {
	Path        string       `xml:"path"`
	Permissions *Permissions `xml:"permissions,omitempty"`
}

type Permissions struct {
	Mode  int `xml:"mode"`
	Owner int `xml:"owner"`
	Group int `xml:"group"`
}

type Capacity struct {
	CharData string `xml:",chardata"`
	Unit     string `xml:"unit,attr"`
}

type Allocation struct {
	CharData string `xml:",chardata"`
	Unit     string `xml:"unit,attr"`
}

type Available struct {
	CharData string `xml:",chardata"`
	Unit     string `xml:"unit,attr"`
}

func New(name string) *Pool {
	return &Pool{Name: name}
}

func NewFromXML(b []byte) (*Pool, error) {
	p := &Pool{}

	if err := xml.Unmarshal(b, p); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Pool) Save(c *libvirt.Libvirt) error {
	// TODO: implement saving of changes
	return nil
}

func (p *Pool) ToXML() ([]byte, error) {
	return xml.Marshal(p)
}

func (p *Pool) ToXMLFormatted() ([]byte, error) {
	return xml.MarshalIndent(p, "", "  ")
}

func (p *Pool) GetVolumes(c *libvirt.Libvirt) ([]volume.Volume, error) {
	u, err := uuid.Parse(p.UUID)
	if err != nil {
		return nil, err
	}

	volumes, err := c.StoragePoolListVolumes(
		libvirt.StoragePool{
			Name: p.Name,
			UUID: libvirt.UUID(u),
		},
		1024,
	)
	if err != nil {
		return nil, err
	}

	volumeList := make([]volume.Volume, len(volumes))

	for i, v := range volumes {
		n, err := c.StorageVolLookupByName(
			libvirt.StoragePool{
				Name: p.Name,
				UUID: libvirt.UUID(u),
			},
			v,
		)
		if err != nil {
			log.Error(err)
			continue
		}

		s, err := c.StorageVolGetXMLDesc(n, 0)
		if err != nil {
			log.Error(err)
			continue
		}

		x, err := volume.NewFromXML([]byte(s))
		if err != nil {
			return nil, err
		}

		volumeList[i] = *x
	}

	return volumeList, nil
}

func List(c *libvirt.Libvirt) ([]Pool, error) {
	pools, _, err := c.ConnectListAllStoragePools(1, 0)
	if err != nil {
		return nil, err
	}

	poolList := make([]Pool, len(pools))

	for i, p := range pools {
		s, err := c.StoragePoolGetXMLDesc(p, 0)
		if err != nil {
			log.Error(err)
			continue
		}

		x, err := NewFromXML([]byte(s))
		if err != nil {
			log.Error(err)
			continue
		}

		poolList[i] = *x
	}

	return poolList, nil
}
