package auth_pg

import (
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

type authUserDTO struct {
	ID           int64  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	Role         string `db:"role"`
}

func fromAuthUserDTO(a *authUserDTO) (*domainAuth.AuthUser, error) {
	authUser, err := domainAuth.NewAuthUserFromValues(a.ID, a.Username, a.PasswordHash, a.Role)
	if err != nil {
		return nil, domainError.Strip(err)
	}
	return authUser, nil
}

func toAuthUserDTO(a *domainAuth.AuthUser) (*authUserDTO, error) {
	return &authUserDTO{
		ID:           a.ID.Value(),
		Username:     a.Username.Value(),
		PasswordHash: a.PasswordHash.Value(),
		Role:         a.Role.Value(),
	}, nil
}
