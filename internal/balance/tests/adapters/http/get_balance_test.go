package http_test

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) setupAuthOK() {
	s.app.SetupAuth(fixtures.AuthToken, auth_validate_token.Out{
		UserID: fixtures.AuthUserID,
		Role:   domainAuth.AuthUserRoleAdmin,
	}, nil)
}

func (s *HttpTestSuite) setupAuthRole(role domainAuth.AuthUserRole) {
	s.app.SetupAuth(fixtures.AuthToken, auth_validate_token.Out{
		UserID: fixtures.AuthUserID,
		Role:   role,
	}, nil)
}

func (s *HttpTestSuite) TestGetBalance_Success() {
	s.setupAuthOK()

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{
		Available: 100,
		Reserved:  0,
	}, nil)

	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})

	_, resp := s.requestAuth(s.httpAdapter.GetBalance)

	s.Require().Equal(http.StatusOK, resp.Code)
	s.Require().Equal(rjson(`{"available":100,"reserved":0}`), resp.Body.String())
}

func (s *HttpTestSuite) TestGetBalance_NotFound() {
	s.setupAuthOK()

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{}, domainAccount.ErrAccountNotFound)

	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)

	s.Require().Equal(http.StatusNotFound, resp.Code)
}

func (s *HttpTestSuite) TestGetBalance_InvalidUserID() {
	s.setupAuthOK()
	s.setRouteParams(map[string]string{"user_id": "invalid_user_id"})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)
	s.Require().Equal(http.StatusBadRequest, resp.Code)
}
