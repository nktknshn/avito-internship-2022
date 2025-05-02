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

// @Summary      Report revenue
// @Description  Report revenue
// @Tags         report_revenue
// @Accept       json
// @Produce      json
// @Security    Bearer
// @Param        year   query      int  true  "Year"
// @Param        month  query      int  true  "Month"
// @Success      200  {object}  handlers_builder.Result[responseBody]
// @Failure      400  {object}  handlers_builder.Error
// @Failure      401  {object}  handlers_builder.Error
// @Failure      403  {object}  handlers_builder.Error
// @Failure      500  {object}  handlers_builder.Error
// @Router       /api/v1/report/revenue [get]
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

type responseRecord struct {
	ProductID    int64  `json:"product_id" example:"1"`
	ProductTitle string `json:"product_title" example:"delivery"`
	TotalRevenue int64  `json:"total_revenue" example:"100"`
}

type responseBody struct {
	Records []responseRecord `json:"records"`
}

func outToResult(out report_revenue.Out) responseBody {
	result := responseBody{
		Records: make([]responseRecord, len(out.Records)),
	}

	for i, record := range out.Records {
		result.Records[i] = responseRecord{
			ProductID:    record.ProductID.Value(),
			ProductTitle: record.ProductTitle.Value(),
			TotalRevenue: record.TotalRevenue,
		}
	}

	return result
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
		out, err := u.Handle(r.Context(), in)
		if err != nil {
			return nil, err
		}

		return outToResult(out), nil
	})
}
