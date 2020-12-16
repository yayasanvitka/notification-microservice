package application

import (
	"whatsapp-microservice/pkg/config"
	"whatsapp-microservice/pkg/db"
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
