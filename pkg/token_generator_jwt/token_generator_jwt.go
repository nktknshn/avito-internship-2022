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
	if signKey == nil {
		panic("signKey is nil")
	}

	if tokenTTL <= 0 {
		panic("tokenTTL is not positive")
	}

	return &TokenGeneratorJWT[T]{
		signKey:  signKey,
		tokenTTL: tokenTTL,
	}
}

func (t *TokenGeneratorJWT[T]) GenerateToken(_ context.Context, claims T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims[T]{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			Audience:  []string{},
			NotBefore: nil,
			ID:        "",
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

	if signKey == nil {
		panic("signKey is nil")
	}

	return &TokenValidatorJWT[T]{
		signKey: signKey,
	}
}

func (t *TokenValidatorJWT[T]) ValidateToken(_ context.Context, token string) (*T, error) {
	var claims TokenClaims[T]

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(_ *jwt.Token) (interface{}, error) {
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
		return &claims.Data, nil
	}

	return nil, token_generator.ErrInvalidToken
}

type TokenGeneratorJWTVerifier[T any] struct {
	TokenGenerator token_generator.TokenGenerator[T]
	TokenValidator token_generator.TokenValidator[T]
}

func NewTokenGeneratorJWTVerifier[T any](signKey []byte, tokenTTL time.Duration) *TokenGeneratorJWTVerifier[T] {
	return &TokenGeneratorJWTVerifier[T]{
		TokenGenerator: NewTokenGeneratorJWT[T](signKey, tokenTTL),
		TokenValidator: NewTokenValidatorJWT[T](signKey),
	}
}

func (t *TokenGeneratorJWTVerifier[T]) GenerateToken(ctx context.Context, claims T) (string, error) {
	return t.TokenGenerator.GenerateToken(ctx, claims)
}

func (t *TokenGeneratorJWTVerifier[T]) ValidateToken(ctx context.Context, token string) (*T, error) {
	return t.TokenValidator.ValidateToken(ctx, token)
}
