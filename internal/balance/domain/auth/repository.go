package auth

import (
	"context"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrAuthUserNotFound  = domainError.New("user not found")
	ErrDuplicateUsername = domainError.New("duplicate username")
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username AuthUserUsername) (*AuthUser, error)
	CreateUser(ctx context.Context, username AuthUserUsername, passwordHash AuthUserPasswordHash, role AuthUserRole) error
}
