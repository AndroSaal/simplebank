package main

import (
	config "github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
)

func main() {

	cfgServer := config.MustLoadServerConfig()

	logger := log.SetUpSlogLogger(cfgServer.Env)

	//TODO: инициализировать слой репозитория

	//TODO: инициализировать слой сервиса

	//TODO: инициализировать слой хэндлеров

	//TODO: инициализация сервера

	//TODO: Запуск сервера

	//TODO: обработка сигналов

	//TODO: остановка сервиса (graicfull shoutdown)

}
