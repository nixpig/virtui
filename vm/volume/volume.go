package volume

import "encoding/xml"

func New(name string) *Volume {
	return &Volume{Name: name}
}

func NewFromXML(b []byte) (*Volume, error) {
	v := &Volume{}

	if err := xml.Unmarshal(b, v); err != nil {
		return nil, err
	}

	return v, nil
}

func NewWithDefaults(name string) *Volume {
	v := New(name)

	v.Type = "file"
	v.Allocation = 0
	v.Capacity = &Capacity{Unit: "G", CharData: 10}
	v.Target = &VolumeTarget{Format: &Format{Type: "qcow2"}}

	return v
}

func (v *Volume) ToXML() ([]byte, error) {
	return xml.Marshal(v)
}

func (v *Volume) ToXMLFormatted() ([]byte, error) {
	return xml.MarshalIndent(v, "", "  ")
}

type Volume struct {
	XMLName    xml.Name      `xml:"volume"`
	Type       string        `xml:"type,attr"`
	Name       string        `xml:"name"`
	Allocation int           `xml:"allocation"`
	Capacity   *Capacity     `xml:"capacity"`
	Target     *VolumeTarget `xml:"target"`
}

type Capacity struct {
	Unit     string `xml:"unit,attr,omitempty"`
	CharData int    `xml:",chardata"`
}

type VolumeTarget struct {
	Format *Format `xml:"format"`
}

type Format struct {
	Type string `xml:"type,attr"`
}
