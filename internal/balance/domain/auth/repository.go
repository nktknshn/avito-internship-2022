package auth

import (
	"context"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrUserNotFound = domainError.New("user not found")
	ErrDuplicateKey = domainError.New("duplicate key")
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username AuthUserUsername) (*AuthUser, error)
	GetBlacklist(ctx context.Context) ([]AuthUserToken, error)
	CreateUser(ctx context.Context, username AuthUserUsername, passwordHash AuthUserPasswordHash, role AuthUserRole) error
}
