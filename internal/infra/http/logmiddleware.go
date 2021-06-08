package http

import (
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LogMiddleware(log services.Logger) func(h echo.HandlerFunc) echo.HandlerFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			ctx := context.Request().Context()
			log.Info(ctx, fmt.Sprintf("Request %v:%v", context.Request().Method, context.Path()))
			err := h(context)
			status := context.Response().Status
			if err != nil {
				status = http.StatusInternalServerError
				he, ok := err.(*echo.HTTPError)
				if ok {
					if he.Internal != nil {
						if herr, ok := he.Internal.(*echo.HTTPError); ok {
							he = herr
						}
					}
					status = he.Code
				}
			}
			log.Info(ctx, fmt.Sprintf("Response %v:%v:%v", context.Request().Method, context.Path(), status))

			return err
		}
	}
}
