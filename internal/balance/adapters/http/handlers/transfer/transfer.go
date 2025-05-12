package transfer

import (
	"context"
	"errors"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type TransferHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in transfer.In) error
	GetName() string
}

// @Summary      Transfer
// @ID           transfer
// @Description  Transfer money between accounts
// @Tags         transfer
// @Accept       json
// @Produce      json
// @Security    Bearer
// @Param        payload   body      requestBody  true  "Payload"
// @Success      200  {object}  handlers_builder.ResultEmpty
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/balance/transfer [post]
func New(auth handlers_auth.AuthUseCase, useCase useCase) *TransferHandler {

	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &TransferHandler{auth: auth, useCase: useCase}
}

func (h *TransferHandler) GetHandler() http.Handler {
	return makeTransferHandler(h.auth, h.useCase)
}

type requestBody struct {
	FromUserID int64 `json:"from_user_id" example:"1"`
	ToUserID   int64 `json:"to_user_id"   example:"2"`
	Amount     int64 `json:"amount"       example:"100"`
}

func makeTransferHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _    = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		payload = ergo.PayloadAttach[requestBody](b)
	)

	return b.BuildHandlerWrapped(func(_ http.ResponseWriter, r *http.Request) (any, error) {
		pl := payload.Get(r)

		in, err := transfer.NewInFromValues(
			pl.FromUserID,
			pl.ToUserID,
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
