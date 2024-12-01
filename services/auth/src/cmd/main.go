package main

import (
	"time"

	config "github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
	auth_repository "github.com/AndtoSaal/simplebank/services/auth/src/repository"
	"github.com/AndtoSaal/simplebank/services/auth/src/service/auth_service"
	"github.com/AndtoSaal/simplebank/services/auth/src/service/usrInfo_service"
	auth_transport "github.com/AndtoSaal/simplebank/services/auth/src/transport/grpc"
)

const (
	tokenTTL time.Duration = time.Hour * 1
)

func main() {

	///инициализация конфигов сервера и дб
	cfgServer, cfgDataBase := config.MustLoadConfig()

	//инициализируем логгер с соотв переменной окружения
	logger := log.SetUpSlogLogger(cfgServer.Env)

	db, err := auth_repository.NewPostgresDB(cfgDataBase)
	if err != nil {
		logger.Error("Cannot connect to databse", (err).Error())
	}

	repositoryLevel := auth_repository.NewAuthPostgresRepo(db)
	authServiceLevel := auth_service.NewAuthService(logger, repositoryLevel, tokenTTL)
	usrInfoServiceLevel := usrInfo_service.NewUserInfoService(logger, repositoryLevel)
	transportLevel := auth_transport.NewAuthTransport()

	//TODO: инициализировать слой репозитория

	//TODO: инициализировать слой сервиса

	//TODO: инициализировать слой хэндлеров

	//TODO: инициализация сервера

	//TODO: Запуск сервера

	//TODO: обработка сигналов

	//TODO: остановка сервиса (graicfull shoutdown)

}
