package auth_validate_token

import "errors"

var (
	ErrEmptyToken       = errors.New("token is empty")
	ErrInvalidToken     = errors.New("invalid token")
	ErrTokenExpired     = errors.New("token expired")
	ErrInvalidClaims    = errors.New("invalid claims")
	ErrTokenBlacklisted = errors.New("token blacklisted")
)
