package handlers_builder

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	ergo "github.com/nktknshn/go-ergo-handler"
)

func NewWithAuth(auth handlers_auth.AuthUseCase, roles []domainAuth.AuthUserRole) (*ergo.Builder, *handlers_auth.AttachedAuthParser) {

	if len(roles) == 0 {
		panic("roles are empty")
	}

	var (
		b          = ergo.New()
		validator  = handlers_auth.NewUseCaseToValidateToken(auth)
		authParser = handlers_auth.AuthParser.Attach(validator, b)
	)

	roleChecker := handlers_auth.NewRoleChecker(roles...)
	roleChecker.Attach(authParser, b)

	return b, authParser
}
