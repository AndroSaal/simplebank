package auth_repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
	"github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	"github.com/jmoiron/sqlx"
)

type AuthPostgresDB struct {
	db *sqlx.DB
}

func NewAuthPostgresRepo(cfgDataBase config.DatabaseConfig, logger *slog.Logger) *AuthPostgresDB {
	db, err := NewPostgresDB(cfgDataBase)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot connect to databse : %s ", (err).Error()))
	}
	return &AuthPostgresDB{db: db}
}

func (p *AuthPostgresDB) Stop() error {
	err := p.db.Close()
	return err
}

func (p *AuthPostgresDB) SaveUser(
	ctx context.Context, email string, passwordHash []byte) (userId int64, err error) {

	var id int64

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

func (r *AuthPostgresDB) GetUserInfo(ctx context.Context, userId int64) (bool, error) {

	var isAdmin bool

	query := fmt.Sprintf("SELECT is_Admin FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&isAdmin, query, userId /*$1 в query*/)

	return isAdmin, err
}
