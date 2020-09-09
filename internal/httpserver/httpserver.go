package httpserver

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/labstack/echo/v4"
)

type (
	ServiceImpl struct {
		serviceManager services.ServiceManager
		ctx            context.Context
		echo           *echo.Echo
		address        string
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s ServiceImpl) WithAddress(address string) ServiceImpl {
	s.address = address
	return s
}

func (s ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	s.echo = echo.New()
	return nil
}

func (s ServiceImpl) Close() error {
	return nil
}

func (s ServiceImpl) WithServiceManager(c services.ServiceManager) services.HttpServer {
	s.serviceManager = c
	return s
}

func (s ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}

func (s ServiceImpl) RegisterRoutes(f func(e *echo.Echo)) {
	f(s.echo)
}

func (s ServiceImpl) Run() error {
	return s.echo.Start(s.address)
}
