package vm

type Volume struct{}

func NewVolume() (*Volume, error) {
	return &Volume{}, nil
}
