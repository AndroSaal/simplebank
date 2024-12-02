package main

import (
	config "github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
	auth_transport "github.com/AndtoSaal/simplebank/services/auth/src/transport/grpc"
)

func main() {

	///инициализация конфигов сервера и дб
	cfgService := config.MustLoadConfig()

	//инициализируем логгер с соотв переменной окружения
	logger := log.SetUpSlogLogger(cfgService.Srv.Env)

	//инициализируем все слои сервиса
	transportLevel := auth_transport.NewAuthTransport(logger, *cfgService)

	transportLevel.Run()
	//посмотреть как завести эту историю !]
	//TODO: инициализировать слой репозитория

	//TODO: инициализировать слой сервиса

	//TODO: инициализировать слой хэндлеров

	//TODO: инициализация сервера

	//TODO: Запуск сервера

	//TODO: обработка сигналов

	//TODO: остановка сервиса (graicfull shoutdown)

}
