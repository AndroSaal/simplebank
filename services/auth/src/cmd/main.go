package main

import (
	"time"

	config "github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
	auth_transport "github.com/AndtoSaal/simplebank/services/auth/src/transport/grpc"
)

const (
	tokenTTL time.Duration = time.Hour * 1
)

func main() {

	///инициализация конфигов сервера и дб
	cfgService := config.MustLoadConfig()

	//инициализируем логгер с соотв переменной окружения
	logger := log.SetUpSlogLogger(cfgService.Srv.Env)

	//инициализируем все слои сервиса
	transportLevel := auth_transport.NewAuthTransport(logger, *cfgService)

	//посмотреть как завести эту историю !
	transportLevel.gRPCServer.Run()
	//TODO: инициализировать слой репозитория

	//TODO: инициализировать слой сервиса

	//TODO: инициализировать слой хэндлеров

	//TODO: инициализация сервера

	//TODO: Запуск сервера

	//TODO: обработка сигналов

	//TODO: остановка сервиса (graicfull shoutdown)

}
