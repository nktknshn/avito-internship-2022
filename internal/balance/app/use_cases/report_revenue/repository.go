package report_revenue

import (
	"context"

	useCaseError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
)

var (
	ErrInvalidYear  = useCaseError.New("invalid year")
	ErrInvalidMonth = useCaseError.New("invalid month")
)

type ReportRevenueRepository interface {
	GetReportRevenueByMonth(ctx context.Context, reportQuery ReportRevenueQuery) (ReportRevenueResponse, error)
}

type Year int

func (y Year) Value() int {
	return int(y)
}

func NewYear(year int) (Year, error) {
	if year <= 0 {
		return 0, ErrInvalidYear
	}
	return Year(year), nil
}

type Month int

func (m Month) Value() int {
	return int(m)
}

func NewMonth(month int) (Month, error) {
	if month < 1 || month > 12 {
		return 0, ErrInvalidMonth
	}
	return Month(month), nil
}

type ReportRevenueQuery struct {
	Year  Year
	Month Month
}

type ReportRevenueResponse struct {
	Records []ReportRevenueRecord
}

type ReportRevenueRecord struct {
	ProductID    domainProduct.ProductID
	ProductTitle domainProduct.ProductTitle
	TotalRevenue int64
}
