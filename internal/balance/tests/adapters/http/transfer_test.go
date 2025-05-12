package http_test

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *HTTPTestSuite) TestTransfer() {

	var validPayload = map[string]any{
		"from_user_id": fixtures.UserID_i64,
		"to_user_id":   fixtures.UserID_2_i64,
		"amount":       fixtures.AmountPositive100_i64,
	}

	testCases := []testCase{
		{
			name:       "success",
			auth:       true,
			payload:    validPayload,
			expectCode: http.StatusOK,
		},
		{
			name:       "invalid from",
			auth:       true,
			payload:    map[string]any{"from_user_id": -1, "to_user_id": fixtures.UserID_2_i64, "amount": fixtures.AmountPositive100_i64},
			expectCode: http.StatusBadRequest,
			expectErr:  domain.ErrInvalidUserID.Error(),
		},
		{
			name:       "invalid to",
			auth:       true,
			payload:    map[string]any{"from_user_id": fixtures.UserID_i64, "to_user_id": -1, "amount": fixtures.AmountPositive100_i64},
			expectCode: http.StatusBadRequest,
			expectErr:  domain.ErrInvalidUserID.Error(),
		},
		{
			name:       "invalid amount",
			auth:       true,
			payload:    map[string]any{"from_user_id": fixtures.UserID_i64, "to_user_id": fixtures.UserID_2_i64, "amount": -1},
			expectCode: http.StatusBadRequest,
			expectErr:  domainAmount.ErrInvalidPositiveAmount.Error(),
		},
		{
			name:       "user is not allowed",
			auth:       true,
			authRole:   domainAuth.AuthUserRoleNobody,
			expectCode: http.StatusForbidden,
			expectErr:  handlers_auth.ErrUserNotAllowed.Error(),
		},
		{
			name:          "use case internal server error",
			auth:          true,
			payload:       validPayload,
			useCaseReturn: returnError(errors.New("internal server error")),
			expectCode:    http.StatusInternalServerError,
			expectErr:     "internal server error",
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.TransferUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.Transfer
	}, testCases)
}
