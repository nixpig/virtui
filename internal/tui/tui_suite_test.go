package tui

import (
	"context"
	"testing"

	"github.com/nixpig/virtui/internal/entity"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TUITestSuite struct {
	suite.Suite
	mockService *MockService
}

func (s *TUITestSuite) SetupTest() {
	s.mockService = new(MockService)
}

func (s *TUITestSuite) TestEventHandling() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	callbackID := 123

	s.mockService.On("GetConnectionDetails").Return(entity.ConnectionDetails{}, nil)
	s.mockService.On("ListAllDomains").Return([]entity.DomainWithState{}, nil)
	s.mockService.On("ListAllNetworks").Return([]entity.Network{}, nil)
	s.mockService.On("ListAllStoragePools").Return([]entity.StoragePool{}, nil)
	s.mockService.On("RegisterDomainEventCallback", mock.Anything).Return(callbackID, nil)
	s.mockService.On("EventLoop", ctx).Return(nil)

	model, err := New(s.mockService, ctx)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), model)

	s.mockService.AssertCalled(s.T(), "GetConnectionDetails")
	s.mockService.AssertCalled(s.T(), "RegisterDomainEventCallback", mock.Anything)
	s.mockService.AssertCalled(s.T(), "EventLoop", ctx)
}

func TestTuiTestSuite(t *testing.T) {
	suite.Run(t, new(TUITestSuite))
}