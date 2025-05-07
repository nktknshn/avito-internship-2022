package token_generator

import (
	"context"
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

type TokenGenerator[T any] interface {
	GenerateToken(ctx context.Context, claims T) (string, error)
}

type TokenValidator[T any] interface {
	ValidateToken(ctx context.Context, token string) (*T, error)
}

type TokenGeneratorVerifier[T any] interface {
	TokenGenerator[T]
	TokenValidator[T]
}
