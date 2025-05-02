package use_cases_test

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var ()

func (s *AuthSuiteIntegrationTest) TestAuthValidateToken_InvalidToken() {
	in := must.Must(auth_validate_token.NewInFromValues(
		"123",
	))
	s.createAuthUser()
	out, err := s.validate.Handle(context.Background(), in)
	s.Require().Error(err)
	s.Require().ErrorIs(err, auth_validate_token.ErrInvalidToken)
	s.Require().Empty(out)
}
