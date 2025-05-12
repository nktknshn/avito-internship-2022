package handlers_builder

import (
	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/roles"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

func NewPublic() *ergo.Builder {
	return ergo.New().
		WithHandlerErrorFunc(handlerErrorFunc).
		WithHandlerResultFunc(handlerResultFunc)
}

func NewWithAuth(auth handlers_auth.AuthUseCase, roles []domainAuth.AuthUserRole) (*ergo.Builder, *handlers_auth.AttachedAuthParser) {

	if len(roles) == 0 {
		panic("roles are empty")
	}

	var (
		b = ergo.New().
			WithHandlerErrorFunc(handlerErrorFunc).
			WithHandlerResultFunc(handlerResultFunc)
		validator  = handlers_auth.NewUseCaseToValidateToken(auth)
		authParser = handlers_auth.AuthParser.Attach(validator, b)
	)

	roleChecker := handlers_auth.NewRoleChecker(roles...)
	roleChecker.Attach(authParser, b)

	return b, authParser
}

func NewWithAuthForUseCase(auth handlers_auth.AuthUseCase, useCase string) (*ergo.Builder, *handlers_auth.AttachedAuthParser) {
	roles := roles.GetUseCaseAuthUserRolesMust(useCase)
	return NewWithAuth(auth, roles)
}
