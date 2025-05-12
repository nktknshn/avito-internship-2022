package auth_signup

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/password_hasher"
)

type AuthSignupUseCase struct {
	trm      trm.Manager
	hasher   password_hasher.Hasher
	authRepo authRepo
}

type authRepo interface {
	CreateUser(
		ctx context.Context,
		username domainAuth.AuthUserUsername,
		passwordHash domainAuth.AuthUserPasswordHash,
		role domainAuth.AuthUserRole,
	) error
}

func New(trm trm.Manager, hasher password_hasher.Hasher, authRepo authRepo) *AuthSignupUseCase {

	if hasher == nil {
		panic("hasher is nil")
	}

	if authRepo == nil {
		panic("authRepo is nil")
	}

	if trm == nil {
		panic("trm is nil")
	}

	return &AuthSignupUseCase{trm: trm, hasher: hasher, authRepo: authRepo}
}

func (u *AuthSignupUseCase) Handle(ctx context.Context, in In) error {
	passwordHash, err := u.hasher.Hash(in.password.String())

	if err != nil {
		return err
	}

	err = u.authRepo.CreateUser(
		ctx,
		in.username,
		domainAuth.AuthUserPasswordHash(passwordHash),
		in.role,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *AuthSignupUseCase) GetName() string {
	return use_cases.NameAuthSignup
}
