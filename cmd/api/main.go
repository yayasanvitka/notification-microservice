// cmd/api/main.go

package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"whatsapp-microservice/cmd/api/router"
	"whatsapp-microservice/pkg/application"
	"whatsapp-microservice/pkg/exithandler"
	"whatsapp-microservice/pkg/logger"
	"whatsapp-microservice/pkg/server"
)

func main() {
	logger.Init()
	zap.S().Info("Starting Application")

	if err := godotenv.Load(); err != nil {
		zap.S().Fatal("Failed to load env vars!")
	}

	app, err := application.Start()
	if err != nil {
		zap.S().Fatal(err.Error())
	}

	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app)).
		WithErrLogger(zap.S())

	go func() {
		zap.S().Info("starting server at ", app.Cfg.GetAPIPort())

		if err := srv.Start(); err != nil {
			zap.S().Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			zap.S().Error(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			zap.S().Error(err.Error())
		}
	})
}
