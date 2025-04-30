package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/mocked"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	ergo "github.com/nktknshn/go-ergo-handler"
	"github.com/stretchr/testify/suite"
)

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}

type HttpTestSuite struct {
	testing_pg.TestSuitePg
	app         *mocked.AppMocked
	httpAdapter *adaptersHttp.HttpAdapter
}

func (s *HttpTestSuite) SetupTest() {
	s.app = mocked.NewMockedApp()
	s.httpAdapter = adaptersHttp.NewHttpAdapter(&s.app.Application)
	s.setRouteParams(map[string]string{})
}

func (s *HttpTestSuite) requestAuth(h adaptersHttp.Handler) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("", "", nil)
	req.Header.Set("Authorization", "Bearer "+fixtures.AuthToken)
	resp := httptest.NewRecorder()
	h.GetHandler().ServeHTTP(resp, req)
	return req, resp
}

func (s *HttpTestSuite) request(h adaptersHttp.Handler) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("", "", nil)
	resp := httptest.NewRecorder()
	h.GetHandler().ServeHTTP(resp, req)
	return req, resp
}

func (s *HttpTestSuite) setRouteParams(params map[string]string) {
	// TODO: исправить на конструктор для NewHttpAdapter
	ergo.DefaultVarsGetter = NewMuxVarsGetterMock(params)
}
