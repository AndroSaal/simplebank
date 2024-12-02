//конфигурация сервиса авторизации
//с помощью библиотеки github.com/ilyakaznacheev/cleanenv@latest

package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	// "google.golang.org/grpc"
)

type ServiceConfig struct {
	DB    DatabaseConfig
	Srv   ServerConfig
	Login LoginCongif
}

// конфигурация сервера
type ServerConfig struct {
	Env  string      `yaml:"env" env-default:"local"`
	GRPC GRPCConfing `yml:"grpc"`
}

// конфигурация grpc сервера
type GRPCConfing struct {
	Port    string        `yaml:"grpc:port" env-default:"50051"`
	Host    string        `yaml:"grpc:host"`
	Timeout time.Duration `yaml:"grpc:timeout"`
}

type LoginCongif struct {
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"15h"`
}

// конфигурация базы
type DatabaseConfig struct {
	Env      string `yaml:"db:env"`
	Host     string `yaml:"db:host" env-default:"localhost"`
	Port     string `yaml:"db:port" env-default:"5432"`
	UserName string `yaml:"db:username" env-default:"postgres"`
	Password string `yaml:"db:password" env-default:"qwerty"`
	Database string `yaml:"db:dbname" env-default:"postgres"`
	SSLMode  string `yaml:"db:sslmode" env-default:"disable"`
}

func MustLoadConfig() *ServiceConfig {
	path := getConfigPath()

	if path == "" {
		panic("empty Server Config path")
	}

	// проверяем, существует ли файл, если нет - паника
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	return &ServiceConfig{
		DB:    *MustLoadDataBaseConfig(path),
		Srv:   *MustLoadServerConfig(path),
		Login: *MustLoadDLoginConfig(path),
	}

}

func MustLoadServerConfig(path string) *ServerConfig { //должна быть паника вместо обработки ошибки
	var cfg ServerConfig

	//запись из файла по пути path в сущность cfg,
	//если возникает ошибка - паникуем

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg

}

func MustLoadDataBaseConfig(path string) *DatabaseConfig {
	var cfg DatabaseConfig

	//запись из файла по пути path в сущность cfg,
	//если возникает ошибка - паникуем
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

func MustLoadDLoginConfig(path string) *LoginCongif {
	var cfg LoginCongif

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// получение пути до файла возможно как через перменную окружения,
// так и через флаг
func getConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
		//добавить лог - каким способом получили путь
	} //else {

	//}

	return res

}

// func MustLoadDataBaseConfig() *ServerConfig {

// }
