package auth

import "errors"

type AuthUserRole string

func (r AuthUserRole) Value() string {
	return string(r)
}

var (
	ErrInvalidAuthUserRole = errors.New("invalid auth user role")
)

const (
	AuthUserRoleAdmin   AuthUserRole = "admin"
	AuthUserRoleAccount AuthUserRole = "account"
	AuthUserRoleReport  AuthUserRole = "report"
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
	case AuthUserRoleAdmin.Value():
		return AuthUserRoleAdmin, nil
	case AuthUserRoleReport.Value():
		return AuthUserRoleReport, nil
	case AuthUserRoleAccount.Value():
		return AuthUserRoleAccount, nil
	}

	return "", ErrInvalidAuthUserRole
}
