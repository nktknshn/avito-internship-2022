package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
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
		log.Fatalf("LoadConfig: %v", err)
	}

	app, cleanup, err := app_impl.NewApplication(ctx, cfg)

	if err != nil {
		log.Fatalf("NewApplication: %v", err)
	}

	defer cleanup()

	serv := server.NewGrpcServer(cfg, app)

	if err := serv.Init(ctx); err != nil {
		log.Fatalf("Init: %v", err)
	}

	if err := serv.Run(ctx); err != nil {
		log.Fatalf("Run: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	serv.GetLogger().Info("GRPC server stopped")

	serv.Shutdown(ctx)

}
