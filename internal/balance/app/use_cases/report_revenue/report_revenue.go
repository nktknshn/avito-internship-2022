package report_revenue

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
)

type ReportRevenueUseCase struct {
	repo ReportRevenueRepository
}

func New(repo ReportRevenueRepository) *ReportRevenueUseCase {

	if repo == nil {
		panic("repo is nil")
	}

	return &ReportRevenueUseCase{repo: repo}
}

func (u *ReportRevenueUseCase) Handle(ctx context.Context, in In) (Out, error) {

	out := Out{
		Records: make([]OutRecord, 0),
	}

	reportQuery := ReportRevenueQuery{
		Year:  in.year,
		Month: in.month,
	}

	report, err := u.repo.GetReportRevenueByMonth(ctx, reportQuery)

	if err != nil {
		return Out{}, err
	}

	for _, record := range report.Records {
		out.Records = append(out.Records, OutRecord(record))
	}

	return out, nil
}

func (u *ReportRevenueUseCase) GetName() string {
	return use_cases.NameReportRevenue
}
