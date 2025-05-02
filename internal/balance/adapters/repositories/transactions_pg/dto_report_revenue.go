package transactions_pg

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
)

type reportRevenueRecordDTO struct {
	ProductTitle string `db:"product_title"`
	ProductID    int64  `db:"product_id"`
	TotalRevenue int64  `db:"total_revenue"`
}

func fromReportRevenueDTO(dto *reportRevenueRecordDTO) (report_revenue.ReportRevenueRecord, error) {

	productID, err := domainProduct.NewProductID(dto.ProductID)

	if err != nil {
		return report_revenue.ReportRevenueRecord{}, domainError.Strip(err)
	}

	productTitle, err := domainProduct.NewProductTitle(dto.ProductTitle)

	if err != nil {
		return report_revenue.ReportRevenueRecord{}, domainError.Strip(err)
	}

	return report_revenue.ReportRevenueRecord{
		ProductID:    productID,
		ProductTitle: productTitle,
		TotalRevenue: dto.TotalRevenue,
	}, nil
}
