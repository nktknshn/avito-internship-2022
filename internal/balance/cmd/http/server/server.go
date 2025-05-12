package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"path"
	"time"

	middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
	ergo "github.com/nktknshn/go-ergo-handler"
	ergoChi "github.com/nktknshn/go-ergo-handler/adapters/chi"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	balanceRouter "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http/chi"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
)

var (
	maxHeaderBytes  = 1 << 20
	shutdownTimeout = time.Second * 10
)

type BalanceHTTPServer struct {
	server     *http.Server
	rootRouter *chi.ChiRouter
	cfg        *config.Config
	app        *app_impl.Application
}

func NewHTTPServer(cfg *config.Config, app *app_impl.Application) *BalanceHTTPServer {

	if cfg == nil {
		panic("cfg is nil")
	}

	if app == nil {
		panic("app is nil")
	}

	return &BalanceHTTPServer{
		cfg: cfg,
		app: app,
	}
}

func (s *BalanceHTTPServer) GetServer() *http.Server {
	return s.server
}

func (s *BalanceHTTPServer) GetHandler() http.Handler {
	return s.rootRouter.GetHandler()
}

func (s *BalanceHTTPServer) GetConfig() *config.Config {
	return s.cfg
}

func (s *BalanceHTTPServer) GetApp() *app_impl.Application {
	return s.app
}

func (s *BalanceHTTPServer) GetLogger() logging.Logger {
	return s.app.GetLogger()
}

func (s *BalanceHTTPServer) Init(_ context.Context) error {

	logger := s.app.GetLogger()

	ergo.SetVarsGetter(ergoChi.New())

	rootRouter := chi.NewChiRouter()
	s.rootRouter = rootRouter

	mwCors := cors.New(cors.Options{
		AllowedOrigins:   s.cfg.GetHTTP().GetCORS().GetAllowedOrigins(),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		//nolint:mnd // TODO: config
		MaxAge: 300,
	})

	httpLogger := httplog.NewLogger("balance-microservice", httplog.Options{
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		JSON:             true,
		TimeFieldName:    "time",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			// "version": "v1.0-81aa4244d9fc8076a",
			"env": "dev",
		},
		QuietDownRoutes: []string{
			"/ping",
			"/metrics",
		},
		//nolint:mnd // TODO: config
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})

	rootRouter.Use(middleware.Heartbeat("/ping"))
	rootRouter.Use(middleware.Recoverer)
	rootRouter.Use(middleware.RequestID)
	rootRouter.Use(middleware.RealIP)
	rootRouter.Use(mwCors.Handler)
	rootRouter.Use(httplog.RequestLogger(httpLogger))

	apiRouter := chi.NewChiRouter()

	balanceRouter.Attach(
		apiRouter,
		adaptersHttp.NewHTTPAdapter(s.app.GetApp()),
	)

	if s.cfg.GetHTTP().GetSwagger().GetEnabled() {
		swaggerPath := s.cfg.GetHTTP().GetSwagger().GetPath()
		logger.Info("Swagger enabled", "path", swaggerPath)
		rootRouter.GetChi().Mount(swaggerPath, swaggerHandler(s.cfg))
	}

	apiPrefix := path.Clean(s.cfg.GetHTTP().GetAPIPrefix())

	rootRouter.GetChi().Mount(apiPrefix, apiRouter.GetHandler())

	logger.Info("Binding revenue exporter download", "url", s.cfg.GetUseCases().GetReportRevenueExport().GetURL())

	rootRouter.GetChi().Mount(
		s.cfg.GetUseCases().GetReportRevenueExport().GetURL(),
		s.app.GetRevenueExporterHandler(),
	)

	if s.app.GetMetricsHandler() != nil {
		rootRouter.Handle("/metrics", s.app.GetMetricsHandler())
	}

	logger.Info("Starting server on", "addr", s.cfg.GetHTTP().GetAddr())

	s.server = &http.Server{
		Addr:           s.cfg.GetHTTP().GetAddr(),
		Handler:        http.TimeoutHandler(rootRouter.GetHandler(), s.cfg.GetHTTP().GetHandlerTimeout(), "Timeout"),
		ReadTimeout:    s.cfg.GetHTTP().GetReadTimeout(),
		WriteTimeout:   s.cfg.GetHTTP().GetWriteTimeout(),
		MaxHeaderBytes: maxHeaderBytes,
	}

	if s.cfg.GetHTTP().GetTLS().GetEnabled() {
		setTLS(s.server, s.cfg.GetHTTP().GetTLS())
	}

	return nil
}

func (s *BalanceHTTPServer) Run(ctx context.Context) error {

	if s.app == nil {
		return errors.New("Init() must be called before Run()")
	}

	s.app.RunRevenueExporterCleanup(ctx)
	go s.runHTTPServer()
	return nil
}

func (s *BalanceHTTPServer) runHTTPServer() {
	logger := s.app.GetLogger()

	var err error

	if s.cfg.GetHTTP().GetTLS().GetEnabled() {
		logger.Info(
			"Starting server with TLS",
			"cert_file",
			s.cfg.GetHTTP().GetTLS().GetCertFile(),
			"key_file",
			s.cfg.GetHTTP().GetTLS().GetKeyFile(),
		)

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

}

func (s *BalanceHTTPServer) Shutdown(ctx context.Context) error {

	logger := s.app.GetLogger()

	logger.Info("Shutting down server")

	shutdownCtx, shutdown := context.WithTimeout(ctx, shutdownTimeout)
	defer shutdown()

	err := s.server.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error("Failed to stop server", "error", err)
	}

	logger.Info("Server stopped")

	return nil
}

func setTLS(_ *http.Server, _ *config.ConfigTLS) {
	// TODO: implement
	// server.TLSConfig = &tls.Config{
	// 	Certificates: []tls.Certificate{cfg.GetCertificates()},
	// }
}
