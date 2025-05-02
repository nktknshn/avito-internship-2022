package auth_signin

import domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"

type In struct {
	username domainAuth.AuthUserUsername
	password domainAuth.AuthUserPassword
}

func NewInFromValues(username string, password string) (In, error) {
	_username, err := domainAuth.NewAuthUserUsername(username)
	if err != nil {
		return In{}, err
	}

	_password, err := domainAuth.NewAuthUserPassword(password)
	if err != nil {
		return In{}, err
	}

	return In{
		username: _username,
		password: _password,
	}, nil
}
