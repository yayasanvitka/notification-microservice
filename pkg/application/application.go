package application

import (
	"notification-microservice/pkg/config"
	"notification-microservice/pkg/db"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Start() (*Application, error) {
	cfg := config.Get()
	database, err := db.Connect(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  database,
		Cfg: cfg,
	}, nil
}
