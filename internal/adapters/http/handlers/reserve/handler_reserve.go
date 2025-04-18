package handler_reserve

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/reserve"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type HandlerReserve struct {
	auth           handlers_auth.TokenValidator
	reserveUseCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in reserve.In) error
}

func NewHandlerReserve(auth handlers_auth.TokenValidator, reserveUseCase useCase) *HandlerReserve {
	if auth == nil {
		panic("auth is nil")
	}

	if reserveUseCase == nil {
		panic("reserveUseCase is nil")
	}

	return &HandlerReserve{auth: auth, reserveUseCase: reserveUseCase}
}

func (h *HandlerReserve) GetHandler() http.Handler {
	return makeHandlerReserve(h.auth, h.reserveUseCase)
}

type payloadType struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	OrderID   int64 `json:"order_id"`
	Amount    int64 `json:"amount"`
}

func (p payloadType) GetIn() (reserve.In, error) {
	return reserve.NewInFromValues(
		p.UserID,
		p.ProductID,
		p.OrderID,
		p.Amount,
	)
}

func makeHandlerReserve(auth handlers_auth.TokenValidator, u useCase) http.Handler {
	var (
		b, _ = handlers_builder.NewWithAuth(auth, []domain.AuthUserRole{
			domain.AuthUserRoleAdmin,
			domain.AuthUserRoleAccount,
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
