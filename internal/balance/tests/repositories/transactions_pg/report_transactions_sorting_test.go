package transactions_pg_test

import (
	"slices"
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

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
	}

	trSpends_saved, err := s.saveTransactions(trs)
	s.Require().NoError(err)

	trSpends_saved_sortedAmountAsc := slices.Clone(trSpends_saved)
	sortAmountAsc(trSpends_saved_sortedAmountAsc)

	trSpends_saved_sortedAmountDesc := slices.Clone(trSpends_saved)
	sortAmountDesc(trSpends_saved_sortedAmountDesc)

	trSpends_saved_sortedUpdatedAtAsc := slices.Clone(trSpends_saved)
	sortUpdatedAtAsc(trSpends_saved_sortedUpdatedAtAsc)

	trSpends_saved_sortedUpdatedAtDesc := slices.Clone(trSpends_saved)
	sortUpdatedAtDesc(trSpends_saved_sortedUpdatedAtDesc)

	report, err := s.transactionsRepo.GetTransactionsByUserID(s.Context(), fixtures.UserID, report_transactions.GetTransactionsQuery{
		Limit:            100,
		Sorting:          report_transactions.SortingAmount,
		SortingDirection: report_transactions.SortingDirectionAsc,
	})

	s.Require().NoError(err)
	s.Require().True(transactionsEqual(trSpends_saved_sortedAmountAsc, report.Transactions))

	report, err = s.transactionsRepo.GetTransactionsByUserID(s.Context(), fixtures.UserID, report_transactions.GetTransactionsQuery{
		Limit:            100,
		Sorting:          report_transactions.SortingAmount,
		SortingDirection: report_transactions.SortingDirectionDesc,
	})

	s.Require().NoError(err)
	s.Require().True(transactionsEqual(trSpends_saved_sortedAmountDesc, report.Transactions))

}
