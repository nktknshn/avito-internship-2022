package auth_get_token

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

type AuthGetTokenUseCase struct {
	trm      trm.Manager
	authRepo authRepo
}

type authRepo interface {
	GetUserByUsername(ctx context.Context, username domainAuth.AuthUserUsername) (*domainAuth.AuthUser, error)
}

func NewAuthGetTokenUseCase(trm trm.Manager, authRepo authRepo) *AuthGetTokenUseCase {
	return &AuthGetTokenUseCase{trm: trm, authRepo: authRepo}
}

func (u *AuthGetTokenUseCase) Handle(ctx context.Context, in In) (Out, error) {
	return Out{}, nil
}
