package vm

type Network struct{}

func NewVNet() (*Network, error) {
	return &Network{}, nil
}
