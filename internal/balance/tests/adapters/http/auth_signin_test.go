package http_test

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

func (s *HTTPTestSuite) TestAuthSignin() {
	testCases := []testCase{
		{
			name: "success",
			payload: map[string]any{
				"username": "test",
				"password": "test12345",
			},
			expectCode: http.StatusOK,
			useCaseReturn: returnSuccess2(auth_signin.Out{
				Token: "test",
			}),
			expectBody: map[string]any{
				"token": "test",
			},
		},
		{
			name: "invalid username",
			payload: map[string]any{
				"username": "test",
				"password": "test12345",
			},
			useCaseReturn: returnError2[auth_signin.Out](domainAuth.ErrInvalidAuthUserPassword),
			expectCode:    http.StatusUnauthorized,
			expectErr:     domainAuth.ErrInvalidAuthUserPassword.Error(),
		},
		{
			name: "invalid username value",
			payload: map[string]any{
				"username": "",
				"password": "",
			},
			expectCode: http.StatusBadRequest,
			expectErr:  domainAuth.ErrInvalidAuthUserUsernameTooShort.Error(),
		},
		{
			name: "invalid password value",
			payload: map[string]any{
				"username": "test",
				"password": "",
			},
			expectCode: http.StatusBadRequest,
			expectErr:  domainAuth.ErrInvalidAuthUserPasswordTooShort.Error(),
		},
		// internal error
		{
			name:          "internal error",
			useCaseReturn: returnError2[auth_signin.Out](errors.New("do not expose")),
			expectCode:    http.StatusInternalServerError,
			expectErr:     "internal server error",
			auth:          true,
			payload: map[string]any{
				"username": "test",
				"password": "test12345",
			},
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.AuthSigninUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.AuthSignIn
	}, testCases)
}
