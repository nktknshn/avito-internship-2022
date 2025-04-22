package auth_signup

import domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"

type In struct {
	Username domainAuth.AuthUserUsername
	Password domainAuth.AuthUserPassword
	Role     domainAuth.AuthUserRole
}

func NewInFromValues(username string, password string, role string) (In, error) {
	_username, err := domainAuth.NewAuthUserUsername(username)
	if err != nil {
		return In{}, err
	}
	_password, err := domainAuth.NewAuthUserPassword(password)
	if err != nil {
		return In{}, err
	}
	_role, err := domainAuth.NewAuthUserRole(role)
	if err != nil {
		return In{}, err
	}
	return In{
		Username: _username,
		Password: _password,
		Role:     _role,
	}, nil
}
