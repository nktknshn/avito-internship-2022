package main

import (
	"context"
	"flag"
	"net/http"

	middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	balanceRouter "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http/chi"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	ergo "github.com/nktknshn/go-ergo-handler"
	ergoChi "github.com/nktknshn/go-ergo-handler/adapters/chi"
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

	ergo.SetVarsGetter(ergoChi.New())

	apiRouter := chi.NewChiRouter()

	apiRouter.Use(middleware.Recoverer)
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.RequestID)
	apiRouter.Use(middleware.RealIP)

	mwCors := cors.New(cors.Options{
		AllowedOrigins:   cfg.GetHTTP().GetCors().GetAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	apiRouter.Use(mwCors.Handler)

	handlers := adaptersHttp.NewHttpAdapter(&app.Application)
	balanceRouter.Attach(apiRouter, handlers)

	apiPrefix := cfg.GetHTTP().GetApiPrefix()

	server := http.NewServeMux()

	if cfg.GetHTTP().GetSwagger().GetEnabled() {
		swaggerPath := cfg.GetHTTP().GetSwagger().GetPath()
		app.GetLogger().Info("Swagger enabled", "path", swaggerPath)
		swaggerRouter := mux.NewRouter()
		swaggerRouter.PathPrefix(swaggerPath).
			Handler(swaggerHandler(cfg))
		server.Handle("/", swaggerRouter)
	}

	server.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, apiRouter.GetHandler()))

	app.GetLogger().Info("Binding revenue exporter", "url", cfg.GetUseCases().GetReportRevenueExport().GetURL())

	server.Handle(
		cfg.GetUseCases().GetReportRevenueExport().GetURL(),
		app.GetRevenueExporterHandler(),
	)

	app.RunRevenueExporterCleanup(ctx)

	if app.GetMetricsHandler() != nil {
		server.Handle("/metrics", app.GetMetricsHandler())
	}

	app.GetLogger().Info("Starting server on", "addr", cfg.GetHTTP().GetAddr())

	http.ListenAndServe(cfg.GetHTTP().GetAddr(), server)

}
