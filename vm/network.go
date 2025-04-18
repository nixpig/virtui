package vm

type Network struct{}

func NewNetwork() (*Network, error) {
	return &Network{}, nil
}
