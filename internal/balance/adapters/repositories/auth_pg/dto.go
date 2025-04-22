package auth_pg

import domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"

type authUserDTO struct {
	ID           int64  `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	Role         string `db:"role"`
}

func fromAuthUserDTO(a *authUserDTO) (*domainAuth.AuthUser, error) {
	return domainAuth.NewAuthUserFromValues(a.ID, a.Username, a.PasswordHash, a.Role)
}

func toAuthUserDTO(a *domainAuth.AuthUser) (*authUserDTO, error) {
	return &authUserDTO{
		ID:           a.ID.Value(),
		Username:     a.Username.Value(),
		PasswordHash: a.PasswordHash.Value(),
		Role:         a.Role.Value(),
	}, nil
}
