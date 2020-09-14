package httpserver

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/labstack/echo/v4"
	"io/ioutil"
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

func (s *ServiceImpl) WithAddress(address string) *ServiceImpl {
	s.address = address
	return s
}

func (s *ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	s.echo = echo.New()
	s.echo.Logger.SetOutput(ioutil.Discard)
	s.echo.Use(LogMiddleware(s.ServiceManager().Logger()))
	s.RegisterRoutes()
	return nil
}

func (s *ServiceImpl) Close() error {
	if err := s.echo.Shutdown(s.ctx); err != nil {
		return err
	}
	return s.echo.Close()
}

func (s *ServiceImpl) WithServiceManager(c services.ServiceManager) services.HttpServer {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}

func (s *ServiceImpl) RegisterRoutes() *ServiceImpl {
	group := s.echo.Group("/v1")
	group.POST("/users", s.routeCreateUser)
	group.PATCH("/users/:userId", s.routeUpdateUser)
	group.GET("/users/:userId", s.routeGetUserById)
	group.GET("/users", s.routeGetUsers)
	group.DELETE("/users/:userId", s.routeDeleteUser)
	return s
}

func (s *ServiceImpl) Run() error {
	return s.echo.Start(s.address)
}
