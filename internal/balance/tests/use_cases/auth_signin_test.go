package use_cases_test

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	inAuthSignin = must.Must(auth_signin.NewInFromValues(
		fixtures.UsernameAdmin_str,
		fixtures.PasswordAdmin_str,
	))
	inAuthSigninInvalidPassword = must.Must(auth_signin.NewInFromValues(
		fixtures.UsernameAdmin_str,
		fixtures.PasswordAdmin_str+"invalid",
	))
)

func (s *AuthSuiteTest) TearDownTest() {
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *AuthSuiteTest) TestAuthSignin_Success() {
	s.createAuthUser()
	out, err := s.signin.Handle(context.Background(), inAuthSignin)
	s.Require().NoError(err)
	s.Require().NotEmpty(out.Token)
}

func (s *AuthSuiteTest) TestAuthSignin_InvalidPassword() {
	s.createAuthUser()
	out, err := s.signin.Handle(context.Background(), inAuthSigninInvalidPassword)
	s.Require().Error(err)
	s.Require().Empty(out)
	s.Require().ErrorIs(err, domainAuth.ErrInvalidAuthUserPassword)
}
