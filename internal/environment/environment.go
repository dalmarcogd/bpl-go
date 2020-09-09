package environment

import (
	"context"

	"github.com/crgimenes/goconfig"
	"github.com/dalmarcogd/bpl-go/internal/services"
)

//Environment this object keep the all variables environment
type (
	environment struct {
		Environment  string `cfg:"ENVIRONMENT" cfgDefault:"UNKNOWN" cfgRequired:"true"`
		Service      string `cfg:"SERVICE" cfgDefault:"hsm-api" cfgRequired:"true"`
		Version      string `cfg:"VERSION" cfgDefault:"UNKNOWN" cfgRequired:"true"`
		DebugPprof   bool   `cfg:"DEBUG_PPROF" cfgDefault:"false" `
		DatabaseDsn  string `cfg:"DATABASE_DSN" cfgDefault:"user=postgres password=postgres dbname=bpl host=localhost port=5432 sslmode=disable TimeZone=UTC" cfgDefault:"false" `
		CacheAddress string `cfg:"CacheAddress" cfgDefault:"localhost:6379" cfgDefault:"false" `
	}

	ServiceImpl struct {
		serviceManager services.ServiceManager
		ctx            context.Context
		environment    *environment
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	s.environment = &environment{}
	if err := goconfig.Parse(s.environment); err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) Close() error {
	return nil
}

func (s *ServiceImpl) WithServiceManager(c services.ServiceManager) services.Environment {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}

func (s *ServiceImpl) Environment() string {
	return s.environment.Environment
}

func (s *ServiceImpl) Service() string {
	return s.environment.Service
}

func (s *ServiceImpl) Version() string {
	return s.environment.Version
}

func (s *ServiceImpl) DebugPprof() bool {
	return s.environment.DebugPprof
}

func (s *ServiceImpl) DatabaseDsn() string {
	return s.environment.DatabaseDsn
}

func (s *ServiceImpl) CacheAddress() string {
	return s.environment.CacheAddress
}
