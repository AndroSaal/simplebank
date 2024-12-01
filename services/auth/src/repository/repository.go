package auth_repository

import (
	"context"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
	"github.com/jmoiron/sqlx"
)

// интерфейс, которым определяющий тип, которым должна обладать конкретная бд
// чтобы его реализовывать
type RepositoryAuth interface {
	SaveUser(ctx context.Context, email string, password_Hash []byte) (userId int64, err error)
	GetUser(ctx context.Context, email string) (models.User, error)
}

type RepositoryUsrInfo interface {
	GetUserInfo(ctx context.Context, id int64) (bool, error)
}

// тип, релизующий интерфейс сервисвного слоя UserRepository
type AuthUserRepo struct {
	repoAuth RepositoryAuth
	repoInf  RepositoryUsrInfo
}

// конструктор типа AuthUserRepo
func NewAuthRepository(db *sqlx.DB) *AuthUserRepo {
	return &AuthUserRepo{
		repoAuth: NewAuthPostgresRepo(db),
		repoInf:  NewAuthPostgresRepo(db),
	}
}
