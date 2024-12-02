package auth_transport

import (
	"context"
	"log/slog"

	grpcAuthV1 "github.com/AndtoSaal/simplebank/services/auth/pb/gateway-auth/v1"
	"github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// обертка для type authServerAPI struct
type AuthTransport struct {
	gRPCServer *grpc.Server
	log        *slog.Logger
	port       string
}

func NewAuthTransport(
	log *slog.Logger, cfg config.ServiceConfig) *AuthTransport {

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

	authServerAPI := NewAuthServerAPI(log, cfg)

	grpcAuthV1.RegisterAuthServer(gRPCServer, authServerAPI)

	return &AuthTransport{
		gRPCServer: gRPCServer,
		log:        log,
		port:       cfg.Srv.GRPC.Port,
	}
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
