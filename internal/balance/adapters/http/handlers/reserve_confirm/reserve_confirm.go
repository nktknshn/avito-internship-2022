package reserve_confirm

import (
	"context"
	"errors"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type HandlerReserveConfirm struct {
	auth           handlers_auth.AuthUseCase
	reserveConfirm useCase
}

type useCase interface {
	Handle(ctx context.Context, in reserve_confirm.In) error
	GetName() string
}

// @Summary      Reserve confirm
// @ID           reserveConfirm
// @Description  Confirm money reservation
// @Tags         reserve_confirm
// @Accept       json
// @Produce      json
// @Security    Bearer
// @Param        payload   body      requestBody  true  "Payload"
// @Success      200  {object}  handlers_builder.ResultEmpty
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      404  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/balance/reserve/confirm [post]
func New(auth handlers_auth.AuthUseCase, reserveConfirm useCase) *HandlerReserveConfirm {
	if auth == nil {
		panic("auth is nil")
	}

	if reserveConfirm == nil {
		panic("reserveConfirm is nil")
	}

	return &HandlerReserveConfirm{auth: auth, reserveConfirm: reserveConfirm}
}

func (h *HandlerReserveConfirm) GetHandler() http.Handler {
	return makeHandlerReserveConfirm(h.auth, h.reserveConfirm)
}

type requestBody struct {
	UserID    int64 `json:"user_id"    example:"1"`
	OrderID   int64 `json:"order_id"   example:"1"`
	ProductID int64 `json:"product_id" example:"1"`
	Amount    int64 `json:"amount"     example:"100"`
}

func makeHandlerReserveConfirm(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _    = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		payload = ergo.PayloadAttach[requestBody](b)
	)

	return b.BuildHandlerWrapped(func(_ http.ResponseWriter, r *http.Request) (any, error) {
		pl := payload.Get(r)
		in, err := reserve_confirm.NewInFromValues(
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

		if errors.Is(err, domainTransaction.ErrTransactionNotFound) {
			return nil, ergo.NewError(http.StatusNotFound, err)
		}

		if errors.Is(err, domainTransaction.ErrTransactionAlreadyPaid) {
			return nil, ergo.NewError(http.StatusConflict, err)
		}

		if err != nil {
			return nil, err
		}

		return nil, nil
	})
}
