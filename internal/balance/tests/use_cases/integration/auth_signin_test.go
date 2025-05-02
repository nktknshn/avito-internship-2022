package use_cases_test

import (
	"context"

	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
)

func (s *AuthSuiteIntegrationTest) TearDownTest() {
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *AuthSuiteIntegrationTest) TestAuthSignin_Success() {
	s.createAuthUser()
	out, err := s.signin.Handle(context.Background(), fixtures.InAuthSignin)
	s.Require().NoError(err)
	s.Require().NotEmpty(out.Token)
}

func (s *AuthSuiteIntegrationTest) TestAuthSignin_NotFound() {
	out, err := s.signin.Handle(context.Background(), fixtures.InAuthSignin)
	s.Require().Error(err)
	s.Require().Empty(out)
	s.Require().ErrorIs(err, domainAuth.ErrAuthUserNotFound)
}

func (s *AuthSuiteIntegrationTest) TestAuthSignin_InvalidPassword() {
	s.createAuthUser()
	out, err := s.signin.Handle(context.Background(), fixtures.InAuthSigninInvalidPassword)
	s.Require().Error(err)
	s.Require().Empty(out)
	s.Require().ErrorIs(err, domainAuth.ErrInvalidAuthUserPassword)
}
