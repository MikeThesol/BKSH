package main

import "github.com/MikeThesol/BKSH/internal/config"

func main() {
	log := config.InitLogger("local")
	log.Info("server start")
}
