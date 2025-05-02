package auth_validate_token

import (
	useCasesError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
)

var (
	ErrEmptyToken       = useCasesError.New("token is empty")
	ErrInvalidToken     = useCasesError.New("invalid token")
	ErrTokenExpired     = useCasesError.New("token expired")
	ErrInvalidClaims    = useCasesError.New("invalid claims")
	ErrTokenBlacklisted = useCasesError.New("token blacklisted")
)
