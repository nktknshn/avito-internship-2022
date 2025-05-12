package transactions_pg

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	accountTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

func getMonthRangeUTC(year int, month int) (time.Time, time.Time) {
	m := time.Month(month)
	t0 := time.Date(year, m, 1, 0, 0, 0, 0, time.UTC)

	if m == time.December {
		year++
		m = time.January
	} else {
		m++
	}

	t1 := time.Date(year, m, 1, 0, 0, 0, 0, time.UTC)

	return t0, t1
}

func (r *TransactionsRepository) GetReportRevenueByMonth(
	ctx context.Context,
	reportQuery report_revenue.ReportRevenueQuery,
) (report_revenue.ReportRevenueResponse, error) {
	t0, t1 := getMonthRangeUTC(
		reportQuery.Year.Value(),
		reportQuery.Month.Value(),
	)

	sqlQuery := `
	SELECT
		product_id,
		STRING_AGG(DISTINCT product_title, ',') as product_title,
		SUM(amount) as total_revenue
	FROM transactions_spend
	WHERE 
		created_at >= :t0 AND created_at < :t1
		AND status = :status
	GROUP BY product_id
	`

	tr := r.getter.DefaultTrOrDB(ctx, r.db)

	if tr == nil {
		return report_revenue.ReportRevenueResponse{}, errors.New("TransactionsRepository.GetReportRevenueByMonth: tr is nil")
	}

	sql, args, err := tr.BindNamed(sqlQuery, map[string]interface{}{
		"t0":     t0,
		"t1":     t1,
		"status": accountTransaction.TransactionSpendStatusConfirmed.Value(),
	})

	if err != nil {
		return report_revenue.ReportRevenueResponse{}, errors.Wrap(err, "TransactionsRepository.GetReportRevenueByMonth.BindNamed")
	}

	records := []reportRevenueRecordDTO{}

	err = tr.SelectContext(ctx, &records, sql, args...)

	if err != nil {
		return report_revenue.ReportRevenueResponse{}, errors.Wrap(err, "TransactionsRepository.GetReportRevenueByMonth.SelectContext")
	}

	modelRecords := make([]report_revenue.ReportRevenueRecord, len(records))

	for i, record := range records {
		modelRecords[i], err = fromReportRevenueDTO(&record)

		if err != nil {
			return report_revenue.ReportRevenueResponse{}, errors.Wrap(
				err,
				"TransactionsRepository.GetReportRevenueByMonth.fromReportRevenueDTO",
			)
		}
	}

	return report_revenue.ReportRevenueResponse{
		Records: modelRecords,
	}, nil
}
