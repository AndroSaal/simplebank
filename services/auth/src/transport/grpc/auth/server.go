//Транспортный слой - Хэндлеры

package auth_transport

import (
	"context"
	"errors"

	grpcAuthV1 "github.com/AndtoSaal/simplebank/services/auth/pb/gateway-auth/v1"
	auth_service "github.com/AndtoSaal/simplebank/services/auth/src/service"

	// "google.golang.org/appengine/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	gRPCstatus "google.golang.org/grpc/status"
)

// структура сервера, инмплементация интерфейса из пакета, сгенеренного протоком
type authServerAPI struct {
	//техническое поле, нужно для обратной совсместимости с proto
	grpcAuthV1.UnimplementedAuthServer
	//интерфейс сервисного слоя
	auth Auth
}

// какими методами должен обладать сервис Auth
type Auth interface {
	LoginExistUser(ctx context.Context, email string, password string) (token string, err error)
	RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error)
}

// cjplfybt
func RegisterAuthServer(gRPC *grpc.Server) {
	grpcAuthV1.RegisterAuthServer(gRPC, &authServerAPI{})
}

func (s *authServerAPI) Loginer(ctx context.Context, req *grpcAuthV1.LoginerRequest) (*grpcAuthV1.LoginerResponse, error) {
	if req.Email == "" {
		return nil, gRPCstatus.Error(codes.InvalidArgument, "email is required pole")
	}

	if req.Password == "" {
		return nil, gRPCstatus.Error(codes.InvalidArgument, "password is required pole")
	}

	token, err := s.auth.LoginExistUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		//auth_service ErrInvalidCredentials - реализовать на сервисном слое
		if errors.Is(err, auth_service.ErrInvalidCredentials) {
			return nil, gRPCstatus.Error(codes.InvalidArgument, "invalid email or password")
		}
		return nil, gRPCstatus.Error(codes.Internal, "failed to login")
	}
	return &grpcAuthV1.LoginerResponse{Token: token}, nil

}

func (s *authServerAPI) Register(ctx context.Context, req *grpcAuthV1.RegisterRequest) (*grpcAuthV1.RegisterResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		//repository.ErrUserExists реализовать
		if errors.Is(err, repository.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &grpcAuthV1.RegisterResponse{UserId: uid}, nil
}

// func IsAdmin(context.Context, *grpcAuthV1.IsAdminRequest) (*grpcAuthV1.IsAdminResponse, error) {
// 	// log.LevelDebug("UNimplemented")
// 	panic("U got wat u want\n")
// }
