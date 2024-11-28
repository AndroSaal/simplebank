package auth

import (
	"context"

	grpcAuthV1 "github.com/AndtoSaal/simplebank/services/auth/pb/gateway-auth/v1"
	// "google.golang.org/appengine/log"
	"google.golang.org/grpc"
)

type serverAPI struct {
	grpcAuthV1.UnimplementedAuthServer
}

func RegisterAuthServer(gRPC *grpc.Server) {
	grpcAuthV1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *grpcAuthV1.LoginerRequest) (*grpcAuthV1.LoginerResponse, error) {
	// log.LevelDebug("UNimplemented")
	panic("U got wat u want\n")
}

func (s *serverAPI) Register(ctx context.Context, req *grpcAuthV1.RegisterRequest) (*grpcAuthV1.RegisterResponse, error) {
	// log.LevelDebug("UNimplemented")
	panic("U got wat u want\n")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *grpcAuthV1.IsAdminRequest) (*grpcAuthV1.IsAdminResponse, error) {
	// log.LevelDebug("UNimplemented")
	panic("U got wat u want\n")
}
