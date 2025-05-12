package use_cases_test

import (
	"context"

	"github.com/stretchr/testify/mock"

	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/password_hasher"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
)

type tokenManagerMock struct {
	mock.Mock
}

func (m *tokenManagerMock) GenerateToken(ctx context.Context, claims domainAuth.AuthUserTokenClaims) (string, error) {
	args := m.Called(ctx, claims)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.String(0), args.Error(1)
}

func (m *tokenManagerMock) ValidateToken(ctx context.Context, token string) (*domainAuth.AuthUserTokenClaims, error) {
	args := m.Called(ctx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainAuth.AuthUserTokenClaims), args.Error(1)
}

type hasherVerifierMock struct {
	mock.Mock
}

func (m *hasherVerifierMock) Verify(password, hash string) (bool, error) {
	args := m.Called(password, hash)
	if args.Get(0) == nil {
		return false, args.Error(1)
	}
	return args.Bool(0), args.Error(1)
}

func (m *hasherVerifierMock) Hash(password string) (string, error) {
	args := m.Called(password)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.String(0), args.Error(1)
}

var _ password_hasher.HashVerifier = &hasherVerifierMock{}
var _ token_generator.TokenGenerator[domainAuth.AuthUserTokenClaims] = &tokenManagerMock{}
var _ token_generator.TokenValidator[domainAuth.AuthUserTokenClaims] = &tokenManagerMock{}
