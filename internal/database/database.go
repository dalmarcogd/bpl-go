package database

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type (
	ServiceImpl struct {
		serviceManager services.ServiceManager
		ctx            context.Context
		client         *gorm.DB
		dsn            string
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s ServiceImpl) WithDsn(dsn string) ServiceImpl {
	s.dsn = dsn
	return s
}

func (s ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	c, err := gorm.Open(postgres.Open(s.dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	s.client = c
	db, err := s.client.DB()
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	s.ServiceManager().Logger()
	return nil
}

func (s ServiceImpl) Close() error {
	db, err := s.client.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (s ServiceImpl) WithServiceManager(c services.ServiceManager) services.Database {
	s.serviceManager = c
	return s
}

func (s ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}
