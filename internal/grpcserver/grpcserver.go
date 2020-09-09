package httpserver

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/services"
)

type (
	ServiceImpl struct {
		serviceManager services.ServiceManager
		ctx            context.Context
		address        string
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) WithAddress(address string) *ServiceImpl {
	s.address = address
	return s
}

func (s *ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	return nil
}

func (s *ServiceImpl) Close() error {
	return nil
}

func (s *ServiceImpl) WithServiceManager(c services.ServiceManager) services.HttpServer {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}

func (s *ServiceImpl) Run() error {
	return nil
}
