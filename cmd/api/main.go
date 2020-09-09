package main

import (
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/cache"
	"github.com/dalmarcogd/bpl-go/internal/database"
	"github.com/dalmarcogd/bpl-go/internal/httpserver"
	"github.com/dalmarcogd/bpl-go/internal/logger"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ss := services.
		New().
		WithDatabase(database.New().WithDsn("user=postgres password=postgres dbname=blp host=localhost port=3306 sslmode=disable TimeZone=UTC")).
		WithCache(cache.New().WithAddress("localhost:6379")).
		WithLogger(logger.New()).
		WithHttpServer(httpserver.New().WithAddress(":8080"))

	if err := ss.Init(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}

	go func() {
		if err := ss.HttpServer().Run(); err != nil {
			ss.Logger().Fatal(ss.Context(), err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-quit

	ss.Logger().Info(ss.Context(), fmt.Sprintf("Shutdown by %v", sig.String()))

	if err := ss.Close(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}
	ss.Logger().Info(ss.Context(), "All services closed")
}
