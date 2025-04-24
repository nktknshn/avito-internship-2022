package auth_validate_token

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"
)

type AuthValidateTokenUseCase struct {
	trm      trm.Manager
	authRepo authRepo
}

type authRepo interface {
}

func New(trm trm.Manager, authRepo authRepo) *AuthValidateTokenUseCase {

	if trm == nil {
		panic("trm is nil")
	}

	if authRepo == nil {
		panic("authRepo is nil")
	}

	return &AuthValidateTokenUseCase{trm: trm, authRepo: authRepo}
}

// Проверить токен по блеклисту
// Проверить токен на валидность
func (u *AuthValidateTokenUseCase) Handle(ctx context.Context, in In) (Out, error) {
	return Out{}, nil
}
