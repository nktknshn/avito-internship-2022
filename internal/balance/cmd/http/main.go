package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"

	middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig(flagConfigPath)
	if err != nil {
		panic(err)
	}

	app, err := app_impl.NewApplication(ctx, cfg)
	if err != nil {
		panic(err)
	}

	logger := app.GetLogger()

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

	apiPrefix := path.Clean(cfg.GetHTTP().GetApiPrefix())

	serveMux := http.NewServeMux()

	if cfg.GetHTTP().GetSwagger().GetEnabled() {
		swaggerPath := cfg.GetHTTP().GetSwagger().GetPath()
		logger.Info("Swagger enabled", "path", swaggerPath)
		serveMux.Handle(path.Clean(swaggerPath)+"/", swaggerHandler(cfg))
	}

	serveMux.Handle(apiPrefix+"/", http.StripPrefix(apiPrefix, apiRouter.GetHandler()))

	logger.Info("Binding revenue exporter download", "url", cfg.GetUseCases().GetReportRevenueExport().GetURL())

	serveMux.Handle(
		cfg.GetUseCases().GetReportRevenueExport().GetURL(),
		app.GetRevenueExporterHandler(),
	)

	if app.GetMetricsHandler() != nil {
		serveMux.Handle("/metrics", app.GetMetricsHandler())
	}

	logger.Info("Starting server on", "addr", cfg.GetHTTP().GetAddr())

	server := &http.Server{
		Addr:         cfg.GetHTTP().GetAddr(),
		Handler:      http.TimeoutHandler(serveMux, cfg.GetHTTP().GetHandlerTimeout(), "Timeout"),
		ReadTimeout:  cfg.GetHTTP().GetReadTimeout(),
		WriteTimeout: cfg.GetHTTP().GetWriteTimeout(),
	}

	if cfg.GetHTTP().GetTLS().GetEnabled() {
		setTLS(server, cfg.GetHTTP().GetTLS())
	}

	app.RunRevenueExporterCleanup(ctx)

	go func() {

		if cfg.GetHTTP().GetTLS().GetEnabled() {
			// TODO
		}

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Failed to start server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("Shutting down server")

	shutdownCtx, shutdown := context.WithTimeout(ctx, time.Second*10)
	defer shutdown()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error("Failed to stop server", "error", err)
	}

	logger.Info("Server stopped")

}
