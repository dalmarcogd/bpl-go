package handlers

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/models"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/google/uuid"
)

type (
	ServiceImpl struct {
		serviceManager services.Sis
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

func (s *ServiceImpl) WithSis(c services.Sis) services.Handlers {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) Sis() services.Sis {
	return s.serviceManager
}

func (s *ServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	user.Id = uuid.New().String()
	result := s.Sis().Database().DB(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) UpdateUser(ctx context.Context, u *models.User) error {
	result := s.Sis().Database().DB(ctx).Save(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) GetUser(ctx context.Context, u *models.User) error {
	result := s.Sis().Database().DB(ctx).Find(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) GetUsers(ctx context.Context, u *[]models.User) error {
	result := s.Sis().Database().DB(ctx).Find(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *ServiceImpl) DeleteUser(ctx context.Context, u *models.User) error {
	result := s.Sis().Database().DB(ctx).Delete(&u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
