package http_test

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestGetBalance() {
	// testCases := []testCaseQuery{
	// 	{
	// 		name:        "success",
	// 		query:       map[string]string{"user_id": fixtures.UserID_str},
	// 		expectCode:  http.StatusOK,
	// 		expectErr:   "",
	// 		auth:        true,
	// 		authRole:    domainAuth.AuthUserRoleAdmin,
	// 		routeParams: map[string]string{"user_id": fixtures.UserID_str},
	// 	},
	// }
}

func (s *HttpTestSuite) TestGetBalance_Success() {
	s.setupAuthAdmin()

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{
		Available: 100,
		Reserved:  0,
	}, nil)

	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})

	_, resp := s.requestAuth(s.httpAdapter.GetBalance)

	s.Require().Equal(http.StatusOK, resp.Code)
	s.Require().Equal(rjsonStr(`{"available":100,"reserved":0}`), resp.Body.String())
}

func (s *HttpTestSuite) TestGetBalance_NotFound() {
	s.setupAuthAdmin()

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{}, domainAccount.ErrAccountNotFound)

	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)

	s.Require().Equal(http.StatusNotFound, resp.Code)
}

func (s *HttpTestSuite) TestGetBalance_InvalidUserID() {
	s.setupAuthAdmin()
	s.setRouteParams(map[string]string{"user_id": "invalid_user_id"})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)
	s.Require().Equal(http.StatusBadRequest, resp.Code)
}
