package main

import (
	config "github.com/AndtoSaal/simplebank/services/auth/src/internal/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/internal/logger"
)

func main() {

	cfgServer := config.MustLoadServerConfig()

	logger := log.SetUpSlogLogger(cfgServer.Env)

	//TODO: запуск сервиса

	//TODO: обработка сигналов

	//TODO: остановка сервиса

}