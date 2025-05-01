package http_test

import (
	"errors"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) Test_DomainError() {
	s.setupAuthAdmin()

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{}, domainAccount.ErrAccountNotFound)

	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)

	// s.Require().Equal(http.StatusNotFound, resp.Code)
	s.Require().Equal(ejsonStr(domainAccount.ErrAccountNotFound.Error()), resp.Body.String())
}

func (s *HttpTestSuite) Test_InternalError() {
	s.setupAuthAdmin()

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{}, errors.New("internal server error that should not be exposed to the client"))

	s.setRouteParams(map[string]string{"user_id": fixtures.UserID_str})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)

	s.Require().Equal(http.StatusInternalServerError, resp.Code)
	s.Require().Equal(ejsonStr("internal server error"), resp.Body.String())
}

func (s *HttpTestSuite) Test_ParserError() {
	s.setupAuthAdmin()

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(get_balance.Out{}, errors.New("parser error"))

	s.setRouteParams(map[string]string{"user_id": "invalid_user_id"})
	_, resp := s.requestAuth(s.httpAdapter.GetBalance)

	s.Require().Equal(http.StatusBadRequest, resp.Code)
	s.Require().Equal(ejsonStr("invalid int64 value: invalid_user_id"), resp.Body.String())
}
