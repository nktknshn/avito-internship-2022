package http_test

import (
	"errors"
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestTransfer() {

	var validPayload = map[string]any{
		"from":   fixtures.AccountID_i64,
		"to":     fixtures.AccountID_2_i64,
		"amount": fixtures.Amount100_i64,
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
			payload:    map[string]any{"from": -1},
			expectCode: http.StatusBadRequest,
			expectErr:  domainAccount.ErrInvalidAccountID.Error(),
		},
		{
			name:       "invalid to",
			auth:       true,
			payload:    map[string]any{"to": -1},
			expectCode: http.StatusBadRequest,
			expectErr:  domainAccount.ErrInvalidAccountID.Error(),
		},
		{
			name:       "invalid amount",
			auth:       true,
			payload:    map[string]any{"amount": -1},
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
