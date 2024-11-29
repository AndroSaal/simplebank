package repository

import "context"

// интерфейс, которым определяющий тип, которым должна обладать конкретная бд
// чтобы его реализовывать
type Repository interface {
	SaveUser(ctx context.Context, email string, password_Hash []byte) (userId int, err error)
	GetUser(ctx context.Context, email string) (userId int, err error)
}

// тип, релизующий интерфейс сервисвного слоя UserRepository
type AuthUserRepo struct {
	repo Repository
}

// конструктор типа AuthUserRepo
func NewAuthUserHandler(repo Repository) *AuthUserRepo {
	return &AuthUserRepo{repo: repo}
}
