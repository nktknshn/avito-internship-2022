package report_transactions

import (
	"context"
	"errors"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type reportTransactionsHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in report_transactions.In) (report_transactions.Out, error)
	GetName() string
}

// @Summary      Report transactions
// @Description  Report transactions
// @Tags         report_transactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id   path      int  true  "User ID"
// @Param        limit     query     int  false  "Limit"
// @Param        cursor    query     string  false  "Cursor"
// @Param        sorting   query     string  false  "Sorting"
// @Param        sorting_direction query     string  false  "Sorting Direction"
// @Success      200  {object}  handlers_builder.Result[responseBody]
// @Failure      400  {object}  handlers_builder.Error
// @Router       /api/v1/report/transactions/{user_id} [get]
func New(auth handlers_auth.AuthUseCase, useCase useCase) *reportTransactionsHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &reportTransactionsHandler{auth, useCase}
}

func (h *reportTransactionsHandler) GetHandler() http.Handler {
	return makeReportTransactionsHandler(h.auth, h.useCase)
}

func makeReportTransactionsHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _                  = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		paramUserID           = ergo.RouterParamInt64("user_id").Attach(b)
		paramLimit            = ergo.QueryParamUInt64Maybe("limit").Attach(b)
		paramCursor           = ergo.QueryParamStringMaybe("cursor").Attach(b)
		paramSorting          = ergo.QueryParamStringMaybe("sorting").Attach(b)
		paramSortingDirection = ergo.QueryParamStringMaybe("sorting_direction").Attach(b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		in, err := report_transactions.NewInFromValues(
			paramUserID.Get(r),
			paramCursor.GetDefault(r, ""),
			paramLimit.GetDefault(r, uint64(0)),
			paramSorting.GetDefault(r, ""),
			paramSortingDirection.GetDefault(r, ""),
		)

		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}

		res, err := u.Handle(r.Context(), in)

		if errors.Is(err, domainAccount.ErrAccountNotFound) {
			return nil, ergo.NewError(http.StatusNotFound, err)
		}

		if err != nil {
			return nil, err
		}

		return outToResult(res), nil
	})
}
