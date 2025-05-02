package use_cases_test

import (
	"context"
	"errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
	"github.com/stretchr/testify/mock"
)

var (
	nilClaims = (*domainAuth.AuthUserTokenClaims)(nil)
)

func (s *AuthSuiteTest) TestAuthValidateToken_InvalidToken() {
	in := must.Must(auth_validate_token.NewInFromValues(
		"123",
	))
	s.createAuthUser()
	out, err := s.validate.Handle(context.Background(), in)
	s.Require().Error(err)
	s.Require().ErrorIs(err, auth_validate_token.ErrInvalidToken)
	s.Require().Empty(out)
}

func (s *AuthSuiteTest) TestAuthValidateToken_Success() {
	in := must.Must(auth_validate_token.NewInFromValues("123"))

	s.mockedTokenVal.On("ValidateToken", mock.Anything, mock.Anything).Return(&domainAuth.AuthUserTokenClaims{
		AuthUserID:   1,
		AuthUserRole: fixtures.AuthUserRole_str,
	}, nil)

	out, err := s.mockedValidate.Handle(context.Background(), in)

	s.Require().NoError(err)
	s.Require().Equal(domainAuth.AuthUserID(1), out.UserID)
	s.Require().Equal(domainAuth.AuthUserRole(fixtures.AuthUserRole_str), out.Role)
}

func (s *AuthSuiteTest) TestAuthValidateToken_InvalidClaims() {
	in := must.Must(auth_validate_token.NewInFromValues("123"))

	s.mockedTokenVal.On("ValidateToken", mock.Anything, mock.Anything).Return(&domainAuth.AuthUserTokenClaims{}, nil)

	out, err := s.mockedValidate.Handle(context.Background(), in)

	s.Require().Empty(out)
	s.Require().ErrorIs(err, auth_validate_token.ErrInvalidClaims)
}

func (s *AuthSuiteTest) TestAuthValidateToken_TokenExpired() {
	in := must.Must(auth_validate_token.NewInFromValues("123"))

	s.mockedTokenVal.On("ValidateToken", mock.Anything, mock.Anything).Return(nilClaims, token_generator.ErrTokenExpired)

	out, err := s.mockedValidate.Handle(context.Background(), in)

	s.Require().ErrorIs(err, auth_validate_token.ErrTokenExpired)
	s.Require().Empty(out)
}

func (s *AuthSuiteTest) TestAuthValidateToken_TokenValidatorError() {
	in := must.Must(auth_validate_token.NewInFromValues("123"))

	s.mockedTokenVal.On("ValidateToken", mock.Anything, mock.Anything).Return(nilClaims, errors.New("token validator error"))

	out, err := s.mockedValidate.Handle(context.Background(), in)

	s.Require().ErrorContains(err, "token validator error")
	s.Require().Empty(out)
}
