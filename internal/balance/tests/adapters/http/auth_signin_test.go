package http_test

import (
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestAuthSignin() {
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
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.AuthSigninUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.AuthSignIn
	}, testCases)
}
