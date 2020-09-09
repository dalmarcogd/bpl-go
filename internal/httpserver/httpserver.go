package httpserver

import (
	"context"
	"github.com/dalmarcogd/bpl-go/internal/errors"
	"github.com/dalmarcogd/bpl-go/internal/models"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
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

	group.POST("/users", func(c echo.Context) error {
		uReq := new(models.UserRequest)
		if err := c.Bind(&uReq); err != nil {
			return err
		}
		user := models.User{
			Name:  uReq.Name,
			Email: uReq.Email,
		}
		err := s.ServiceManager().Handlers().CreateUser(c.Request().Context(), &user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, &models.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	})
	group.PATCH("/users/:userId", func(c echo.Context) error {
		uReq := new(models.UserRequest)
		userId := c.Param("userId")
		if userId != "" {
			return errors.UserIdRequired
		}
		if err := c.Bind(&uReq); err != nil {
			return err
		}
		user := models.User{
			Id:    userId,
			Name:  uReq.Name,
			Email: uReq.Email,
		}
		err := s.ServiceManager().Handlers().CreateUser(c.Request().Context(), &user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &models.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	})
	group.GET("/users/:userId", func(c echo.Context) error {
		userId := c.Param("userId")
		if userId != "" {
			return errors.UserIdRequired
		}
		user := models.User{
			Id: userId,
		}

		err := s.ServiceManager().Handlers().GetUser(c.Request().Context(), &user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &models.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	})
	group.GET("/users", func(c echo.Context) error {
		var users []models.User
		err := s.ServiceManager().Handlers().GetUsers(c.Request().Context(), &users)
		if err != nil {
			return err
		}

		uResponses := make([]*models.UserResponse, 0)
		for _, user := range users {
			uResponses = append(uResponses, &models.UserResponse{
				Id:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			})
		}
		return c.JSON(http.StatusOK, &uResponses)
	})
	group.DELETE("/users/:userId", func(c echo.Context) error {
		userId := c.Param("userId")
		if userId != "" {
			return errors.UserIdRequired
		}
		user := models.User{
			Id: userId,
		}

		err := s.ServiceManager().Handlers().DeleteUser(c.Request().Context(), &user)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &models.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	})
	return s
}

func (s *ServiceImpl) Run() error {
	return s.echo.Start(s.address)
}
