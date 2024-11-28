package auth_service

import (
	"context"
	"log/slog"
	"time"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
)

type UserSaver interface {
	SaveUser(ctx context.Context, email string, password_Hash []byte) (userId int, err error)
}

type UserGetter interface {
	GetUser(ctx context.Context, email string) (*models.User, error)
}

type UserHandler interface {
	UserGetter
	UserSaver
}

// реализация интерфейса Auth из services/auth/src/transport/grpc/auth/server.go
type AuthService struct {
	log         *slog.Logger
	userHandler UserHandler
	tokenTTL    time.Duration
}

func New(
	log *slog.Logger,
	userHandler UserHandler,
	tokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		userHandler: userHandler,
		log:         log,
		tokenTTL:    tokenTTL, // Время жизни возвращаемых токенов
	}
}

func (as *AuthService) RegisterNewUser()
