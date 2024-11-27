package main

import (
	"fmt"

	"github.com/AndtoSaal/simplebank/services/auth/src/internal/config"
)

func main() {
	cfgServer := config.MustLoadServerConfig()
	fmt.Println(cfgServer)

	//TODO: логгер

	//TODO: инициализация сервиса

	//TODO: запуск сервиса

	//TODO: обработка сигналов

	//TODO: остановка сервиса

}
