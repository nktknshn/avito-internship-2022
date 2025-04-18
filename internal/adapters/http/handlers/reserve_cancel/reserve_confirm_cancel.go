package reserve_cancel

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
	"github.com/nktknshn/avito-internship-2022/internal/use_cases/reserve_cancel"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type HandlerReserveCancel struct {
	auth          handlers_auth.TokenValidator
	reserveCancel useCase
}

type useCase interface {
	Handle(ctx context.Context, in reserve_cancel.In) error
}

func NewHandlerReserveCancel(auth handlers_auth.TokenValidator, reserveCancel useCase) *HandlerReserveCancel {
	if auth == nil {
		panic("auth is nil")
	}

	if reserveCancel == nil {
		panic("reserveCancel is nil")
	}

	return &HandlerReserveCancel{auth: auth, reserveCancel: reserveCancel}
}

func (h *HandlerReserveCancel) GetHandler() http.Handler {
	return makeHandlerReserveCancel(h.auth, h.reserveCancel)
}

type payloadType struct {
	UserID    int64 `json:"user_id"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Amount    int64 `json:"amount"`
}

func (p payloadType) GetIn() (reserve_cancel.In, error) {
	return reserve_cancel.NewInFromValues(
		p.UserID,
		p.OrderID,
		p.ProductID,
		p.Amount,
	)
}

func makeHandlerReserveCancel(auth handlers_auth.TokenValidator, u useCase) http.Handler {
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
