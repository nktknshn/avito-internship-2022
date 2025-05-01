package report_revenue

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type reportRevenueHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in report_revenue.In) (report_revenue.Out, error)
	GetName() string
}

func New(auth handlers_auth.AuthUseCase, useCase useCase) *reportRevenueHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &reportRevenueHandler{auth, useCase}
}

func (h *reportRevenueHandler) GetHandler() http.Handler {
	return makeReportRevenueHandler(h.auth, h.useCase)
}

func makeReportRevenueHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	var (
		b, _       = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		paramYear  = ergo.QueryParamInt("year").Attach(b)
		paramMonth = ergo.QueryParamInt("month").Attach(b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		in, err := report_revenue.NewInFromValues(
			paramYear.Get(r),
			paramMonth.Get(r),
		)

		if err != nil {
			return nil, ergo.NewError(http.StatusBadRequest, err)
		}

		return u.Handle(r.Context(), in)
	})
}
