package token_generator

import (
	"context"
	"errors"
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenExpired  = errors.New("token expired")
	ErrInvalidClaims = errors.New("invalid claims")
)

type TokenGenerator[T any] interface {
	GenerateToken(ctx context.Context, claims T) (string, error)
}

type TokenValidator[T any] interface {
	ValidateToken(ctx context.Context, token string) (*T, error)
}
