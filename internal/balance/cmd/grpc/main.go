package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	adaptersGrpc "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var (
	flagConfigPath = "config.yaml"
)

func main() {

	flag.StringVar(&flagConfigPath, "cfg_path", flagConfigPath, "config path")
	flag.Parse()

	ctx := context.Background()

	cfg, err := config.LoadConfig(flagConfigPath)
	if err != nil {
		panic(err)
	}

	app, err := app_impl.NewApplication(ctx, cfg)
	if err != nil {
		panic(err)
	}

	grpcAdapter := adaptersGrpc.New(app.GetApp())

	opts := []grpc.ServerOption{}

	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:                  cfg.GRPC.GetKeepalive().GetTime(),
		Timeout:               cfg.GRPC.GetKeepalive().GetTimeout(),
		MaxConnectionIdle:     cfg.GRPC.GetKeepalive().GetMaxConnectionIdle(),
		MaxConnectionAge:      cfg.GRPC.GetKeepalive().GetMaxConnectionAge(),
		MaxConnectionAgeGrace: cfg.GRPC.GetKeepalive().GetMaxConnectionAgeGrace(),
	}))

	opts = append(opts, grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			// grpc_ctxtags.UnaryServerInterceptor(),
			grpcAdapter.UnaryServerInterceptor(),
		)),
	)

	grpcServer := grpc.NewServer(opts...)

	balance.RegisterBalanceServiceServer(grpcServer, grpcAdapter)

	if cfg.GetMode() == "dev" {
		reflection.Register(grpcServer)
	}

	listen, err := net.Listen("tcp", cfg.GRPC.GetAddr())
	if err != nil {
		panic(err)
	}

	go func() {
		app.GetLogger().Info("GRPC server started", "address", listen.Addr().String())
		err = grpcServer.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	app.GetLogger().Info("GRPC server stopped")

	grpcServer.GracefulStop()

}
