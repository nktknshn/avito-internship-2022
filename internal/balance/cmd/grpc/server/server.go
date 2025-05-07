package server

import (
	"context"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	adaptersGrpc "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"

	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type BalanceGrpcServer struct {
	grpcServer *grpc.Server
	cfg        *config.Config
	app        *app_impl.Application
}

func NewGrpcServer(cfg *config.Config) *BalanceGrpcServer {
	if cfg == nil {
		panic("cfg is nil")
	}
	return &BalanceGrpcServer{
		cfg: cfg,
	}
}

func (s *BalanceGrpcServer) GetServer() *grpc.Server {
	return s.grpcServer
}

func (s *BalanceGrpcServer) GetConfig() *config.Config {
	return s.cfg
}

func (s *BalanceGrpcServer) GetApp() *app_impl.Application {
	return s.app
}

func (s *BalanceGrpcServer) GetLogger() logging.Logger {
	return s.app.GetLogger()
}

func (s *BalanceGrpcServer) Init(ctx context.Context) error {

	var err error

	s.app, err = app_impl.NewApplication(ctx, s.cfg)
	if err != nil {
		return err
	}

	grpcAdapter := adaptersGrpc.New(s.app.GetApp())

	opts := []grpc.ServerOption{}

	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:                  s.cfg.GRPC.GetKeepalive().GetTime(),
		Timeout:               s.cfg.GRPC.GetKeepalive().GetTimeout(),
		MaxConnectionIdle:     s.cfg.GRPC.GetKeepalive().GetMaxConnectionIdle(),
		MaxConnectionAge:      s.cfg.GRPC.GetKeepalive().GetMaxConnectionAge(),
		MaxConnectionAgeGrace: s.cfg.GRPC.GetKeepalive().GetMaxConnectionAgeGrace(),
	}))

	opts = append(opts, grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			// grpc_ctxtags.UnaryServerInterceptor(),
			grpcAdapter.UnaryServerInterceptor(),
		)),
	)

	s.grpcServer = grpc.NewServer(opts...)
	balance.RegisterBalanceServiceServer(s.grpcServer, grpcAdapter)

	if s.cfg.GetMode() == "dev" {
		reflection.Register(s.grpcServer)
	}

	return nil
}

func (s *BalanceGrpcServer) Run(ctx context.Context) error {

	listen, err := net.Listen("tcp", s.cfg.GRPC.GetAddr())
	if err != nil {
		panic(err)
	}

	go func() {
		s.app.GetLogger().Info("GRPC server started", "address", listen.Addr().String())
		err = s.grpcServer.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

	return nil
}

func (s *BalanceGrpcServer) Shutdown(ctx context.Context) error {
	s.grpcServer.GracefulStop()
	return nil
}
