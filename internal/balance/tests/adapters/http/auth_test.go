package http_test

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	ergo "github.com/nktknshn/go-ergo-handler"
	"github.com/stretchr/testify/mock"
)

func (s *HTTPTestSuite) TestAuth_Success() {
	s.setupAuthAdmin()
	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{}, nil)
	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)
	s.Require().Equal(http.StatusOK, resp.Code)
}

func (s *HTTPTestSuite) TestAuth_MissingToken() {
	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})
	_, resp := s.request(s.httpAdapter.GetBalance)
	s.Require().Equal(http.StatusUnauthorized, resp.Code)
	s.Require().Equal(ejsonStr(ergo.ErrAuthMissingToken.Error()), resp.Body.String())
}

func (s *HTTPTestSuite) TestAuth_InvalidRole() {
	s.setupAuthRole(domainAuth.AuthUserRoleReport)
	_, resp := s.requestAuth(s.httpAdapter.Deposit)
	s.Require().Equal(http.StatusForbidden, resp.Code)
	s.Require().Equal(ejsonStr(handlers_auth.ErrUserNotAllowed.Error()), resp.Body.String())
}
