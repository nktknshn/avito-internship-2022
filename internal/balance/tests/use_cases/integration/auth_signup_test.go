package use_cases_test

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

var ()

func (s *AuthSuiteIntegrationTest) TestAuthSignup_Success() {
	err := s.signup.Handle(context.Background(), fixtures.InAuthSignup)
	s.Require().NoError(err)

	result, err := s.ExecSQL(
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
