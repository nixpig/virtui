package vm

type Volume struct{}

func NewStorage() (*Volume, error) {
	return &Volume{}, nil
}
