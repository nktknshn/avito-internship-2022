package server

import (
	"context"
	"errors"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	adaptersGrpc "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
)

type BalanceGrpcServer struct {
	grpcServer *grpc.Server
	cfg        *config.Config
	app        *app_impl.Application
}

func NewGrpcServer(cfg *config.Config, app *app_impl.Application) *BalanceGrpcServer {

	if cfg == nil {
		panic("cfg is nil")
	}

	if app == nil {
		panic("app is nil")
	}

	return &BalanceGrpcServer{
		cfg: cfg,
		app: app,
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

func (s *BalanceGrpcServer) Init(_ context.Context) error {

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

func (s *BalanceGrpcServer) Run(_ context.Context) error {

	if s.app == nil {
		return errors.New("Init() must be called before Run()")
	}

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

func (s *BalanceGrpcServer) Shutdown(_ context.Context) error {
	s.grpcServer.GracefulStop()
	return nil
}
