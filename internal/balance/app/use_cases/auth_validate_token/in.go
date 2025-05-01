package auth_validate_token

import (
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

type In struct {
	token string
}

func NewInFromValues(token string) (In, error) {
	if token == "" {
		return In{}, ErrEmptyToken
	}
	return In{
		token: token,
	}, nil
}

type Out struct {
	UserID domainAuth.AuthUserID
	Role   domainAuth.AuthUserRole
}
