package report_revenue_export

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
)

type ReportRevenueExportUseCase struct {
}

func New() *ReportRevenueExportUseCase {
	return &ReportRevenueExportUseCase{}
}

func (u *ReportRevenueExportUseCase) Handle(ctx context.Context, in In) (Out, error) {
	return Out{}, nil
}

func (u *ReportRevenueExportUseCase) GetName() string {
	return use_cases.NameReportRevenueExport
}
