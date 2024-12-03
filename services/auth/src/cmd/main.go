package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	config "github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
	auth_transport "github.com/AndtoSaal/simplebank/services/auth/src/transport/grpc"
)

func main() {

	///инициализация конфигов сервера и дб
	cfgService := config.MustLoadConfig()

	fmt.Println(cfgService)
	fmt.Println("OK configs")

	//инициализируем логгер с соотв переменной окружения
	logger := log.SetUpSlogLogger(cfgService.Srv.Env)

	fmt.Println("OK logger")

	//инициализируем все слои сервиса
	transportLevel := auth_transport.NewAuthTransport(logger, *cfgService)

	fmt.Println("OK transort")
	//запуск всего сервиса, при ошибки - паника
	go func() {
		transportLevel.MustRun()
	}()

	fmt.Println("OK MUSTRUN")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Waiting for SIGINT (pkill -2) or SIGTERM
	<-stop

	//graceful shutdown, закрытие коннекта к базе, остановка сервра
	transportLevel.Stop()

	//TODO: остановка сервиса (graicfull shoutdown)

}
