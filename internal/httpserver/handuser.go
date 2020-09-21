package httpserver

import (
	"github.com/dalmarcogd/bpl-go/internal/errors"
	"github.com/dalmarcogd/bpl-go/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *ServiceImpl) handleCreateUser(c echo.Context) error {
	uReq := new(models.UserRequest)
	if err := c.Bind(uReq); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}
	user := models.User{
		Name:  uReq.Name,
		Email: uReq.Email,
	}
	err := s.ServiceManager().Handlers().CreateUser(c.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusCreated, &models.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (s *ServiceImpl) handleUpdateUser(c echo.Context) error {
	userId := c.Param("userId")
	if userId == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors.UserIdRequired.Error()).SetInternal(errors.UserIdRequired)
	}
	uReq := new(models.UserRequest)
	if err := c.Bind(uReq); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}
	user := models.User{
		Id:    userId,
		Name:  uReq.Name,
		Email: uReq.Email,
	}
	err := s.ServiceManager().Handlers().UpdateUser(c.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, &models.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (s *ServiceImpl) handleGetUserById(c echo.Context) error {
	userId := c.Param("userId")
	if userId == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors.UserIdRequired.Error()).SetInternal(errors.UserIdRequired)
	}
	user := models.User{
		Id: userId,
	}
	err := s.ServiceManager().Handlers().GetUser(c.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, &models.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (s *ServiceImpl) handleGetUsers(c echo.Context) error {
	var users []models.User
	err := s.ServiceManager().Handlers().GetUsers(c.Request().Context(), &users)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
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
}

func (s *ServiceImpl) handleDeleteUser(c echo.Context) error {
	userId := c.Param("userId")
	if userId == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors.UserIdRequired.Error()).SetInternal(errors.UserIdRequired)
	}
	user := models.User{
		Id: userId,
	}

	err := s.ServiceManager().Handlers().DeleteUser(c.Request().Context(), &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, &models.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}
