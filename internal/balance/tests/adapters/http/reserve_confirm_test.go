package http_test

import (
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestReserveConfirm() {
	var validPayload = map[string]any{
		"user_id":    fixtures.UserID_i64,
		"product_id": fixtures.ProductID_i64,
		"order_id":   fixtures.OrderID_i64,
		"amount":     fixtures.Amount100_i64,
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
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.ReserveConfirmUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.ReserveConfirm
	}, testCases)
}
