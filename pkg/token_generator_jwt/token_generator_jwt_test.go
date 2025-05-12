package token_generator_jwt_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
	"github.com/nktknshn/avito-internship-2022/pkg/token_generator_jwt"
)

var (
	secret     = []byte("secret")
	tokenTTL   = time.Hour * 24
	expired    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDYxNzIwNTQsImlhdCI6MTc0NjE3MjA1MywiRGF0YSI6eyJBdXRoVXNlcklEIjoxLCJBdXRoVXNlclJvbGUiOiJhZG1pbiJ9fQ.SLwygvWQPzmVwv2jl9GdDQynoYD5tKghRUFOn0ZcIvE"
	validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEwOTY5NTQ0NDU3LCJpYXQiOjE3NDYxNzI0MjAsIkRhdGEiOnsiQXV0aFVzZXJJRCI6MSwiQXV0aFVzZXJSb2xlIjoiYWRtaW4ifX0.167DZUKxfxj2sRcGAvz_SsAS-_vyV4plR2NcEKaXoqU"
)

func TestTokenGeneratorJWT_GenerateToken(t *testing.T) {
	tokenGen := token_generator_jwt.NewTokenGeneratorJWT[auth.AuthUserTokenClaims](secret, tokenTTL)
	token, err := tokenGen.GenerateToken(t.Context(), auth.AuthUserTokenClaims{
		AuthUserID:   1,
		AuthUserRole: "admin",
	})
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestTokenGeneratorJWT_GenerateToken_Valid(t *testing.T) {
	tokenVal := token_generator_jwt.NewTokenValidatorJWT[auth.AuthUserTokenClaims](secret)
	claims, err := tokenVal.ValidateToken(t.Context(), validToken)
	require.NoError(t, err)
	require.Equal(t, &auth.AuthUserTokenClaims{
		AuthUserID:   1,
		AuthUserRole: "admin",
	}, claims)
}

func TestTokenValidatorJWT_ExpiredToken(t *testing.T) {
	tokenVal := token_generator_jwt.NewTokenValidatorJWT[auth.AuthUserTokenClaims](secret)
	claims, err := tokenVal.ValidateToken(t.Context(), expired)
	require.Error(t, err)
	require.Nil(t, claims)
	require.ErrorIs(t, err, token_generator.ErrTokenExpired)
}

func TestTokenValidatorJWT_InvalidToken(t *testing.T) {
	tokenVal := token_generator_jwt.NewTokenValidatorJWT[auth.AuthUserTokenClaims](secret)
	claims, err := tokenVal.ValidateToken(t.Context(), "invalid")
	require.Error(t, err)
	require.Nil(t, claims)
	require.ErrorIs(t, err, token_generator.ErrInvalidToken)
}
