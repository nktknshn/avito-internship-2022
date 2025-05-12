package report_transactions

import (
	"context"
	"errors"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type ReportTransactionsHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in report_transactions.In) (report_transactions.Out, error)
	GetName() string
}

// @Summary      Report transactions
// @ID           reportTransactions
// @Description  Report transactions
// @Tags         report_transactions
// @Accept       json
// @Produce      json
// @Security    Bearer
// @Param        user_id   path      int  true  "User ID"
// @Param        limit     query     int  true  "Limit"
// @Param        cursor    query     string  false  "Cursor"
// @Param        sorting   query     sortingType  true  "Sorting"
// @Param        sorting_direction query     sortingDirection  true  "Sorting Direction"
// @Success      200  {object}  handlers_builder.Result[responseBody]
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/report/transactions/{user_id} [get]
func New(auth handlers_auth.AuthUseCase, useCase useCase) *ReportTransactionsHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &ReportTransactionsHandler{auth, useCase}
}

func (h *ReportTransactionsHandler) GetHandler() http.Handler {
	return makeReportTransactionsHandler(h.auth, h.useCase)
}

type sortingType string

func (s sortingType) String() string {
	return string(s)
}

func (s sortingType) Parse(_ context.Context, v string) (sortingType, error) {
	if v == sortingUpdatedAt.String() ||
		v == sortingAmount.String() {
		return sortingType(v), nil
	}
	return "", errors.New("invalid sorting type")
}

const (
	sortingUpdatedAt sortingType = sortingType(report_transactions.SortingUpdatedAt)
	sortingAmount    sortingType = sortingType(report_transactions.SortingAmount)
)

type sortingDirection string

func (s sortingDirection) Parse(_ context.Context, v string) (sortingDirection, error) {
	if v == sortingDirectionAsc.String() ||
		v == sortingDirectionDesc.String() {
		return sortingDirection(v), nil
	}
	return "", errors.New("invalid sorting direction")
}

func (s sortingDirection) String() string {
	return string(s)
}

const (
	sortingDirectionAsc  sortingDirection = sortingDirection(report_transactions.SortingDirectionAsc)
	sortingDirectionDesc sortingDirection = sortingDirection(report_transactions.SortingDirectionDesc)
)

func makeReportTransactionsHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _                  = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		paramUserID           = ergo.RouterParamInt64("user_id").Attach(b)
		paramLimit            = ergo.QueryParamUInt64Maybe("limit").Attach(b)
		paramCursor           = ergo.QueryParamStringMaybe("cursor").Attach(b)
		paramSorting          = ergo.QueryParamWithParserMaybe[sortingType]("sorting").Attach(b)
		paramSortingDirection = ergo.QueryParamWithParserMaybe[sortingDirection]("sorting_direction").Attach(b)
	)

	return b.BuildHandlerWrapped(func(_ http.ResponseWriter, r *http.Request) (any, error) {
		in, err := report_transactions.NewInFromValues(
			paramUserID.Get(r),
			paramCursor.GetDefault(r, ""),
			paramLimit.GetDefault(r, uint64(0)),
			paramSorting.GetDefault(r, "").String(),
			paramSortingDirection.GetDefault(r, "").String(),
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

		return outToResponse(res), nil
	})
}
