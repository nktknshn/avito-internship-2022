package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	adapters_grpc "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/nktknshn/avito-internship-2022/pkg/config_cleanenv"
	"github.com/nktknshn/avito-internship-2022/pkg/server_grpc"
	"google.golang.org/grpc"
)

var (
	flagConfigPath = "config.yaml"
)

func main() {

	flag.StringVar(&flagConfigPath, "cfg_path", flagConfigPath, "config path")
	flag.Parse()

	ctx := context.Background()

	cfg, err := config_cleanenv.LoadConfigType[config.Config](flagConfigPath)

	if err != nil {
		panic(err)
	}

	app, err := app_impl.NewApplication(ctx, &cfg)

	if err != nil {
		panic(err)
	}

	server := server_grpc.RunGRPCServerOnAddr(
		cfg.GRPC.GetAddr(),
		app.Logger,
		func(server *grpc.Server) {
			grpcServer := adapters_grpc.NewGrpcServer(app.Application)
			balance.RegisterBalanceServiceServer(server, grpcServer)
		},
	)

	app.Logger.Info(ctx, "GRPC server started on %s", "address", server.Listen.Addr().String())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	app.Logger.Info(ctx, "GRPC server stopped")

	server.GrpcServer.GracefulStop()

}
