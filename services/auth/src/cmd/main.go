package main

import (
	"log/slog"

	config "github.com/AndtoSaal/simplebank/services/auth/src/internal/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/internal/logger"
)

func main() {

	cfgServer := config.MustLoadServerConfig()

	logger := log.SetUpSlogLogger(cfgServer.Env)

	logger.Info("starting aplication",
		slog.String("env", cfgServer.Env),
		slog.Int("port", cfgServer.GRPC.Port),
	)

	//TODO: инициализация сервиса

	//TODO: запуск сервиса

	//TODO: обработка сигналов

	//TODO: остановка сервиса

}
