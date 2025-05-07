package server

import (
	"context"
	"errors"
	"net/http"
	"path"
	"time"

	middleware "github.com/go-chi/chi/v5/middleware"
	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	balanceRouter "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	ergo "github.com/nktknshn/go-ergo-handler"
	ergoChi "github.com/nktknshn/go-ergo-handler/adapters/chi"

	"github.com/go-chi/cors"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http/chi"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
)

type BalanceHttpServer struct {
	server    *http.Server
	serveMux  *http.ServeMux
	cfg       *config.Config
	app       *app_impl.Application
	apiRouter *chi.ChiRouter
}

func NewHttpServer(cfg *config.Config) *BalanceHttpServer {
	return &BalanceHttpServer{
		cfg: cfg,
	}
}

func (s *BalanceHttpServer) GetServer() *http.Server {
	return s.server
}

func (s *BalanceHttpServer) GetConfig() *config.Config {
	return s.cfg
}

func (s *BalanceHttpServer) GetApp() *app_impl.Application {
	return s.app
}

func (s *BalanceHttpServer) GetServeMux() *http.ServeMux {
	return s.serveMux
}

func (s *BalanceHttpServer) GetLogger() logging.Logger {
	return s.app.GetLogger()
}

func (s *BalanceHttpServer) Init(ctx context.Context) error {

	var err error

	s.app, err = app_impl.NewApplication(ctx, s.cfg)
	if err != nil {
		return err
	}

	logger := s.app.GetLogger()

	ergo.SetVarsGetter(ergoChi.New())

	apiRouter := chi.NewChiRouter()

	apiRouter.Use(middleware.Recoverer)
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.RequestID)
	apiRouter.Use(middleware.RealIP)

	mwCors := cors.New(cors.Options{
		AllowedOrigins:   s.cfg.GetHTTP().GetCors().GetAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	apiRouter.Use(mwCors.Handler)

	handlers := adaptersHttp.NewHttpAdapter(s.app.GetApp())
	balanceRouter.Attach(apiRouter, handlers)

	apiPrefix := path.Clean(s.cfg.GetHTTP().GetApiPrefix())

	s.serveMux = http.NewServeMux()

	if s.cfg.GetHTTP().GetSwagger().GetEnabled() {
		swaggerPath := s.cfg.GetHTTP().GetSwagger().GetPath()
		swaggerPath = path.Clean(swaggerPath)
		logger.Info("Swagger enabled", "path", swaggerPath)
		s.serveMux.Handle(swaggerPath+"/", swaggerHandler(s.cfg))
	}

	s.serveMux.Handle(
		apiPrefix+"/",
		http.StripPrefix(apiPrefix, apiRouter.GetHandler()),
	)

	logger.Info("Binding revenue exporter download", "url", s.cfg.GetUseCases().GetReportRevenueExport().GetURL())

	s.serveMux.Handle(
		s.cfg.GetUseCases().GetReportRevenueExport().GetURL(),
		s.app.GetRevenueExporterHandler(),
	)

	if s.app.GetMetricsHandler() != nil {
		s.serveMux.Handle("/metrics", s.app.GetMetricsHandler())
	}

	logger.Info("Starting server on", "addr", s.cfg.GetHTTP().GetAddr())

	s.server = &http.Server{
		Addr:         s.cfg.GetHTTP().GetAddr(),
		Handler:      http.TimeoutHandler(s.serveMux, s.cfg.GetHTTP().GetHandlerTimeout(), "Timeout"),
		ReadTimeout:  s.cfg.GetHTTP().GetReadTimeout(),
		WriteTimeout: s.cfg.GetHTTP().GetWriteTimeout(),
	}

	if s.cfg.GetHTTP().GetTLS().GetEnabled() {
		setTLS(s.server, s.cfg.GetHTTP().GetTLS())
	}

	return nil
}

func (s *BalanceHttpServer) Run(ctx context.Context) error {
	logger := s.app.GetLogger()

	s.app.RunRevenueExporterCleanup(ctx)

	go func() {
		var err error
		if s.cfg.GetHTTP().GetTLS().GetEnabled() {
			logger.Info("Starting server with TLS", "cert_file", s.cfg.GetHTTP().GetTLS().GetCertFile(), "key_file", s.cfg.GetHTTP().GetTLS().GetKeyFile())

			err = s.server.ListenAndServeTLS(
				s.cfg.GetHTTP().GetTLS().GetCertFile(),
				s.cfg.GetHTTP().GetTLS().GetKeyFile(),
			)
		} else {
			err = s.server.ListenAndServe()
		}

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Failed to start server", "error", err)
		}
	}()

	return nil
}

func (s *BalanceHttpServer) Shutdown(ctx context.Context) error {

	logger := s.app.GetLogger()

	logger.Info("Shutting down server")

	shutdownCtx, shutdown := context.WithTimeout(ctx, time.Second*10)
	defer shutdown()

	err := s.server.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error("Failed to stop server", "error", err)
	}

	logger.Info("Server stopped")

	return nil
}

func setTLS(server *http.Server, cfg *config.ConfigTLS) {
	// TODO
	// server.TLSConfig = &tls.Config{
	// 	Certificates: []tls.Certificate{cfg.GetCertificates()},
	// }
}
