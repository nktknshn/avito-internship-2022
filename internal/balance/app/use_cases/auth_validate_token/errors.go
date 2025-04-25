package auth_validate_token

import "errors"

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrTokenExpired     = errors.New("token expired")
	ErrTokenBlacklisted = errors.New("token blacklisted")
)
