//Транспортный слой - Хэндлеры

package auth_transport

import (
	"context"
	"errors"
	"log/slog"

	grpcAuthV1 "github.com/AndtoSaal/simplebank/services/auth/pb/gateway-auth/v1"
	"github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	"github.com/AndtoSaal/simplebank/services/auth/src/service/auth_service"
	authServiceError "github.com/AndtoSaal/simplebank/services/auth/src/service/auth_service/errors"
	"github.com/AndtoSaal/simplebank/services/auth/src/service/usrInfo_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// структура сервера, инмплементация интерфейса из пакета, сгенеренного протоком
type AuthServerAPI struct {
	//техническое поле, нужно для обратной совсместимости с proto
	grpcAuthV1.UnimplementedAuthServer
	//интерфейс сервисного слоя
	auth     Auth
	userInfo UserINFO
}

func NewAuthServerAPI(log *slog.Logger, serviceConfig config.ServiceConfig) *AuthServerAPI {
	return &AuthServerAPI{
		auth:     auth_service.NewAuthService(log, serviceConfig),
		userInfo: usrInfo_service.NewUserInfoService(log, serviceConfig),
	}
}

// какими методами должен обладать сервис Auth
type Auth interface {
	LoginExistUser(ctx context.Context, email string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error)
}

// структура информации о пользователе (userInfo)
type UserINFO interface {
	IsAdminById(ctx context.Context, userID int64) (isAdmin bool, err error)
}

func (s AuthServerAPI) Loginer(ctx context.Context, req *grpcAuthV1.LoginerRequest) (*grpcAuthV1.LoginerResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required field")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required field")
	}

	token, err := s.auth.LoginExistUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		//auth_service ErrInvalidCredentials - реализовать на сервисном слое
		if errors.Is(err, authServiceError.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}
		return nil, status.Error(codes.Internal, "failed to login")
	}
	return &grpcAuthV1.LoginerResponse{Token: token}, nil

}

func (s AuthServerAPI) Register(ctx context.Context, req *grpcAuthV1.RegisterRequest) (*grpcAuthV1.RegisterResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required field")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required field")
	}

	uid, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		//repository.ErrUserExists реализовать на сервисном слое
		if errors.Is(err, authServiceError.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &grpcAuthV1.RegisterResponse{UserId: uid}, nil
}
func (s AuthServerAPI) IsAdminChecker(ctx context.Context, req *grpcAuthV1.IsAdminRequest) (*grpcAuthV1.IsAdminResponse, error) {
	if req.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id is required field")
	}
	isAdmin, err := s.userInfo.IsAdminById(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to check user")
	}
	return &grpcAuthV1.IsAdminResponse{IsAdmin: isAdmin}, nil
}
