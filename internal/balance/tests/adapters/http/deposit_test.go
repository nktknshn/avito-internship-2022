package http_test

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestDeposit_Success() {
	s.setupAuthOK()

	s.app.DepositUseCaseMock.On("Handle", mock.Anything, fixtures.InDeposit).Return(nil)
	_, resp := s.requestAuthPayload(s.httpAdapter.Deposit, map[string]any{
		"user_id": fixtures.UserID_i64,
		"amount":  fixtures.Amount100_i64,
		"source":  fixtures.DepositSource_str,
	})
	s.Require().Equal(http.StatusOK, resp.Code)
	s.Require().Equal(rjson(`{}`), resp.Body.String())
}
