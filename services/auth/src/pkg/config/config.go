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

// конфигурация сервера
type ServerConfig struct {
	Env  string      `yaml:"env" env-default:"local"`
	GRPC GRPCConfing `yml:"grpc"`
}

// конфигурация grpc сервера
type GRPCConfing struct {
	Port    int           `yaml:"port"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

// конфигурация базы
type DatabaseConfig struct {
	Env      string `yaml:"env"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	UserName string `yaml:"username" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Database string `yaml:"dbname" env-default:"postgres"`
}

func MustLoadServerConfig() *ServerConfig { //должна быть паника вместо обработки ошибки
	path := getConfigPath()
	if path == "" {
		panic("empty Server Config path")
	}

	// проверяем, существует ли файл, если нет - паника
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg ServerConfig

	//запись из файла по пути path в сущность cfg,
	//если возникает ошибка - паникуем
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
