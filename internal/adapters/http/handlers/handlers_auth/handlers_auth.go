package handlers_auth

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/domain"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type authParserKeyType string

const authParserKey authParserKeyType = "auth"

type TokenData interface {
	GetAuthUserID() domain.AuthUserID
	GetAuthUserRole() domain.AuthUserRole
}

type TokenValidator interface {
	ValidateToken(ctx context.Context, token string) (*TokenData, bool, error)
}

var AuthParser = ergo.AuthParser[TokenData](authParserKey, ergo.TokenBearerFromHeader)

type AttachedAuthParser = ergo.AttachedAuthParser[TokenData, authParserKeyType]
