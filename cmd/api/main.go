package main

import (
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/handlers"
	cache2 "github.com/dalmarcogd/bpl-go/internal/infra/cache"
	database2 "github.com/dalmarcogd/bpl-go/internal/infra/database"
	environment2 "github.com/dalmarcogd/bpl-go/internal/infra/environment"
	logger2 "github.com/dalmarcogd/bpl-go/internal/infra/logger"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"os"
	"os/signal"
)

func main() {
	ss := services.
		New().
		WithDatabase(database2.New()).
		WithCache(cache2.New()).
		WithLogger(logger2.New()).
		WithHttpServer(http.New().WithAddress(":8080")).
		WithHandlers(handlers.New()).
		WithEnvironment(environment2.New())

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
