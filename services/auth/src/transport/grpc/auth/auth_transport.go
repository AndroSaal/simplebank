package auth_transport

import (
	"context"
	"log/slog"

	authgrpc "grpc-service-ref/internal/grpc/auth"

	"github.com/AndtoSaal/simplebank/services/auth/src/service/auth_service"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
)

type AuthTransport struct {
	gRPCServer *grpc.Server
	log        *slog.Logger
	port       string
}

func NewAuthTransport(log *slog.Logger, authService *auth_service.AuthService, port int) *AuthTransport {

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
	))

	authgrpc.RegisterAuthServiceServer(gRPCServer, authService)

	return &AuthTransport{
		gRPCServer: gRPCServer,
		log:        log,
		port:       string(port),
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
