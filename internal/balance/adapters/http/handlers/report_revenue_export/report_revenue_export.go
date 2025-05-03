package report_revenue_export

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue_export"
)

type reportRevenueExportHandler struct {
	auth    handlers_auth.AuthUseCase
	useCase useCase
}

type useCase interface {
	Handle(ctx context.Context, in report_revenue_export.In) (report_revenue_export.Out, error)
	GetName() string
}

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

func makeReportRevenueExportHandler(auth handlers_auth.AuthUseCase, u useCase) http.Handler {
	return nil
}
