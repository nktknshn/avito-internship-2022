package auth

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDuplicateKey = errors.New("duplicate key")
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username AuthUserUsername) (*AuthUser, error)
	GetBlacklist(ctx context.Context) ([]AuthUserToken, error)
	CreateUser(ctx context.Context, username AuthUserUsername, passwordHash AuthUserPasswordHash, role AuthUserRole) error
}
