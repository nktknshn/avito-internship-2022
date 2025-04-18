package handlers_builder

import (
	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
	ergo "github.com/nktknshn/go-ergo-handler"
)

func NewWithAuth(auth handlers_auth.TokenValidator, roles []domain.AuthUserRole) (*ergo.Builder, *handlers_auth.AttachedAuthParser) {
	var (
		b          = ergo.New()
		authParser = handlers_auth.AuthParser.Attach(auth, b)
	)

	roleChecker := handlers_auth.NewRoleChecker(roles...)
	roleChecker.Attach(authParser, b)

	return b, authParser
}
