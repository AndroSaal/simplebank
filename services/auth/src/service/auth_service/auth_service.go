package auth_service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/AndtoSaal/simplebank/services/auth/src/entities/models"
	"github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	jwtAuth "github.com/AndtoSaal/simplebank/services/auth/src/pkg/jwt"
	log "github.com/AndtoSaal/simplebank/services/auth/src/pkg/logger"
	auth_repository "github.com/AndtoSaal/simplebank/services/auth/src/repository"
	auth_service "github.com/AndtoSaal/simplebank/services/auth/src/service/auth_service/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserAdder interface {
	SaveUser(ctx context.Context, email string, password_Hash []byte) (userId int64, err error)
}

type UserGetter interface {
	GetUser(ctx context.Context, email string) (models.User, error)
}

type StopperConnecion interface {
	Stop() error
}

// интерфейс для хранилища (repository слой)
type UserRepository interface {
	UserGetter
	UserAdder
	StopperConnecion
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
	serviceConfig config.ServiceConfig,
) *AuthService {
	return &AuthService{
		userRepositoryHandler: auth_repository.NewAuthPostgresRepo(serviceConfig.DB, log),
		log:                   log,
		tokenTTL:              serviceConfig.Login.TokenTTL, // Время жизни возвращаемых токенов
	}
}

// метод для реализации интерфейса из транспортного слоя - регистрация нового полььзователя
func (as *AuthService) RegisterNewUser(
	ctx context.Context,
	email string, password string) (userId int64, err error) {

	const tracelog = "auth_service.RegisterNewUser"

	logLocal := as.log.With(
		slog.String("tracelog", tracelog),
		slog.String("email", email),
	)

	logLocal.Info("Register new user")
	//добавить пароля с помощью паkета crypto

	//генерируем хэш и соль для пароля
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

func (as *AuthService) LoginExistUser(ctx context.Context, email string, password string) (token string, err error) {

	const tracelog = "auth_service.LoginExistUser"

	logLocal := as.log.With(
		slog.String("tracelog", tracelog),
		slog.String("email", email),
		//пароль нужно засекретить, если проект станет чуть рентабильнее
		slog.String("password", password),
	)

	logLocal.Info("Login exist user")

	//Получение пользователя из базы
	user, err := as.userRepositoryHandler.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, auth_repository.ErrUserNotFound) {
			as.log.Warn("user not found", log.Err(err))
			return "", fmt.Errorf("%s: %w", tracelog, auth_service.ErrInvalidCredentials)
		}

		as.log.Error("failed to get user", log.Err(err))
		return "", fmt.Errorf("%s: %w", tracelog, err)
	}

	//Проверка пароля для найденного пользователя
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		as.log.Info("invalid credentials", log.Err(err))

		return "", fmt.Errorf("%s: %w", tracelog, auth_service.ErrInvalidCredentials)
	}

	logLocal.Info("user logged in successfully")

	//создание токена для пользователя
	token, err = jwtAuth.NewToken(user, as.tokenTTL)
	if err != nil {
		as.log.Error("failed to create token", log.Err(err))
		return "", fmt.Errorf("%s: %w", tracelog, err)
	}

	return token, nil

}

func (as *AuthService) Stop() error {
	err := as.userRepositoryHandler.Stop()
	return err
}
