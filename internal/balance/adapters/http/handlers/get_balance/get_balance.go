package get_balance

import (
	"context"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
)

type GetBalanceHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in get_balance.In) (get_balance.Out, error)
	GetName() string
}

// @Summary      Get balance
// @ID           getBalance
// @Description  Get balance
// @Tags         balance
// @Accept       json
// @Produce      json
// @Security    Bearer
// @Param        user_id   path      int  true  "User ID" example=1
// @Success      200  {object}  handlers_builder.Result[responseBody]
// @Failure      400  {object}  handlers_builder.Error
// @Failure      404  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/balance/{user_id} [get]
func New(auth handlers_auth.AuthUseCase, useCase useCase) *GetBalanceHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &GetBalanceHandler{auth, useCase}
}

func (h *GetBalanceHandler) GetHandler() http.Handler {
	return makeGetBalanceHandler(h.auth, h.useCase)
}

type responseBody struct {
	Available int64 `json:"available"`
	Reserved  int64 `json:"reserved"`
}

func makeGetBalanceHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _        = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		paramUserID = ergo.RouterParamInt64("user_id").Attach(b)
	)

	return b.BuildHandlerWrapped(func(_ http.ResponseWriter, r *http.Request) (any, error) {
		in, err := get_balance.NewInFromValues(paramUserID.Get(r))

		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}

		result, err := u.Handle(r.Context(), in)

		if err != nil {
			return nil, err
		}

		return responseBody{
			Available: result.Available.Value(),
			Reserved:  result.Reserved.Value(),
		}, nil
	})
}
