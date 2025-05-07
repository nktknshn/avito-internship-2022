package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http/server"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/suite"
)

func TestE2ETestSuite(t *testing.T) {
	s := &E2ETestSuite{}
	s.NeedsPostgres = true
	s.MigrationsDir = "../../../migrations/postgres"
	suite.Run(t, s)
}

type E2ETestSuite struct {
	testing_pg.TestSuitePg
	server *server.BalanceHttpServer
}

func (s *E2ETestSuite) SetupSubTest() {

}

func (s *E2ETestSuite) TearDownSubTest() {

}

func (s *E2ETestSuite) SetupTest() {

	cfg, err := config.LoadConfigFromReader(strings.NewReader(cfg))
	s.Require().NoError(err)

	port, err := s.DT.GetRunningPort()
	s.Require().NoError(err)
	cfg.Postgres.Addr = "localhost:" + port

	cfg.Postgres.UserName = "postgres"
	cfg.Postgres.Database = "postgres"
	cfg.Postgres.Schema = "public"
	cfg.Postgres.Password = s.DT.GetDockerConfig().DockerPassword
	cfg.Postgres.MigrationsDir = ""

	s.server = server.NewHttpServer(cfg)
	s.Require().NoError(s.server.Init(context.Background()))

}

func (s *E2ETestSuite) TearDownTest() {
	s.Require().NoError(s.server.Shutdown(context.Background()))
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *E2ETestSuite) TestHttpServer() {
	s.Run("should return metrics", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/metrics", nil)
		s.server.GetServeMux().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusOK, recorder.Code)
	})

	s.Run("should return swagger", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
		s.server.GetServeMux().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusOK, recorder.Code)
	})

	s.Run("should return api", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/v1/balance/1", nil)
		s.server.GetServeMux().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusUnauthorized, recorder.Code)
	})

	s.Run("should return OPTIONS", func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodOptions, "/api/v1/balance/1", nil)
		request.Header.Set("Access-Control-Request-Method", "POST")
		s.server.GetServeMux().ServeHTTP(recorder, request)
		s.Require().Equal(http.StatusOK, recorder.Code)
	})
}
