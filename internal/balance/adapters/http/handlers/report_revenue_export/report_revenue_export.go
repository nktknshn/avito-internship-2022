package report_revenue_export

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_builder"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue_export"
	ergo "github.com/nktknshn/go-ergo-handler"
)

type reportRevenueExportHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in report_revenue_export.In) (report_revenue_export.Out, error)
	GetName() string
}

// @Summary      Report revenue export
// @ID           reportRevenueExport
// @Description  Report revenue export
// @Tags         report_revenue_export
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
// @Router       /api/v1/report/revenue/export [get]
func New(auth handlers_auth.AuthUseCase, useCase useCase) *reportRevenueExportHandler {
	if auth == nil {
		panic("auth is nil")
	}

	if useCase == nil {
		panic("useCase is nil")
	}

	return &reportRevenueExportHandler{auth, useCase}
}

func (h *reportRevenueExportHandler) GetHandler() http.Handler {
	return makeReportRevenueExportHandler(h.auth, h.useCase)
}

type responseBody struct {
	URL string `json:"url" example:"https://example.com/data/report_revenue_export.csv"`
}

func makeReportRevenueExportHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {

	var (
		b, _       = handlers_builder.NewWithAuthForUseCase(auth, u.GetName())
		paramYear  = ergo.QueryParamInt("year").Attach(b)
		paramMonth = ergo.QueryParamInt("month").Attach(b)
	)

	return b.BuildHandlerWrapped(func(w http.ResponseWriter, r *http.Request) (any, error) {
		in, err := report_revenue_export.NewInFromValues(
			paramYear.Get(r),
			paramMonth.Get(r),
		)

		if err != nil {
			return nil, err
		}

		resp, err := u.Handle(r.Context(), in)

		if err != nil {
			return nil, err
		}

		return responseBody{
			URL: resp.URL,
		}, nil
	})
}
