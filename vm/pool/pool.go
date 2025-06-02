package pool

import (
	"encoding/xml"

	"github.com/digitalocean/go-libvirt"
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
