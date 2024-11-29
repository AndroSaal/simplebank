package auth_service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
)

type UserAdder interface {
	SaveUser(ctx context.Context, email string, password_Hash []byte) (userId int, err error)
}

type UserGetter interface {
	GetUser(ctx context.Context, email string) (*models.User, error)
}

// интерфейс для хранилища (repository слой)
type UserRepository interface {
	UserGetter
	UserAdder
}

// реализация интерфейса Auth из services/auth/src/transport/grpc/auth/server.go
// интерфейса из транспортного слоя
type AuthService struct {
	log *slog.Logger
	//методы уровня базы данных (repository)
	userRepositoryHandler UserRepository
	tokenTTL              time.Duration
}

// конструктор AuthService
func NewAuthService(
	log *slog.Logger,
	userRepositoryHandler UserRepository,
	tokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		userRepositoryHandler: userRepositoryHandler,
		log:                   log,
		tokenTTL:              tokenTTL, // Время жизни возвращаемых токенов
	}
}

// метод для реализации интерфейса из транспортного слоя
func (as *AuthService) RegisterNewUser(
	ctx context.Context,
	email string, passwordHash []byte) (userId int, err error) {

	const tracelog = "auth_service.RegisterNewUser"

	logLocal := as.log.With(
		slog.String("tracelog", tracelog),
		slog.String("email", email),
	)

	logLocal.Info("Register new user")
	//добавить пароля с помощью паkета crypto

	if err != nil {
		logLocal.Error("failed to generate password hash", log.Err(err))

		return 0, fmt.Errorf("%s: %w", tracelog, err)
	}

	userId, err = as.userRepositoryHandler.SaveUser(ctx, email, passwordHash)
	if err != nil {
		logLocal.Error("failed to save user", log.Err(err))
		return 0, fmt.Errorf("%s: %w", tracelog, err)
	}

	return userId, nil

}
