package token_generator_jwt

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
)

type claimsType[T any] struct {
	Data T
}

type TokenGeneratorJWT[T any] struct {
	signKey  []byte
	tokenTTL time.Duration
}

type TokenClaims[T any] struct {
	jwt.RegisteredClaims
	claimsType[T]
}

func NewTokenGeneratorJWT[T any](signKey []byte, tokenTTL time.Duration) *TokenGeneratorJWT[T] {
	return &TokenGeneratorJWT[T]{
		signKey:  signKey,
		tokenTTL: tokenTTL,
	}
}

func (t *TokenGeneratorJWT[T]) GenerateToken(ctx context.Context, claims T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims[T]{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		claimsType: claimsType[T]{
			Data: claims,
		},
	})
	return token.SignedString(t.signKey)
}

type TokenValidatorJWT[T any] struct {
	signKey []byte
}

func NewTokenValidatorJWT[T any](signKey []byte) *TokenValidatorJWT[T] {
	return &TokenValidatorJWT[T]{
		signKey: signKey,
	}
}

func (t *TokenValidatorJWT[T]) ValidateToken(ctx context.Context, token string) (*T, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims[T]{}, func(token *jwt.Token) (interface{}, error) {
		return t.signKey, nil
	})

	if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, token_generator.ErrInvalidToken
	}

	if errors.Is(err, jwt.ErrSignatureInvalid) {
		return nil, token_generator.ErrInvalidToken
	}

	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, token_generator.ErrTokenExpired
	}

	if err != nil {
		return nil, token_generator.ErrInvalidToken
	}

	if parsedToken.Valid {
		return &parsedToken.Claims.(*TokenClaims[T]).Data, nil
	}

	return nil, token_generator.ErrInvalidToken
}
