package config

import (
	"log/slog"
	"os"
)

const (
	LOCAL_ENV = "local"
	PROD_ENV  = "prod"
)

func InitLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case LOCAL_ENV:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case PROD_ENV:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
