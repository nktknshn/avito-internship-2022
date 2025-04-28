package main

import (
	"context"
	"flag"
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/pkg/config_cleanenv"
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

	handlers := adaptersHttp.NewHttpAdapter(&app.Application)
	router := router.NewMuxRouter(handlers)

	mux := http.NewServeMux()

	prefix := cfg.GetHTTP().GetApiPrefix()

	mux.Handle(prefix, http.StripPrefix(prefix, router))
	mux.Handle("/metrics", app.MetricsHandler)

	http.ListenAndServe(cfg.GetHTTP().GetAddr(), mux)

	_ = app
}
