package config

import "google.golang.org/grpc"

type ServerConfig struct {
	Env  string `yaml:"env" env-default:"local"`
	GRPC grpc.Server
}

type DatabaseConfig struct {
	Env      string `yaml:"env"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	UserName string `yaml:"username" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Database string `yaml:"dbname" env-default:"postgres"`
}
