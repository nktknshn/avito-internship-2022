package auth

import domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"

type AuthUserRole string

func (r AuthUserRole) Value() string {
	return string(r)
}

func (r AuthUserRole) IsEmpty() bool {
	return r == ""
}

var (
	ErrInvalidAuthUserRole = domainError.New("invalid auth user role")
)

const (
	AuthUserRoleAdmin   AuthUserRole = "admin"
	AuthUserRoleAccount AuthUserRole = "account"
	AuthUserRoleReport  AuthUserRole = "report"
	AuthUserRoleNobody  AuthUserRole = "nobody"
)

func (a AuthUserRole) Validate() error {
	switch a {
	case AuthUserRoleAdmin, AuthUserRoleAccount, AuthUserRoleReport, AuthUserRoleNobody:
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
	case AuthUserRoleNobody.Value():
		return AuthUserRoleNobody, nil
	}

	return "", ErrInvalidAuthUserRole
}
