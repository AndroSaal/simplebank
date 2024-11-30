package auth_repository

import (
	"context"
	"fmt"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgresDB struct {
	db *sqlx.DB
}

func NewAuthPostgresRepo(db *sqlx.DB) *AuthPostgresDB {
	return &AuthPostgresDB{db: db}
}

func (p *AuthPostgresDB) SaveUser(
	ctx context.Context, email string, passwordHash []byte) (userId int, err error) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := p.db.QueryRow(query, email, string(passwordHash))

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// TODO : Добавить ошибки
func (r *AuthPostgresDB) GetUser(ctx context.Context, email string) (models.User, error) {

	var (
		user models.User
	)

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, email /*$1 в query*/)

	return user, err
}
