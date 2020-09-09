package handlers

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/models"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/google/uuid"
)

type (
	ServiceImpl struct {
		serviceManager services.ServiceManager
		ctx            context.Context
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	return nil
}

func (s *ServiceImpl) Close() error {
	return nil
}

func (s *ServiceImpl) WithServiceManager(c services.ServiceManager) services.Handlers {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}

func (s *ServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	user.Id = newUUID.String()
	result := s.ServiceManager().Database().DB(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) UpdateUser(ctx context.Context, u *models.User) error {
	result := s.ServiceManager().Database().DB(ctx).Save(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) GetUser(ctx context.Context, u *models.User) error {
	result := s.ServiceManager().Database().DB(ctx).Find(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) GetUsers(ctx context.Context, u *[]models.User) error {
	result := s.ServiceManager().Database().DB(ctx).Find(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) DeleteUser(ctx context.Context, u *models.User) error {
	result := s.ServiceManager().Database().DB(ctx).Delete(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
