package main

import (
	"github.com/MikeThesol/BKSH/internal/config"
	"github.com/MikeThesol/BKSH/internal/repository"
	"github.com/MikeThesol/BKSH/internal/service"
)

func main() {
	cfg := config.MustLoad()
	log := config.InitLogger("local")
	log.Info("server start")

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Error("DB not create", err.Error())
		return
	}
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	_ = service
}
