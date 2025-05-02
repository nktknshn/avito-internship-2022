package use_cases_test

import (
	"context"

	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

func (s *AuthSuiteTest) TestAuthSignin_HasherError() {
	hasherError := errors.New("hasher error")
	s.mockedAuthRepo.On("GetUserByUsername", mock.Anything, mock.Anything).Return(&fixtures.AuthUser, nil)

	s.mockedHasher.On("Verify", mock.Anything, mock.Anything).Return(false, hasherError)

	out, err := s.mockedSignin.Handle(context.Background(), fixtures.InAuthSignin)
	s.Require().Error(err)
	s.Require().Empty(out)
	s.Require().ErrorIs(err, domainAuth.ErrInvalidAuthUserPassword)

	s.Require().Equal(errors.Cause(err).Error(), hasherError.Error())
}

func (s *AuthSuiteTest) TestAuthSignin_TokenGenError() {
	s.mockedAuthRepo.On("GetUserByUsername", mock.Anything, mock.Anything).Return(&fixtures.AuthUser, nil)

	s.mockedHasher.On("Verify", mock.Anything, mock.Anything).Return(true, nil)

	s.mockedTokenGen.On("GenerateToken", mock.Anything, mock.Anything).Return("", errors.New("token gen error"))

	out, err := s.mockedSignin.Handle(context.Background(), fixtures.InAuthSignin)
	s.Require().Error(err)
	s.Require().Empty(out)
	s.Require().ErrorContains(err, "token gen error")

}
