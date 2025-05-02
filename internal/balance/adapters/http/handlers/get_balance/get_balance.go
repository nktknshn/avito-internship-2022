package get_balance

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type getBalanceHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in get_balance.In) (get_balance.Out, error)
	GetName() string
}

// @Summary      Get balance
// @Description  Get balance
// @Tags         balance
// @Accept       json
// @Produce      json
// @Param        user_id   path      int  true  "User ID"
// @Success      200  {object}  handlers_builder.Result[get_balance.Out]
// @Failure      400  {object}  handlers_builder.Error
// @Failure      404  {object}  handlers_builder.Error
// @Router       /api/v1/balance/{user_id} [get]
func New(auth handlers_auth.AuthUseCase, useCase useCase) *getBalanceHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &getBalanceHandler{auth, useCase}
}

func (h *getBalanceHandler) GetHandler() http.Handler {
	return makeGetBalanceHandler(h.auth, h.useCase)
}

func makeGetBalanceHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _        = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		paramUserID = ergo.RouterParamInt64("user_id").Attach(b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		fmt.Println("paramUserID.Get(r)", paramUserID.Get(r))

		in, err := get_balance.NewInFromValues(paramUserID.Get(r))
		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}
		result, err := u.Handle(r.Context(), in)

		if errors.Is(err, domainAccount.ErrAccountNotFound) {
			return nil, ergo.NewError(http.StatusNotFound, err)
		}

		if err != nil {
			return nil, err
		}

		return result, nil
	})
}
