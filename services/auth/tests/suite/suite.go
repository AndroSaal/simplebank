package suite

import (
	"context"
	"net"
	"testing"

	v1 "github.com/AndtoSaal/simplebank/services/auth/pb/gateway-auth/v1"
	"github.com/AndtoSaal/simplebank/services/auth/src/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	Cfg        *config.ServiceConfig
	AuthClient v1.AuthClient
}

func NewSuite(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadConfig()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Srv.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	grpcAdress := net.JoinHostPort(cfg.Srv.GRPC.Host, cfg.Srv.GRPC.Port)

	clientConnection, err := grpc.NewClient(grpcAdress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc connection failed %v", err)
	}

	grpcClient := v1.NewAuthClient(clientConnection)

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: grpcClient,
	}

}
