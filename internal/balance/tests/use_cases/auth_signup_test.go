package use_cases_test

import (
	"context"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *AuthSuiteTest) TestAuthSignup_HasherError() {
	s.mockedHasher.On("Hash", mock.Anything).Return("", errors.New("hasher error"))

	err := s.mockedSignup.Handle(context.Background(), fixtures.InAuthSignup)
	s.Require().Error(err)
	s.Require().ErrorContains(err, "hasher error")
}

func (s *AuthSuiteTest) TestAuthSignup_RepositoryError() {
	s.mockedHasher.On("Hash", mock.Anything).Return(fixtures.PasswordHashAdmin_str, nil)

	s.mockedAuthRepo.On(
		"CreateUser",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("repository error"))

	err := s.mockedSignup.Handle(context.Background(), fixtures.InAuthSignup)

	s.Require().Error(err)
	s.Require().ErrorContains(err, "repository error")
}
