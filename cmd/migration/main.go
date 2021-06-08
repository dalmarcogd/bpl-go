package main

import (
	database2 "github.com/dalmarcogd/bpl-go/internal/infra/database"
	environment2 "github.com/dalmarcogd/bpl-go/internal/infra/environment"
	logger2 "github.com/dalmarcogd/bpl-go/internal/infra/logger"
	"github.com/dalmarcogd/bpl-go/internal/models"
	"github.com/dalmarcogd/bpl-go/internal/services"
)

func main() {
	ss := services.
		New().
		WithDatabase(database2.New()).
		WithLogger(logger2.New()).
		WithEnvironment(environment2.New())

	if err := ss.Init(); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}

	if err := ss.Database().DB(ss.Context()).AutoMigrate(&models.User{}); err != nil {
		ss.Logger().Fatal(ss.Context(), err.Error())
		return
	}
	ss.Logger().Info(ss.Context(), "Migration finished")
}
