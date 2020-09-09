package main

import (
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/cache"
	"github.com/dalmarcogd/bpl-go/internal/database"
	"github.com/dalmarcogd/bpl-go/internal/environment"
	"github.com/dalmarcogd/bpl-go/internal/handlers"
	"github.com/dalmarcogd/bpl-go/internal/httpserver"
	"github.com/dalmarcogd/bpl-go/internal/logger"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"os"
	"os/signal"
)

func main() {
	ss := services.
		New().
		WithDatabase(database.New()).
		WithCache(cache.New()).
		WithLogger(logger.New()).
		WithHttpServer(httpserver.New().WithAddress(":8080")).
		WithHandlers(handlers.New()).
		WithEnvironment(environment.New())

	if err := ss.Init(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}

	go func() {
		ss.Logger().Info(ss.Context(), "Http server started")
		if err := ss.HttpServer().Run(); err != nil {
			ss.Logger().Fatal(ss.Context(), err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	sig := <-quit

	ss.Logger().Info(ss.Context(), fmt.Sprintf("Shutdown by %v", sig.String()))

	if err := ss.Close(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}
	ss.Logger().Info(ss.Context(), "All services closed")
}
