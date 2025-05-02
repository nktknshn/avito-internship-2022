package reserve_cancel

import (
	"context"
	"errors"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type HandlerReserveCancel struct {
	auth          handlers_auth.AuthUseCase
	reserveCancel useCase
}

type useCase interface {
	Handle(ctx context.Context, in reserve_cancel.In) error
	GetName() string
}

// @Summary      Reserve cancel
// @Description  Cancel money reservation
// @Tags         reserve_cancel
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        payload   body      requestBody  true  "Payload"
// @Success      200  {object}  handlers_builder.ResultEmpty
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/balance/reserve/cancel [post]
func New(auth handlers_auth.AuthUseCase, reserveCancel useCase) *HandlerReserveCancel {
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

type requestBody struct {
	UserID    int64 `json:"user_id"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
	Amount    int64 `json:"amount"`
}

func makeHandlerReserveCancel(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _    = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		payload = ergo.PayloadAttach[requestBody](b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		pl := payload.Get(r)

		in, err := reserve_cancel.NewInFromValues(
			pl.UserID,
			pl.OrderID,
			pl.ProductID,
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
