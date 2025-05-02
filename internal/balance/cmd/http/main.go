package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	balanceRouter "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
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
	apiRouter := NewGorillaRouter()

	apiRouter.Use(NewMiddlewareRecover(
		func(ctx context.Context, w http.ResponseWriter, r *http.Request, err any) {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		},
	))

	apiRouter.Use(NewOptionsMiddleware())

	balanceRouter.AttachHandlers(apiRouter, handlers)

	apiPrefix := cfg.GetHTTP().GetApiPrefix()

	server := http.NewServeMux()

	if cfg.GetHTTP().GetSwagger().GetEnabled() {
		app.Logger.Info(ctx, "Swagger enabled")
		swaggerRouter := mux.NewRouter()
		swaggerRouter.PathPrefix(cfg.GetHTTP().GetSwagger().GetPath()).
			Handler(swaggerHandler(cfg))
		server.Handle("/", swaggerRouter)
	}

	server.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, apiRouter))
	server.Handle("/metrics", app.MetricsHandler)

	app.Logger.Info(ctx, "Starting server on", "addr", cfg.GetHTTP().GetAddr())

	http.ListenAndServe(cfg.GetHTTP().GetAddr(), server)

}
