package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
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

	handlers := adaptersHttp.NewHttpAdapter(&app.Application)
	router := router.NewMuxRouter(handlers)

	mux := http.NewServeMux()

	prefix := cfg.GetHTTP().GetApiPrefix()

	mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	mux.Handle("/metrics", app.MetricsHandler)

	fmt.Println(prefix)

	app.Logger.Info(ctx, "Starting server on", "addr", cfg.GetHTTP().GetAddr())

	http.ListenAndServe(cfg.GetHTTP().GetAddr(), mux)

	_ = app
}
