package handlers_auth

import (
	"context"
	"errors"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

type authParserKeyType string

const authParserKey authParserKeyType = "auth"

type TokenData struct {
	UserID domainAuth.AuthUserID
	Role   domainAuth.AuthUserRole
}

func (t *TokenData) GetRole() domainAuth.AuthUserRole {
	return t.Role
}

func (t *TokenData) GetUserID() domainAuth.AuthUserID {
	return t.UserID
}

type AuthUseCase interface {
	Handle(ctx context.Context, in auth_validate_token.In) (auth_validate_token.Out, error)
}

type UseCaseToValidateToken struct {
	useCase AuthUseCase
}

func NewUseCaseToValidateToken(useCase AuthUseCase) *UseCaseToValidateToken {
	return &UseCaseToValidateToken{useCase: useCase}
}

func (u *UseCaseToValidateToken) ValidateToken(ctx context.Context, token string) (*TokenData, bool, error) {

	in, err := auth_validate_token.NewInFromValues(token)

	if err != nil {
		return nil, false, ergo.WrapWithStatusCode(err, http.StatusBadRequest)
	}

	tokenData, err := u.useCase.Handle(ctx, in)

	if errors.Is(err, auth_validate_token.ErrInvalidToken) {
		return nil, false, ergo.WrapWithStatusCode(err, http.StatusBadRequest)
	}

	if errors.Is(err, auth_validate_token.ErrTokenExpired) {
		return nil, false, ergo.WrapWithStatusCode(err, http.StatusUnauthorized)
	}

	if err != nil {
		return nil, false, err
	}

	return &TokenData{
		UserID: tokenData.UserID,
		Role:   tokenData.Role,
	}, true, nil
}

var AuthParser = ergo.AuthParser[TokenData](authParserKey, ergo.TokenBearerFromHeader)

type AttachedAuthParser = ergo.AttachedAuthParser[TokenData, authParserKeyType]
