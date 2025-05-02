package use_cases_test

import (
	"context"

	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/password_hasher"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
	"github.com/stretchr/testify/mock"
)

type authRepoMock struct {
	mock.Mock
}

func (m *authRepoMock) GetUserByUsername(ctx context.Context, username domainAuth.AuthUserUsername) (*domainAuth.AuthUser, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*domainAuth.AuthUser), args.Error(1)
}

func (m *authRepoMock) CreateUser(ctx context.Context, username domainAuth.AuthUserUsername, passwordHash domainAuth.AuthUserPasswordHash, role domainAuth.AuthUserRole) error {
	args := m.Called(ctx, username, passwordHash, role)
	return args.Error(0)
}

type tokenManagerMock struct {
	mock.Mock
}

func (m *tokenManagerMock) GenerateToken(ctx context.Context, claims domainAuth.AuthUserTokenClaims) (string, error) {
	args := m.Called(ctx, claims)
	return args.String(0), args.Error(1)
}

func (m *tokenManagerMock) ValidateToken(ctx context.Context, token string) (*domainAuth.AuthUserTokenClaims, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(*domainAuth.AuthUserTokenClaims), args.Error(1)
}

type hasherVerifierMock struct {
	mock.Mock
}

func (m *hasherVerifierMock) Verify(password, hash string) (bool, error) {
	args := m.Called(password, hash)
	return args.Bool(0), args.Error(1)
}

func (m *hasherVerifierMock) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

var _ password_hasher.HashVerifier = &hasherVerifierMock{}
var _ token_generator.TokenGenerator[domainAuth.AuthUserTokenClaims] = &tokenManagerMock{}
var _ token_generator.TokenValidator[domainAuth.AuthUserTokenClaims] = &tokenManagerMock{}
