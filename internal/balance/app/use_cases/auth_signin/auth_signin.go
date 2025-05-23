package auth_signin

import (
	"context"
	"errors"

	"github.com/avito-tech/go-transaction-manager/trm"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/password_hasher"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
)

type AuthSigninUseCase struct {
	trm      trm.Manager
	hasher   password_hasher.HashVerifier
	tokenGen token_generator.TokenGenerator[domainAuth.AuthUserTokenClaims]
	authRepo authRepo
}

type authRepo interface {
	GetUserByUsername(ctx context.Context, username domainAuth.AuthUserUsername) (*domainAuth.AuthUser, error)
}

func New(
	trm trm.Manager,
	hasher password_hasher.HashVerifier,
	tokenGen token_generator.TokenGenerator[domainAuth.AuthUserTokenClaims],
	authRepo authRepo,
) *AuthSigninUseCase {

	if trm == nil {
		panic("trm is nil")
	}

	if hasher == nil {
		panic("hasher is nil")
	}

	if tokenGen == nil {
		panic("tokenGen is nil")
	}

	if authRepo == nil {
		panic("authRepo is nil")
	}

	return &AuthSigninUseCase{
		trm:      trm,
		hasher:   hasher,
		tokenGen: tokenGen,
		authRepo: authRepo,
	}
}

// Handle проверяет пользователя по имени и паролю.
// Генерирует и возвращает токен
func (u *AuthSigninUseCase) Handle(ctx context.Context, in In) (Out, error) {
	user, err := u.authRepo.GetUserByUsername(ctx, in.username)

	if err != nil {
		return Out{}, err
	}

	if user == nil {
		return Out{}, errors.New("user == nil")
	}

	ok, err := u.hasher.Verify(in.password.String(), user.PasswordHash.Value())

	if err != nil {
		return Out{}, domainAuth.ErrInvalidAuthUserPassword.WithCause(err)
	}

	if !ok {
		return Out{}, domainAuth.ErrInvalidAuthUserPassword
	}

	token, err := u.tokenGen.GenerateToken(ctx, domainAuth.AuthUserTokenClaims{
		AuthUserID:   user.ID.Value(),
		AuthUserRole: user.Role.Value(),
	})

	if err != nil {
		return Out{}, err
	}

	authToken, err := domainAuth.NewAuthToken(token)

	if err != nil {
		return Out{}, err
	}

	return Out{Token: authToken}, nil
}

func (u *AuthSigninUseCase) GetName() string {
	return use_cases.NameAuthSignin
}
