package token_generator_jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenGeneratorJWT[T any] struct {
	signKey  []byte
	tokenTTL time.Duration
}

type TokenClaims[T any] struct {
	jwt.StandardClaims
	Data T
}

func NewTokenGeneratorJWT[T any](signKey []byte, tokenTTL time.Duration) *TokenGeneratorJWT[T] {
	return &TokenGeneratorJWT[T]{
		signKey:  signKey,
		tokenTTL: tokenTTL,
	}
}

func (t *TokenGeneratorJWT[T]) GenerateToken(ctx context.Context, claims T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims[T]{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(t.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Data: claims,
	})
	return token.SignedString(t.signKey)
}
