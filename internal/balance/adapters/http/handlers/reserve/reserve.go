package reserve

import (
	"context"
	"errors"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type HandlerReserve struct {
	auth           handlers_auth.AuthUseCase
	reserveUseCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in reserve.In) error
	GetName() string
}

func New(auth handlers_auth.AuthUseCase, reserveUseCase useCase) *HandlerReserve {
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

type requestBody struct {
	UserID       int64  `json:"user_id"`
	ProductID    int64  `json:"product_id"`
	ProductTitle string `json:"product_title"`
	OrderID      int64  `json:"order_id"`
	Amount       int64  `json:"amount"`
}

func makeHandlerReserve(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _    = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		payload = ergo.PayloadAttach[requestBody](b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		pl := payload.Get(r)

		in, err := reserve.NewInFromValues(
			pl.UserID,
			pl.ProductID,
			pl.ProductTitle,
			pl.OrderID,
			pl.Amount,
		)

		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}

		err = u.Handle(r.Context(), in)

		if errors.Is(err, domainAccount.ErrAccountNotFound) {
			return nil, ergo.NewError(http.StatusNotFound, err)
		}

		if err != nil {
			return nil, err
		}

		return nil, nil
	})
}
