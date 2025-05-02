package deposit

import (
	"context"
	"errors"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type depositHandler struct {
	auth           handlers_auth.AuthUseCase
	depositUseCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in deposit.In) error
	GetName() string
}

// @Summary      Deposit money
// @Description  Deposit money to account
// @Tags         deposit
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        payload   body      requestBody  true  "Payload"
// @Success      200  {object}  handlers_builder.ResultEmpty
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/balance/deposit [post]
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

type requestBody struct {
	UserID int64  `json:"user_id" example:"1"`
	Source string `json:"source" example:"credit card"`
	Amount int64  `json:"amount" example:"100"`
}

func makeDepositHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _    = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		payload = ergo.PayloadAttach[requestBody](b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {

		pl := payload.Get(r)

		in, err := deposit.NewInFromValues(
			pl.UserID,
			pl.Amount,
			pl.Source,
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
