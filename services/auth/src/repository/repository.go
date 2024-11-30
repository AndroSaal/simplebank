package auth_repository

import (
	"context"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
	"github.com/jmoiron/sqlx"
)

// интерфейс, которым определяющий тип, которым должна обладать конкретная бд
// чтобы его реализовывать
type Repository interface {
	SaveUser(ctx context.Context, email string, password_Hash []byte) (userId int, err error)
	GetUser(ctx context.Context, email string) (models.User, error)
}

// тип, релизующий интерфейс сервисвного слоя UserRepository
type AuthUserRepo struct {
	repo Repository
}

// конструктор типа AuthUserRepo
func NewAuthRepository(db *sqlx.DB) *AuthUserRepo {
	return &AuthUserRepo{repo: NewAuthPostgresRepo(db)}
}
