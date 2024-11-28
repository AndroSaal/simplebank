package log

import (
	"log/slog"
	"os"
)

const (
	local = "local"
	prod  = "prod"
	dev   = "dev"
)

func SetUpSlogLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case local:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case dev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prod:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}
