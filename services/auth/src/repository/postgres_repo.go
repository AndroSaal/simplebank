package auth_repository

import (
	"fmt"

	"github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

// кокнструктор из конфига - инициализация
func NewPostgresDB(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	//заполняем структурку в конструкторе
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.Database, cfg.SSLMode))
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
