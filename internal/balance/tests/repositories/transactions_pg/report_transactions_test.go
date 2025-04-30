package transactions_pg_test

import (
	"time"

	"github.com/google/uuid"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *Suite) TestReportTransactions_Success() {

	acc := s.getAccount1()
	acc2 := s.getAccount2()

	tc := time.Now()
	tu := tc.Add(time.Second)

	trSpend, err := domainTransaction.NewTransactionSpendFromValues(
		uuid.Nil,
		acc.ID.Value(),
		fixtures.UserID_i64,
		fixtures.OrderID_i64,
		fixtures.ProductID_i64,
		fixtures.ProductTitle_str,
		fixtures.AmountPositive_i64,
		domainTransaction.TransactionSpendStatusConfirmed.Value(),
		tc,
		tu,
	)

	s.Require().NoError(err)

	trSpend_saved, err := s.transactionsRepo.SaveTransactionSpend(s.Context(), trSpend)
	s.Require().NoError(err)

	trDeposit, err := domainTransaction.NewTransactionDepositFromValues(
		uuid.Nil,
		acc.ID.Value(),
		fixtures.UserID_i64,
		fixtures.DepositSource_str,
		domainTransaction.TransactionDepositStatusConfirmed.Value(),
		fixtures.AmountPositive_i64+1,
		tc,
		tu,
	)

	s.Require().NoError(err)

	trDeposit_saved, err := s.transactionsRepo.SaveTransactionDeposit(s.Context(), trDeposit)
	s.Require().NoError(err)

	trTransfer, err := domainTransaction.NewTransactionTransferFromValues(
		uuid.Nil,
		acc.ID.Value(),
		acc2.ID.Value(),
		fixtures.AmountPositive_i64+2,
		domainTransaction.TransactionTransferStatusConfirmed.Value(),
		tc,
		tu,
	)

	s.Require().NoError(err)

	trTransfer_saved, err := s.transactionsRepo.SaveTransactionTransfer(s.Context(), trTransfer)
	s.Require().NoError(err)

	trTransferReceive, err := domainTransaction.NewTransactionTransferFromValues(
		uuid.Nil,
		acc2.ID.Value(),
		acc.ID.Value(),
		fixtures.AmountPositive_i64+3,
		domainTransaction.TransactionTransferStatusConfirmed.Value(),
		tc,
		tu,
	)

	s.Require().NoError(err)

	trTransferReceive_saved, err := s.transactionsRepo.SaveTransactionTransfer(s.Context(), trTransferReceive)
	s.Require().NoError(err)

	//
	report, err := s.transactionsRepo.GetTransactionsByUserID(s.Context(), fixtures.UserID, report_transactions.GetTransactionsQuery{
		Limit: 10,
	})

	s.Require().NoError(err)
	s.Require().NotNil(report)

	s.Require().Equal(len(report.Transactions), 4)

	// Проверяем, что все транзакции нормально демаршалятся
	for _, transaction := range report.Transactions {
		switch t := transaction.(type) {
		case *domainTransaction.TransactionSpend:
			s.Require().Equal(trSpend_saved.Amount, t.Amount)
			s.Require().Equal(trSpend_saved.ID, t.ID)
			s.Require().Equal(trSpend_saved.CreatedAt, t.CreatedAt)
			s.Require().Equal(trSpend_saved.UpdatedAt, t.UpdatedAt)
			s.Require().Equal(trSpend_saved.Status, t.Status)
			s.Require().Equal(trSpend_saved.OrderID, t.OrderID)
			s.Require().Equal(trSpend_saved.ProductID, t.ProductID)
			s.Require().Equal(trSpend_saved.UserID, t.UserID)
			s.Require().Equal(trSpend_saved.AccountID, t.AccountID)
		case *domainTransaction.TransactionDeposit:
			s.Require().Equal(trDeposit_saved.Amount, t.Amount)
			s.Require().Equal(trDeposit_saved.ID, t.ID)
			s.Require().Equal(trDeposit_saved.CreatedAt, t.CreatedAt)
			s.Require().Equal(trDeposit_saved.UpdatedAt, t.UpdatedAt)
			s.Require().Equal(trDeposit_saved.Status, t.Status)
			s.Require().Equal(trDeposit_saved.UserID, t.UserID)
			s.Require().Equal(trDeposit_saved.AccountID, t.AccountID)
			s.Require().Equal(trDeposit_saved.DepositSource, t.DepositSource)
		case *domainTransaction.TransactionTransfer:
			if t.ToAccountID == acc.ID {
				s.Require().Equal(trTransferReceive_saved.Amount, t.Amount)
				s.Require().Equal(trTransferReceive_saved.ID, t.ID)
				s.Require().Equal(trTransferReceive_saved.CreatedAt, t.CreatedAt)
				s.Require().Equal(trTransferReceive_saved.UpdatedAt, t.UpdatedAt)
				s.Require().Equal(trTransferReceive_saved.Status, t.Status)
				s.Require().Equal(trTransferReceive_saved.FromAccountID, t.FromAccountID)
				s.Require().Equal(trTransferReceive_saved.ToAccountID, t.ToAccountID)
			} else {
				s.Require().Equal(trTransfer_saved.Amount, t.Amount)
				s.Require().Equal(trTransfer_saved.ID, t.ID)
				s.Require().Equal(trTransfer_saved.CreatedAt, t.CreatedAt)
				s.Require().Equal(trTransfer_saved.UpdatedAt, t.UpdatedAt)
				s.Require().Equal(trTransfer_saved.Status, t.Status)
				s.Require().Equal(trTransfer_saved.FromAccountID, t.FromAccountID)
				s.Require().Equal(trTransfer_saved.ToAccountID, t.ToAccountID)
			}
		default:
			s.FailNow("unknown transaction type", t)
		}
	}
}
