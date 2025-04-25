package handlers_builder

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	ergo "github.com/nktknshn/go-ergo-handler"
)

func NewWithAuth(auth handlers_auth.AuthUseCase, roles []domainAuth.AuthUserRole) (*ergo.Builder, *handlers_auth.AttachedAuthParser) {

	if len(roles) == 0 {
		panic("roles are empty")
	}

	var (
		b = ergo.New().
			WithHandlerResultFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request, result any) {}).
			WithHandlerErrorFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
				if ergo.IsWrappedError(err) {
					ergo.DefaultHandlerErrorFunc(ctx, w, r, err)
					return
				}
				if domainError.IsDomainError(err) {
					ergo.DefaultHandlerErrorFunc(ctx, w, r, ergo.WrapWithStatusCode(err, http.StatusNotFound))
					return
				}
				http.Error(w, "internal server error", http.StatusInternalServerError)
			})
		validator  = handlers_auth.NewUseCaseToValidateToken(auth)
		authParser = handlers_auth.AuthParser.Attach(validator, b)
	)

	roleChecker := handlers_auth.NewRoleChecker(roles...)
	roleChecker.Attach(authParser, b)

	return b, authParser
}
