package http_test

import (
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestReserve() {

	var validPayload = map[string]any{
		"user_id":       fixtures.UserID_i64,
		"product_id":    fixtures.ProductID_i64,
		"product_title": fixtures.ProductTitle_str,
		"order_id":      fixtures.OrderID_i64,
		"amount":        fixtures.Amount100_i64,
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
			useCaseReturn: returnError(domainAccount.ErrAccountNotFound),
			expectCode:    http.StatusNotFound,
			expectErr:     domainAccount.ErrAccountNotFound.Error(),
		},
		{
			name:       "user is not allowed",
			auth:       true,
			authRole:   domainAuth.AuthUserRoleNobody,
			expectCode: http.StatusForbidden,
			expectErr:  handlers_auth.ErrUserNotAllowed.Error(),
		},
		{
			name:       "invalid user_id",
			auth:       true,
			payload:    map[string]any{"user_id": -1},
			expectCode: http.StatusBadRequest,
			expectErr:  domain.ErrInvalidUserID.Error(),
		},
		{
			name:       "invalid product_id",
			auth:       true,
			payload:    map[string]any{"user_id": fixtures.UserID_i64, "product_id": -1},
			expectCode: http.StatusBadRequest,
			expectErr:  domainProduct.ErrInvalidProductID.Error(),
		},
		// order_id
		{
			name:       "invalid order_id",
			auth:       true,
			payload:    map[string]any{"user_id": fixtures.UserID_i64, "product_id": fixtures.ProductID_i64, "order_id": -1},
			expectCode: http.StatusBadRequest,
			expectErr:  domain.ErrInvalidOrderID.Error(),
		},
		// amount
		{
			name:       "invalid amount",
			auth:       true,
			payload:    map[string]any{"user_id": fixtures.UserID_i64, "product_id": fixtures.ProductID_i64, "order_id": fixtures.OrderID_i64, "amount": -1},
			expectCode: http.StatusBadRequest,
			expectErr:  domainAmount.ErrInvalidPositiveAmount.Error(),
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.ReserveUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.Reserve
	}, testCases)
}
