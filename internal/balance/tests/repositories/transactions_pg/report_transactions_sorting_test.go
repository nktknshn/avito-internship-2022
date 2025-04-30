package transactions_pg_test

import (
	"slices"
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *Suite) reportTransactionGetAll(q report_transactions.GetTransactionsQuery) []report_transactions.Transaction {

	result := make([]report_transactions.Transaction, 0)

	for {
		report, err := s.transactionsRepo.GetTransactionsByUserID(s.Context(), fixtures.UserID, q)
		s.Require().NoError(err)
		result = append(result, report.Transactions...)
		if !report.HasMore {
			break
		}
		q.Cursor = report.Cursor
	}

	return result
}

func (s *Suite) TestReportTransactions_Sorting_Amount() {

	acc1 := s.getAccount1()
	acc2 := s.getAccount2()
	tu := time.Now()

	var rtime = func() time.Time {
		return rTime(tu, tu.Add(time.Hour*24))
	}

	trs := []*transactionWrapper{
		tSpend(acc1, int64(55), rtime()),
		tSpend(acc1, int64(55), rtime()),
		tSpend(acc1, int64(55), rtime()),
		tDeposit(acc1, rInt64(1, 100), rtime()),
		tTransfer(acc1, acc2, rInt64(1, 200), rtime()),
		tTransfer(acc2, acc1, rInt64(1, 200), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tDeposit(acc1, rInt64(1, 100), rtime()),
		tDeposit(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tTransfer(acc2, acc1, rInt64(1, 200), rtime()),
		tTransfer(acc2, acc1, rInt64(1, 200), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
		tSpend(acc1, rInt64(1, 100), rtime()),
	}

	trs_saved, err := s.saveTransactions(trs)
	s.Require().NoError(err)

	trs_saved_sortedAmountAsc := slices.Clone(trs_saved)
	sortAmountAsc(trs_saved_sortedAmountAsc)

	trs_saved_sortedAmountDesc := slices.Clone(trs_saved)
	sortAmountDesc(trs_saved_sortedAmountDesc)

	trs_saved_sortedUpdatedAtAsc := slices.Clone(trs_saved)
	sortUpdatedAtAsc(trs_saved_sortedUpdatedAtAsc)

	trs_saved_sortedUpdatedAtDesc := slices.Clone(trs_saved)
	sortUpdatedAtDesc(trs_saved_sortedUpdatedAtDesc)

	type testCase struct {
		testName         string
		sorting          report_transactions.Sorting
		sortingDirection report_transactions.SortingDirection
		limit            int
		expected         []*transactionWrapper
	}

	testCases := []testCase{
		{
			testName:         "Sorting AmountA sc",
			sorting:          report_transactions.SortingAmount,
			sortingDirection: report_transactions.SortingDirectionAsc,
			limit:            100,
			expected:         trs_saved_sortedAmountAsc,
		},
		{
			testName:         "Sorting Amount Desc",
			sorting:          report_transactions.SortingAmount,
			sortingDirection: report_transactions.SortingDirectionDesc,
			limit:            100,
			expected:         trs_saved_sortedAmountDesc,
		},
		{
			testName:         "Sorting UpdatedAt Asc",
			sorting:          report_transactions.SortingUpdatedAt,
			sortingDirection: report_transactions.SortingDirectionAsc,
			limit:            100,
			expected:         trs_saved_sortedUpdatedAtAsc,
		},
		{
			testName:         "Sorting UpdatedAt Desc",
			sorting:          report_transactions.SortingUpdatedAt,
			sortingDirection: report_transactions.SortingDirectionDesc,
			limit:            100,
			expected:         trs_saved_sortedUpdatedAtDesc,
		},
		{
			testName:         "Sorting AmountAsc pagination limit 3",
			sorting:          report_transactions.SortingAmount,
			sortingDirection: report_transactions.SortingDirectionAsc,
			limit:            3,
			expected:         trs_saved_sortedAmountAsc,
		},
		{
			testName:         "Sorting AmountDesc pagination limit 3",
			sorting:          report_transactions.SortingAmount,
			sortingDirection: report_transactions.SortingDirectionDesc,
			limit:            3,
			expected:         trs_saved_sortedAmountDesc,
		},
		{
			testName:         "Sorting UpdatedAtAsc pagination limit 3",
			sorting:          report_transactions.SortingUpdatedAt,
			sortingDirection: report_transactions.SortingDirectionAsc,
			limit:            3,
			expected:         trs_saved_sortedUpdatedAtAsc,
		},
		{
			testName:         "Sorting UpdatedAtDesc pagination limit 3",
			sorting:          report_transactions.SortingUpdatedAt,
			sortingDirection: report_transactions.SortingDirectionDesc,
			limit:            3,
			expected:         trs_saved_sortedUpdatedAtDesc,
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			report := s.reportTransactionGetAll(report_transactions.GetTransactionsQuery{
				Limit:            report_transactions.Limit(testCase.limit),
				Sorting:          testCase.sorting,
				SortingDirection: testCase.sortingDirection,
			})
			s.Require().True(transactionsEqual(testCase.expected, report))
		})
	}

}
