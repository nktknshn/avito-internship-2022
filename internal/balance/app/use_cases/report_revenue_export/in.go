package report_revenue_export

import "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"

type In struct {
	year  report_revenue.Year
	month report_revenue.Month
}

func NewInFromValues(year int, month int) (In, error) {
	_year, err := report_revenue.NewYear(year)
	if err != nil {
		return In{}, err
	}
	_month, err := report_revenue.NewMonth(month)
	if err != nil {
		return In{}, err
	}
	return In{
		year:  _year,
		month: _month,
	}, nil
}

type Out struct {
	File []byte
}
