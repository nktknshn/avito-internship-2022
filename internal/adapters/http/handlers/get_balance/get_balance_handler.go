package get_balance

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_params"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/get_balance"
)

type GetBalanceHandler struct {
	auth              handlers_auth.TokenValidator
	getBalanceUseCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in get_balance.In) (get_balance.Out, error)
}

func NewGetBalanceHandler(auth handlers_auth.TokenValidator, getBalanceUseCase *get_balance.GetBalanceUseCase) *GetBalanceHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if getBalanceUseCase == nil {
		panic("getBalanceUseCase is nil")
	}

	return &GetBalanceHandler{auth: auth, getBalanceUseCase: getBalanceUseCase}
}

func (h *GetBalanceHandler) GetHandler() http.Handler {
	return makeGetBalanceHandler(h.auth, h.getBalanceUseCase)
}

func makeGetBalanceHandler(auth handlers_auth.TokenValidator, u useCase) http.Handler {
	var (
		b, _ = handlers_builder.NewWithAuth(auth, []domain.AuthUserRole{
			domain.AuthUserRoleAdmin,
			domain.AuthUserRoleReport,
		})
		paramUserID = handlers_params.RouterParamUserID.Attach(b)
	)

	return b.BuildHandlerWrapped(func(h http.ResponseWriter, r *http.Request) (any, error) {
		return u.Handle(r.Context(), get_balance.In{
			UserID: paramUserID.Get(r),
		})
	})
}
