package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/grpc/server"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
)

var (
	flagConfigPath = "config.yaml"
)

func main() {

	flag.StringVar(&flagConfigPath, "cfg_path", flagConfigPath, "config path")
	flag.Parse()

	ctx := context.Background()

	cfg, err := config.LoadConfigFromFile(flagConfigPath)
	if err != nil {
		panic(err)
	}

	serv := server.NewGrpcServer(cfg)

	if err := serv.Init(ctx); err != nil {
		panic(err)
	}

	if err := serv.Run(ctx); err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	serv.GetLogger().Info("GRPC server stopped")

	serv.Shutdown(ctx)

}
