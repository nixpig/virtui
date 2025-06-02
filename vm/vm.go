package vm

import "github.com/digitalocean/go-libvirt"

type Newable[T any] interface {
	New(name string) *T
	NewWithDefaults(name string) *T
	NewFromXML(b []byte) (*T, error)
}

type Savable interface {
	Save(c *libvirt.Libvirt) error
}

type XMLable interface {
	ToXML() ([]byte, error)
	ToXMLFormatted([]byte, error)
}
