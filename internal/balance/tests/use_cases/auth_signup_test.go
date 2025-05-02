package use_cases_test

import (
	"context"
	"errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
	"github.com/stretchr/testify/mock"
)

var (
	inAuthSignup = must.Must(auth_signup.NewInFromValues(
		fixtures.UsernameAdmin_str,
		fixtures.PasswordAdmin_str,
		fixtures.AuthUserRole_str,
	))
)

func (s *AuthSuiteTest) TestAuthSignup_Success() {
	err := s.signup.Handle(context.Background(), inAuthSignup)
	s.Require().NoError(err)

	result, err := s.ExecSql(
		`SELECT * FROM auth_users`,
	)
	s.Require().NoError(err)
	s.Require().Len(result.Rows, 1)

	row := result.Rows[0]

	s.Require().Equal(fixtures.UsernameAdmin_str, row["username"])
	s.Require().Equal(fixtures.AuthUserRole_str, row["role"])
	s.Require().NotEmpty(row["password_hash"])

	ok, err := s.hasher.Verify(fixtures.PasswordAdmin_str, row["password_hash"].(string))
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s *AuthSuiteTest) TestAuthSignup_HasherError() {
	s.mockedHasher.On("Hash", mock.Anything).Return("", errors.New("hasher error"))

	err := s.mockedSignup.Handle(context.Background(), inAuthSignup)
	s.Require().Error(err)
	s.Require().ErrorContains(err, "hasher error")
}

func (s *AuthSuiteTest) TestAuthSignup_RepositoryError() {
	s.mockedHasher.On("Hash", mock.Anything).Return(fixtures.PasswordHashAdmin_str, nil)
	s.mockedAuthRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("repository error"))

	err := s.mockedSignup.Handle(context.Background(), inAuthSignup)
	s.Require().Error(err)
	s.Require().ErrorContains(err, "repository error")
}
