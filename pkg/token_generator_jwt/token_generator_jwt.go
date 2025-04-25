package token_generator_jwt

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
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
	jwt.StandardClaims
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
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(t.tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
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

	var validationError *jwt.ValidationError
	if errors.As(err, &validationError) {
		if validationError.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, token_generator.ErrTokenExpired
		}
	}

	if err != nil {
		return nil, token_generator.ErrInvalidToken
	}

	if claims, ok := parsedToken.Claims.(*TokenClaims[T]); ok && parsedToken.Valid {
		return &claims.Data, nil
	}

	return nil, token_generator.ErrInvalidClaims
}
