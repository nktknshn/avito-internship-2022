package domain

import "errors"

type AuthUserRole string

var (
	ErrInvalidAuthUserRole = errors.New("invalid auth user role")
)

const (
	AuthUserRoleAdmin   = "admin"
	AuthUserRoleAccount = "account"
	AuthUserRoleReport  = "report"
)

func (a AuthUserRole) Validate() error {
	switch a {
	case AuthUserRoleAdmin, AuthUserRoleAccount, AuthUserRoleReport:
		return nil
	}

	return ErrInvalidAuthUserRole
}

func NewAuthUserRole(s string) (AuthUserRole, error) {
	switch s {
	case AuthUserRoleAdmin:
		return AuthUserRoleAdmin, nil
	case AuthUserRoleReport:
		return AuthUserRoleReport, nil
	case AuthUserRoleAccount:
		return AuthUserRoleAccount, nil
	}

	return "", ErrInvalidAuthUserRole
}
