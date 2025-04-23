package deposit

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type depositHandler struct {
	auth           handlers_auth.AuthUseCase
	depositUseCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in deposit.In) error
}

func New(auth handlers_auth.AuthUseCase, depositUseCase useCase) *depositHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if depositUseCase == nil {
		panic("depositUseCase is nil")
	}

	return &depositHandler{auth: auth, depositUseCase: depositUseCase}
}

func (h *depositHandler) GetHandler() http.Handler {
	return makeDepositHandler(h.auth, h.depositUseCase)
}

type payloadType struct {
	UserID int64  `json:"user_id"`
	Source string `json:"source"`
	Amount int64  `json:"amount"`
}

func (p payloadType) GetIn() (deposit.In, error) {
	return deposit.NewInFromValues(p.UserID, p.Amount, p.Source)
}

func makeDepositHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _ = handlers_builder.NewWithAuth(auth, []domainAuth.AuthUserRole{
			domainAuth.AuthUserRoleAdmin,
			domainAuth.AuthUserRoleAccount,
		})
		payload = ergo.PayloadAttach[payloadType](b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {

		pl := payload.Get(r)
		in, err := pl.GetIn()

		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}

		err = u.Handle(r.Context(), in)

		if err != nil {
			return nil, err
		}

		return nil, nil
	})
}
