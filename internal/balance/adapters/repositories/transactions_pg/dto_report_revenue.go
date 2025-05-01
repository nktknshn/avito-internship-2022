package transactions_pg

import "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"

type reportRevenueRecordDTO struct {
	ProductTitle string `db:"product_title"`
	TotalRevenue int64  `db:"total_revenue"`
}

func fromReportRevenueDTO(dto *reportRevenueRecordDTO) (report_revenue.ReportRevenueRecord, error) {
	return report_revenue.ReportRevenueRecord{
		ProductTitle: dto.ProductTitle,
		TotalRevenue: dto.TotalRevenue,
	}, nil
}
