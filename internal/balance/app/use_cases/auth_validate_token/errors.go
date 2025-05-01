package auth_validate_token

import "errors"

var (
	ErrEmptyToken       = errors.New("token is empty")
	ErrInvalidToken     = errors.New("invalid token")
	ErrTokenExpired     = errors.New("token expired")
	ErrTokenBlacklisted = errors.New("token blacklisted")
)
