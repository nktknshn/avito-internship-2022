package transactions_pg_test

import (
	"context"
	"math/rand"
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
)

var products = []struct {
	productID   int64
	productName string
}{
	{1, "Product 1"},
	{2, "Product 2"},
	{3, "Product 3"},
	{4, "Product 4"},
	{5, "Product 5"},
	{6, "Product 6"},
	{7, "Product 7"},
	{8, "Product 8"},
	{9, "Product 9"},
	{10, "Product 10"},
}

func randomProduct() *struct {
	productID   int64
	productName string
} {
	return &products[rand.Intn(len(products))]
}

func randomDate(year int, month int) time.Time {
	t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	return t.Add(
		time.Duration(rand.Intn(31)*24*60) * time.Minute,
	)
}

func (s *Suite) TestGetReportRevenueByMonth() {
	acc := s.getAccount1()

	trs := []*transactionWrapper{}

	for month := range 12 {
		for _ = range 50 {
			t := tSpend(acc,
				rInt64(100, 1000),
				randomDate(2024, month+1),
			)
			t.setProduct(randomProduct())
			trs = append(trs, t)
		}
	}

	_, err := s.saveTransactions(trs)
	s.Require().NoError(err)

	report, err := s.transactionsRepo.GetReportRevenueByMonth(context.Background(), report_revenue.ReportRevenueQuery{
		Year:  2024,
		Month: 10,
	})
	s.Require().NoError(err)
	s.Require().Greater(len(report.Records), 0)
}
