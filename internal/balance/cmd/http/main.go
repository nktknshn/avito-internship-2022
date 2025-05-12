package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl/lagging"
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

	cfg, err := config.LoadConfigFromFile(flagConfigPath)

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	var deps *app_impl.AppDeps
	var cleanup func()

	if cfg.GetLagging().GetEnabled() {
		deps, cleanup, err = lagging.NewLaggingDeps(ctx, cfg)
		if err != nil {
			log.Fatalf("NewLaggingDeps: %v", err)
		}
	} else {
		deps, cleanup, err = app_impl.NewDeps(ctx, cfg)
		if err != nil {
			log.Fatalf("NewDeps: %v", err)
		}
	}

	defer cleanup()

	app, err := app_impl.NewApplicationFromDeps(ctx, deps)

	if err != nil {
		log.Fatalf("NewApplicationFromDeps: %v", err)
	}

	server := server.NewHTTPServer(cfg, app)

	if err := server.Init(ctx); err != nil {
		log.Fatalf("Init: %v", err)
	}

	err = server.Run(ctx)
	if err != nil {
		log.Fatalf("Run: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	err = server.Shutdown(ctx)

	if err != nil {
		log.Fatalf("Shutdown: %v", err)
	}
}
