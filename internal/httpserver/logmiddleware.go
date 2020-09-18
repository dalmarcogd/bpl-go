package httpserver

import (
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/labstack/echo/v4"
)

func LogMiddleware(log services.Logger) func(h echo.HandlerFunc) echo.HandlerFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			ctx := context.Request().Context()
			log.Info(ctx, fmt.Sprintf("Request %v:%v", context.Request().Method, context.Path()))
			err := h(context)
			log.Info(ctx, fmt.Sprintf("Response %v:%v:%v", context.Request().Method, context.Path(), context.Response().Status))
			return err
		}
	}
}
