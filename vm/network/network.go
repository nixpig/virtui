package network

import (
	"encoding/xml"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
)

func New(name string) *Network {
	return &Network{
		Name: name,
		UUID: uuid.NewString(),
	}
}

func NewFromXML(b []byte) (*Network, error) {
	n := &Network{}

	if err := xml.Unmarshal(b, n); err != nil {
		return nil, err
	}

	return n, nil
}

func (n *Network) Apply(c *libvirt.Libvirt) error {
	b, err := n.ToXML()
	if err != nil {
		return err
	}

	if _, err := c.NetworkDefineXML(string(b)); err != nil {
		return err
	}

	return nil
}

func (n *Network) ToXML() ([]byte, error) {
	return xml.Marshal(n)
}

func (n *Network) ToXMLFormatted() ([]byte, error) {
	return xml.MarshalIndent(n, "", "  ")
}

type Network struct {
	XMLName     xml.Name         `xml:"network"`
	Name        string           `xml:"name"`
	Title       string           `xml:"title,omitempty"`
	UUID        string           `xml:"uuid,omitempty"`
	Description string           `xml:"description,omitempty"`
	Metadata    *NetworkMetadata `xml:"metadata,omitempty"`
	Forward     *Forward         `xml:"forward"`
	Bridge      *Bridge          `xml:"bridge"`
	IP          *IP              `xml:"ip"`
}

type NetworkMetadata struct{}

type Forward struct {
	Mode string `xml:"mode,attr,omitempty"`
	NAT  *NAT   `xml:"nat"`
}

type NAT struct {
	Address *AddressRange `xml:"address"`
	Port    *PortRange    `xml:"port"`
}

type PortRange struct {
	Start string `xml:"start,attr,omitempty"`
	End   string `xml:"end,attr,omitempty"`
}

type Bridge struct {
	Name  string `xml:"name,attr"`
	STP   string `xml:"stp,attr,omitempty"`
	Delay int    `xml:"delay,attr"`
}

type IP struct {
	Address string `xml:"address,attr,omitempty"`
	Netmask string `xml:"netmask,attr,omitempty"`
	DHCP    *DHCP  `xml:"dhcp"`
}

type DHCP struct {
	Ranges []AddressRange `xml:"range"`
}

type AddressRange struct {
	Start string `xml:"start,attr,omitempty"`
	End   string `xml:"end,attr,omitempty"`
}

func NewWithDefaults(name string) *Network {
	n := New(name)

	n.Forward = &Forward{
		Mode: "nat",
		NAT: &NAT{
			Port: &PortRange{
				Start: "1024",
				End:   "65535",
			},
		},
	}

	n.Bridge = &Bridge{
		Name:  "virbr1",
		STP:   "on",
		Delay: 0,
	}

	n.IP = &IP{
		Address: "192.168.133.1",
		Netmask: "255.255.255.0",
		DHCP: &DHCP{
			Ranges: []AddressRange{
				{Start: "192.168.133.2", End: "192.168.133.254"},
			},
		},
	}

	return n
}
