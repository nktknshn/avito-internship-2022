package report_revenue

type In struct {
	year  Year
	month Month
}

func NewIn(year int, month int) (In, error) {
	_year, err := NewYear(year)
	if err != nil {
		return In{}, err
	}
	_month, err := NewMonth(month)
	if err != nil {
		return In{}, err
	}
	return In{year: _year, month: _month}, nil
}
