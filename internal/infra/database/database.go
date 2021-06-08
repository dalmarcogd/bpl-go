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
		serviceManager services.Sis
		ctx            context.Context
		client         *gorm.DB
		dsn            string
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) WithDsn(dsn string) *ServiceImpl {
	s.dsn = dsn
	return s
}

func (s *ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	s.dsn = s.Sis().Environment().DatabaseDsn()
	c, err := gorm.Open(postgres.Open(s.dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}
	s.client = c
	db, err := s.client.DB()
	if err != nil {
		return err
	}
	if err := db.PingContext(s.ctx); err != nil {
		return err
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return nil
}

func (s *ServiceImpl) Close() error {
	db, err := s.client.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (s *ServiceImpl) WithSis(c services.Sis) services.Database {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) Sis() services.Sis {
	return s.serviceManager
}

func (s *ServiceImpl) DB(ctx context.Context) *gorm.DB {
	return s.client.WithContext(ctx)
}
