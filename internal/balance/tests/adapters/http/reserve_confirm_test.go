package http_test

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *HTTPTestSuite) TestReserveConfirm() {
	var validPayload = map[string]any{
		"user_id":    fixtures.UserID_i64,
		"product_id": fixtures.ProductID_i64,
		"order_id":   fixtures.OrderID_i64,
		"amount":     fixtures.Amount100_i64,
	}

	var invalidPayload = map[string]any{
		"user_id":    fixtures.UserID_i64,
		"product_id": fixtures.ProductID_i64,
	}

	testCases := []testCase{
		{
			name:          "success",
			auth:          true,
			payload:       validPayload,
			expectCode:    http.StatusOK,
			useCaseReturn: returnNilError(),
		},
		{
			name:       "invalid payload",
			auth:       true,
			payload:    invalidPayload,
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid order id",
			// useCaseReturn: returnError(),
		},
		{
			name:          "account not found",
			auth:          true,
			payload:       validPayload,
			expectCode:    http.StatusNotFound,
			useCaseReturn: returnError(domainAccount.ErrAccountNotFound),
			expectErr:     domainAccount.ErrAccountNotFound.Error(),
		},
		{
			name:          "user is not allowed",
			auth:          true,
			authRole:      domainAuth.AuthUserRoleNobody,
			payload:       validPayload,
			expectCode:    http.StatusForbidden,
			expectErr:     handlers_auth.ErrUserNotAllowed.Error(),
			useCaseReturn: returnError(domainAccount.ErrAccountNotFound),
		},
		{
			name:          "transaction is already paid",
			auth:          true,
			payload:       validPayload,
			expectCode:    http.StatusConflict,
			useCaseReturn: returnError(domainTransaction.ErrTransactionAlreadyPaid),
			expectErr:     domainTransaction.ErrTransactionAlreadyPaid.Error(),
		},
		{
			name:          "transaction not found",
			auth:          true,
			payload:       validPayload,
			expectCode:    http.StatusNotFound,
			useCaseReturn: returnError(domainTransaction.ErrTransactionNotFound),
			expectErr:     domainTransaction.ErrTransactionNotFound.Error(),
		},
		{
			name:          "internal error",
			auth:          true,
			payload:       validPayload,
			expectCode:    http.StatusInternalServerError,
			useCaseReturn: returnError(errors.New("some internal error")),
			expectErr:     "internal server error",
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.ReserveConfirmUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.ReserveConfirm
	}, testCases)
}
