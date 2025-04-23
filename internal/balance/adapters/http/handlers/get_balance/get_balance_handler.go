package get_balance

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_params"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

type getBalanceHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in get_balance.In) (get_balance.Out, error)
}

func New(auth handlers_auth.AuthUseCase, useCase useCase) *getBalanceHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &getBalanceHandler{auth, useCase}
}

func (h *getBalanceHandler) GetHandler() http.Handler {
	return makeGetBalanceHandler(h.auth, h.useCase)
}

func makeGetBalanceHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _ = handlers_builder.NewWithAuth(auth, []domainAuth.AuthUserRole{
			domainAuth.AuthUserRoleAdmin,
			domainAuth.AuthUserRoleReport,
		})
		paramUserID = handlers_params.RouterParamUserID.Attach(b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		return u.Handle(r.Context(), get_balance.In{
			UserID: paramUserID.Get(r),
		})
	})
}
