package main

import (
	"fmt"

	config "github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
)

func main() {

	///инициализация конфигов сервера и дб
	cfgServer, cfgDataBase := config.MustLoadConfig()

	//инициализируем логгер с соотв переменной окружения
	logger := log.SetUpSlogLogger(cfgServer.Env)
	logger.Debug("проверка логгера")
	fmt.Println(cfgDataBase, cfgServer)

	//TODO: инициализировать слой репозитория

	//TODO: инициализировать слой сервиса

	//TODO: инициализировать слой хэндлеров

	//TODO: инициализация сервера

	//TODO: Запуск сервера

	//TODO: обработка сигналов

	//TODO: остановка сервиса (graicfull shoutdown)

}
