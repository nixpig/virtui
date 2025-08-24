package tui

import (
	"context"

	"github.com/nixpig/virtui/internal/entity"
	
	"github.com/stretchr/testify/mock"
	"libvirt.org/go/libvirt"
)

// MockService is a mock type for the Service interface
type MockService struct {
	mock.Mock
}

func (m *MockService) GetConnectionDetails() (entity.ConnectionDetails, error) {
	args := m.Called()
	return args.Get(0).(entity.ConnectionDetails), args.Error(1)
}

func (m *MockService) RegisterDomainEventCallback(events chan *libvirt.DomainEventLifecycle) (int, error) {
	args := m.Called(events)
	return args.Int(0), args.Error(1)
}

func (m *MockService) DeregisterDomainEventCallback(callbackID int) error {
	args := m.Called(callbackID)
	return args.Error(0)
}

func (m *MockService) EventLoop(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockService) ListAllDomains() ([]entity.DomainWithState, error) {
	args := m.Called()
	return args.Get(0).([]entity.DomainWithState), args.Error(1)
}

func (m *MockService) LookupDomainByUUIDString(uuid string) (entity.Domain, error) {
	args := m.Called(uuid)
	return args.Get(0).(entity.Domain), args.Error(1)
}

func (m *MockService) StartDomain(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}

func (m *MockService) PauseResumeDomain(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}

func (m *MockService) ShutdownDomain(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}

func (m *MockService) RebootDomain(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}

func (m *MockService) ForceResetDomain(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}

func (m *MockService) ForceOffDomain(uuid string) error {
	args := m.Called(uuid)
	return args.Error(0)
}

func (m *MockService) ListAllNetworks() ([]entity.Network, error) {
	args := m.Called()
	return args.Get(0).([]entity.Network), args.Error(1)
}

func (m *MockService) ListAllStoragePools() ([]entity.StoragePool, error) {
	args := m.Called()
	return args.Get(0).([]entity.StoragePool), args.Error(1)
}

func (m *MockService) Close() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
