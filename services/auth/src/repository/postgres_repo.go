package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable = "users"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	//заполняем структурку в конструкторе
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	fmt.Printf("struct %v\n", cfg)

	//методом Ping проверяем, можем ли мы достучаться до нашей БД
	if err = db.Ping(); err != nil {
		return nil, err
	}

	//Успешное завершение - возвращаем экземпляр БД
	return db, nil
}
