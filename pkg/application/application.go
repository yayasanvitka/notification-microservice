package application

import (
	"notification-microservice/pkg/config"
)

type Application struct {
	Cfg *config.Config
}

func Start() (*Application, error) {
	cfg := config.Get()

	return &Application{
		Cfg: cfg,
	}, nil
}
