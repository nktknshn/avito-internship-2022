package reserve

import (
	"context"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
)

type HandlerReserve struct {
	auth           handlers_auth.AuthUseCase
	reserveUseCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in reserve.In) error
	GetName() string
}

// @Summary      Reserve
// @ID           reserve
// @Description  Reserve money
// @Tags         reserve
// @Accept       json
// @Produce      json
// @Security    Bearer
// @Param        payload   body      requestBody  true  "Payload"
// @Success      200  {object}  handlers_builder.ResultEmpty
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/balance/reserve [post]
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
	UserID       int64  `json:"user_id"       example:"1"`
	ProductID    int64  `json:"product_id"    example:"1"`
	ProductTitle string `json:"product_title" example:"delivery"`
	OrderID      int64  `json:"order_id"      example:"1"`
	Amount       int64  `json:"amount"        example:"100"`
}

func makeHandlerReserve(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _    = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		payload = ergo.PayloadAttach[requestBody](b)
	)

	return b.BuildHandlerWrapped(func(_ http.ResponseWriter, r *http.Request) (any, error) {
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

		if err != nil {
			return nil, err
		}

		return nil, nil
	})
}
