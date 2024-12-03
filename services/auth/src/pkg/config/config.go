//конфигурация сервиса авторизации
//с помощью библиотеки github.com/ilyakaznacheev/cleanenv@latest

package config

import (
	"flag"
	"os"
	"time"

	"github.com/spf13/viper"
	// "github.com/ilyakaznacheev/cleanenv"
	// "google.golang.org/grpc"
)

type ServiceConfig struct {
	DB    DatabaseConfig
	Srv   ServerConfig
	Login LoginConfig
}

// конфигурация сервера
type ServerConfig struct {
	Env  string      `yaml:"env" env-default:"local"`
	GRPC GRPCConfing `yml:"grpc"`
}

// конфигурация grpc сервера
type GRPCConfing struct {
	Port    string        `yaml:"port" env-default:"50051"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

type LoginConfig struct {
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"15h"`
}

// конфигурация базы
type DatabaseConfig struct {
	Env      string `yaml:"env"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	UserName string `yaml:"username" env-default:"postgres"`
	Password string `yaml:"password" env-default:"qwerty"`
	Database string `yaml:"dbname" env-default:"postgres"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

func MustLoadConfig() *ServiceConfig {
	path, name := getConfigPath()
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	// viper.SetConfigName()

	var (
		srvconf   ServerConfig
		dbconf    DatabaseConfig
		loginconf LoginConfig
	)

	if path == "" {
		panic("empty Server Config path")
	}

	// проверяем, существует ли файл, если нет - паника
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	if err := viper.ReadInConfig(); err != nil {
		panic("can not read config file: " + err.Error())
	}

	if err := viper.UnmarshalKey("grpc", &srvconf); err != nil {
		panic("troubles with grpc config: " + err.Error())
	}

	if err := viper.UnmarshalKey("db", &dbconf); err != nil {
		panic("troubles with grpc config: " + err.Error())
	}

	if err := viper.UnmarshalKey("login", &loginconf); err != nil {
		panic("troubles with grpc config: " + err.Error())
	}

	return &ServiceConfig{
		DB:    dbconf,
		Srv:   srvconf,
		Login: loginconf,
	}

}

// func MustLoadServerConfig(path string) *ServerConfig { //должна быть паника вместо обработки ошибки
// 	var cfg ServerConfig

// 	//запись из файла по пути path в сущность cfg,
// 	//если возникает ошибка - паникуем

// 	if err := viper.UnmarshalKey(path, &cfg); err != nil {
// 		panic("config path is empty: " + err.Error())
// 	}

// 	return &cfg

// }

// func MustLoadDataBaseConfig(path string) *DatabaseConfig {
// 	var cfg DatabaseConfig

// 	//запись из файла по пути path в сущность cfg,
// 	//если возникает ошибка - паникуем
// 	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
// 		panic("config path is empty: " + err.Error())
// 	}

// 	return &cfg
// }

// func MustLoadDLoginConfig(path string) *LoginCongif {
// 	var cfg LoginCongif

// 	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
// 		panic("config path is empty: " + err.Error())
// 	}

// 	return &cfg
// }

// получение пути до файла возможно как через перменную окружения,
// так и через флаг
func getConfigPath() (string, string) {
	var (
		path string
		name string
	)

	flag.StringVar(&path, "config_path", "", "path to config file")
	flag.StringVar(&name, "config_name", "", "name of config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
		//добавить лог - каким способом получили путь
	} //else

	if name == "" {
		name = os.Getenv("CONFIG_NAME")
		//добавить лог - каким способом получили путь
	}

	//}

	return path, name

}

// func MustLoadDataBaseConfig() *ServerConfig {

// }
