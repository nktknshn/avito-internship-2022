package http_test

import (
	"errors"
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HTTPTestSuite) TestGetBalance() {

	var routeParams = map[string]string{"user_id": fixtures.UserID_str}

	testCases := []testCase{
		{
			name:       "success",
			expectCode: http.StatusOK,
			auth:       true,
			useCaseReturn: returnSuccess2(get_balance.Out{
				Available: fixtures.Amount100,
				Reserved:  fixtures.Amount0,
			}),
			routeParams: routeParams,
			expectBody:  map[string]any{"available": 100, "reserved": 0},
		},
		{
			name:          "not found",
			auth:          true,
			expectCode:    http.StatusNotFound,
			useCaseReturn: returnError2[get_balance.Out](domainAccount.ErrAccountNotFound),
			expectErr:     domainAccount.ErrAccountNotFound.Error(),
			routeParams:   routeParams,
		},
		{
			name:          "invalid user id",
			auth:          true,
			expectCode:    http.StatusBadRequest,
			expectErr:     "invalid int64 value: invalid_user_id",
			useCaseReturn: returnError2[get_balance.Out](domain.ErrInvalidUserID),
			routeParams:   map[string]string{"user_id": "invalid_user_id"},
		},
		{
			name:          "not found",
			auth:          true,
			expectCode:    http.StatusNotFound,
			useCaseReturn: returnError2[get_balance.Out](domainAccount.ErrAccountNotFound),
			expectErr:     domainAccount.ErrAccountNotFound.Error(),
			routeParams:   map[string]string{"user_id": fixtures.UserID_str},
		},
		{
			name:       "user is not allowed",
			expectCode: http.StatusForbidden,
			expectErr:  handlers_auth.ErrUserNotAllowed.Error(),
			auth:       true,
			authRole:   domainAuth.AuthUserRoleNobody,
		},
		{
			name:          "use case internal server error",
			auth:          true,
			expectCode:    http.StatusInternalServerError,
			useCaseReturn: returnError2[get_balance.Out](errors.New("internal server error")),
			expectErr:     "internal server error",
			routeParams:   routeParams,
		},
	}
	s.runTestCases(func() *mock.Mock {
		return &s.app.GetBalanceUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.GetBalance
	}, testCases)
}
