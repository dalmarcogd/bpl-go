package cache

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/go-redis/redis/v8"
)

type (
	ServiceImpl struct {
		serviceManager services.ServiceManager
		ctx            context.Context
		client         *redis.Client
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
	c := redis.NewClient(&redis.Options{
		Addr:     s.address,
		DB:       0,
		Password: "",
	})
	_, err := c.Ping(s.ctx).Result()
	if err != nil {
		return err
	}
	s.client = c
	return nil
}

func (s ServiceImpl) Close() error {
	return s.client.Close()
}

func (s ServiceImpl) WithServiceManager(c services.ServiceManager) services.Cache {
	s.serviceManager = c
	return s
}

func (s ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}
