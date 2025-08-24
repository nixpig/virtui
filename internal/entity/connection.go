package entity

// ConnectionDetails contains information about the hypervisor connection.
type ConnectionDetails struct {
	Hostname  string
	URI       string
	ConnType  string
	HvVersion string
	LvVersion string
}
