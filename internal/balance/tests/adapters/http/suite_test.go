package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/mocked"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	ergo "github.com/nktknshn/go-ergo-handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type testCase struct {
	name          string
	payload       map[string]any
	url           string
	routeParams   map[string]string
	useCaseReturn []any
	expectCode    int
	expectErr     string
	expectBody    map[string]any
	auth          bool
	authRole      domainAuth.AuthUserRole
}

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

func (s *HttpTestSuite) SetupSubTest() {
	s.SetupTest()
}

func (s *HttpTestSuite) runTestCases(useCase func() *mock.Mock, handler func() adaptersHttp.Handler, testCases []testCase) {
	for _, tc := range testCases {
		s.Run(tc.name, func() {

			if len(tc.routeParams) > 0 {
				s.setRouteParams(tc.routeParams)
			} else {
				s.setRouteParams(map[string]string{})
			}

			if tc.auth && !tc.authRole.IsEmpty() {
				s.setupAuthRole(tc.authRole)
			} else if tc.auth {
				s.setupAuthAdmin()
			}

			useCaseReturn := []any{nil}

			if tc.useCaseReturn != nil {
				useCaseReturn = tc.useCaseReturn
			}

			useCase().On("Handle", mock.Anything, mock.Anything).Return(useCaseReturn...)

			var resp *httptest.ResponseRecorder

			if tc.auth {
				_, resp = s.requestAuthPayload(handler(), tc.payload, tc.url)
			} else {
				_, resp = s.requestPayload(handler(), tc.payload, tc.url)
			}

			// fmt.Println(resp.Body.String())
			s.Require().Equal(tc.expectCode, resp.Code)

			if tc.expectErr != "" {
				s.Require().JSONEq(ejson(tc.expectErr), resp.Body.String(), "expectedErr "+tc.expectErr+" but got "+resp.Body.String())
			} else if tc.expectBody == nil {
				s.Require().JSONEq(rjsonStr(`{}`), resp.Body.String(), "expectedBody {} but got "+resp.Body.String())
			} else {
				s.Require().JSONEq(rjson(tc.expectBody), resp.Body.String(), "expectedBody "+rjson(tc.expectBody)+" but got "+resp.Body.String())
			}
		})
	}
}

func (s *HttpTestSuite) requestAuthPayload(h adaptersHttp.Handler, payload map[string]any, url string) (*http.Request, *httptest.ResponseRecorder) {
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+fixtures.AuthToken)
	resp := httptest.NewRecorder()
	h.GetHandler().ServeHTTP(resp, req)
	return req, resp
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

func (s *HttpTestSuite) requestPayload(h adaptersHttp.Handler, payload any, url string) (*http.Request, *httptest.ResponseRecorder) {
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("", url, bytes.NewBuffer(jsonPayload))
	resp := httptest.NewRecorder()
	h.GetHandler().ServeHTTP(resp, req)
	return req, resp
}

func (s *HttpTestSuite) setRouteParams(params map[string]string) {
	// TODO: исправить на конструктор для NewHttpAdapter
	ergo.DefaultVarsGetter = NewMuxVarsGetterMock(params)
}

// setupAuthAdmin авторизация вернет роль администратора
func (s *HttpTestSuite) setupAuthAdmin() {
	s.app.SetupAuth(fixtures.AuthToken, auth_validate_token.Out{
		UserID: fixtures.AuthUserID,
		Role:   domainAuth.AuthUserRoleAdmin,
	}, nil)
}

// setupAuthRole авторизация вернет указанную роль
func (s *HttpTestSuite) setupAuthRole(role domainAuth.AuthUserRole) {
	s.app.SetupAuth(fixtures.AuthToken, auth_validate_token.Out{
		UserID: fixtures.AuthUserID,
		Role:   role,
	}, nil)
}
