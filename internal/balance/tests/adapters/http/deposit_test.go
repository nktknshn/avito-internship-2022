package http_test

import (
	"errors"
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	ergo "github.com/nktknshn/go-ergo-handler"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestDeposit() {

	var validPayload = map[string]any{
		"user_id": fixtures.UserID_i64,
		"amount":  fixtures.Amount100_i64,
		"source":  fixtures.DepositSource_str,
	}

	testCases := []testCase{
		{
			name:       "success",
			payload:    validPayload,
			expectCode: http.StatusOK,
			auth:       true,
		},
		{
			name: "invalid user id",
			payload: map[string]any{
				"user_id": -1,
				"amount":  fixtures.Amount100_i64,
				"source":  fixtures.DepositSource_str,
			},
			expectCode: http.StatusBadRequest,
			expectErr:  domain.ErrInvalidUserID.Error(),
			auth:       true,
		},
		{
			name: "invalid amount",
			payload: map[string]any{
				"user_id": fixtures.UserID_i64,
				"amount":  -1,
				"source":  fixtures.DepositSource_str,
			},
			expectCode: http.StatusBadRequest,
			expectErr:  domainAmount.ErrInvalidPositiveAmount.Error(),
			auth:       true,
		},
		{
			name: "error parsing payload",
			payload: map[string]any{
				"user_id": "invalid_user_id",
				"amount":  "invalid_amount",
				"source":  1,
			},
			expectCode: http.StatusBadRequest,
			expectErr:  `error parsing payload`,
			auth:       true,
		},
		{
			name:       "no auth",
			payload:    validPayload,
			expectCode: http.StatusUnauthorized,
			expectErr:  ergo.ErrAuthMissingToken.Error(),
			auth:       false,
		},
		{
			name:       "user is not allowed",
			payload:    validPayload,
			expectCode: http.StatusForbidden,
			expectErr:  handlers_auth.ErrUserNotAllowed.Error(),
			auth:       true,
			authRole:   domainAuth.AuthUserRoleNobody,
		},
		{
			name:          "use case error",
			payload:       validPayload,
			expectCode:    http.StatusInternalServerError,
			expectErr:     "internal server error",
			auth:          true,
			useCaseReturn: returnError(errors.New(`error`)),
		},
		{
			name:          "use domain error",
			payload:       validPayload,
			expectCode:    http.StatusNotFound,
			expectErr:     domainAccount.ErrAccountNotFound.Error(),
			auth:          true,
			useCaseReturn: returnError(domainAccount.ErrAccountNotFound),
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.DepositUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.Deposit
	}, testCases)
}
