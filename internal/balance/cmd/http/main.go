package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http/server"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
)

var (
	flagConfigPath = "config.yaml"
)

func main() {

	flag.StringVar(&flagConfigPath, "cfg_path", flagConfigPath, "config path")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig(flagConfigPath)

	if err != nil {
		panic(err)
	}

	server := server.NewHttpServer(cfg)

	if err := server.Init(ctx); err != nil {
		panic(err)
	}

	err = server.Run(ctx)
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	err = server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
