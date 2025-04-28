package auth_validate_token

import (
	"context"
	"errors"

	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
)

type AuthValidateTokenUseCase struct {
	trm            trm.Manager
	tokenValidator token_generator.TokenValidator[domainAuth.AuthUserTokenClaims]
	authRepo       authRepo
}

type authRepo interface{}

func New(
	trm trm.Manager,
	tokenValidator token_generator.TokenValidator[domainAuth.AuthUserTokenClaims],
	authRepo authRepo,
) *AuthValidateTokenUseCase {

	if trm == nil {
		panic("trm is nil")
	}

	if authRepo == nil {
		panic("authRepo is nil")
	}

	if tokenValidator == nil {
		panic("tokenValidator is nil")
	}

	return &AuthValidateTokenUseCase{trm: trm, tokenValidator: tokenValidator, authRepo: authRepo}
}

// TODO: Проверить токен по блеклисту
// Проверить токен на валидность
func (u *AuthValidateTokenUseCase) Handle(ctx context.Context, in In) (Out, error) {

	claims, err := u.tokenValidator.ValidateToken(ctx, in.token)

	if errors.Is(err, token_generator.ErrInvalidToken) {
		return Out{}, ErrInvalidToken
	}

	if errors.Is(err, token_generator.ErrTokenExpired) {
		return Out{}, ErrTokenExpired
	}

	if err != nil {
		return Out{}, err
	}

	return Out{
		UserID: claims.AuthUserID,
		Role:   claims.AuthUserRole,
	}, nil
}

func (u *AuthValidateTokenUseCase) GetName() string {
	return use_cases.NameAuthValidateToken
}
