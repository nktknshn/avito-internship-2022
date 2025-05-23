package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http/server"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/e2e"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func TestE2ETestSuite(t *testing.T) {
	s := &E2ETestSuite{}
	s.NeedsPostgres = true
	s.MigrationsDir = "../../../migrations/postgres"
	suite.Run(t, s)
}

type E2ETestSuite struct {
	testing_pg.TestSuitePg
	server  *server.BalanceHTTPServer
	cleanup func()
}

func (s *E2ETestSuite) SetupSubTest() {

}

func (s *E2ETestSuite) TearDownSubTest() {

}

func (s *E2ETestSuite) SetupTest() {

	cfg, err := config.LoadConfigFromReader(strings.NewReader(e2e.Cfg))
	s.Require().NoError(err)

	port, err := s.DT.GetRunningPort()
	s.Require().NoError(err)

	cfg.Postgres.Addr = "localhost:" + port
	cfg.Postgres.User = "postgres"
	cfg.Postgres.Database = "postgres"
	cfg.Postgres.Schema = "public"
	cfg.Postgres.Password = s.DT.GetDockerConfig().DockerPassword
	cfg.Postgres.MigrationsDir = ""

	app, cleanup, err := app_impl.NewApplication(context.Background(), cfg)
	s.Require().NoError(err)

	s.cleanup = cleanup
	s.server = server.NewHTTPServer(cfg, app)
	s.Require().NoError(s.server.Init(context.Background()))

}

func (s *E2ETestSuite) TearDownTest() {
	s.Require().NoError(s.server.Shutdown(context.Background()))
	s.cleanup()
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *E2ETestSuite) TestHttpServer() {
	s.Run("should expose metrics", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/metrics", nil)
		s.server.GetHandler().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusOK, recorder.Code)
	})

	s.Run("should expose swagger", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
		s.server.GetHandler().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusOK, recorder.Code)
	})

	s.Run("should expose api", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/balance/1", nil)
		s.server.GetHandler().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusUnauthorized, recorder.Code)
	})

	s.Run("should handle OPTIONS", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodOptions, "/api/v1/balance/1", nil)
		request.Header.Set("Access-Control-Request-Method", "POST")
		s.server.GetHandler().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusOK, recorder.Code)
	})

	s.Run("should handle ping", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/ping", nil)
		s.server.GetHandler().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusOK, recorder.Code)
	})
}
