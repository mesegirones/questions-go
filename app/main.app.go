package main

import (
	"context"
	"fmt"
	"questions-go/health"
	"questions-go/internal/config"
	"questions-go/internal/proxy/logger"
	"questions-go/internal/rest"

	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	ctx := context.Background()

	// Setting up dependencies
	config := config.NewConfig()
	loggerProxy := logger.NewLoggerProxy(ctx, config.GetRestConfig())

	r := rest.NewGinEngine(config.GetRestConfig(), loggerProxy)

	healthService := health.NewService(config.GetHealthConfig(), loggerProxy)
	rest.NewHealthHandler(r, healthService)

	if err := r.Run(fmt.Sprintf(":%s", config.GetRestConfig().GetPort())); err != nil {
		loggerProxy.Error(err)
	}
}
