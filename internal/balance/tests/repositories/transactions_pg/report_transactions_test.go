package transactions_pg_test

import (
	"time"

	"github.com/google/uuid"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	test_transactions_pg "github.com/nktknshn/avito-internship-2022/internal/balance/tests/repositories/transactions_pg"
)

func (s *Suite) generateTransactions() {
	account1, err := s.accountsRepo.Save(s.Context(), &domainAccount.Account{
		UserID: fixtures.UserID,
	})
	s.Require().NoError(err)
	s.Require().NotNil(account1)

	account2, err := s.accountsRepo.Save(s.Context(), &domainAccount.Account{
		UserID: fixtures.UserID_2,
	})
	s.Require().NoError(err)
	s.Require().NotNil(account2)

	account3, err := s.accountsRepo.Save(s.Context(), &domainAccount.Account{
		UserID: fixtures.UserID_3,
	})
	s.Require().NoError(err)
	s.Require().NotNil(account3)

	transactionsDeposit := test_transactions_pg.GenerateTransactionsDeposit(account1.ID, fixtures.UserID, 100, test_transactions_pg.GenerateTransactionsDepositParams{
		TimeMin: time.Now().Add(-time.Hour * 24),
		TimeMax: time.Now(),

		AmountMin: 100,
		AmountMax: 1000,
	})

	for _, transaction := range transactionsDeposit {
		_, err := s.transactionsRepo.SaveTransactionDeposit(s.Context(), &transaction)
		s.Require().NoError(err)
	}

	transactionsSpend := test_transactions_pg.GenerateTransactionsSpend(account1.ID, fixtures.UserID, 100, test_transactions_pg.GenerateTransactionsSpendParams{
		TimeMin: time.Now().Add(-time.Hour * 24),
		TimeMax: time.Now(),

		AmountMin: 100,
		AmountMax: 1000,
	})

	for _, transaction := range transactionsSpend {
		_, err := s.transactionsRepo.SaveTransactionSpend(s.Context(), &transaction)
		s.Require().NoError(err)
	}

	transactionsTransfer := test_transactions_pg.GenerateTransactionsTransfer(account1.ID, account2.ID, 100, test_transactions_pg.GenerateTransactionsTransferParams{
		TimeMin: time.Now().Add(-time.Hour * 24),
		TimeMax: time.Now(),

		AmountMin: 100,
		AmountMax: 1000,
	})

	for _, transaction := range transactionsTransfer {
		_, err := s.transactionsRepo.SaveTransactionTransfer(s.Context(), &transaction)
		s.Require().NoError(err)
	}
}

func (s *Suite) TestReportTransactions_Success() {

	acc, err := s.accountsRepo.Save(s.Context(), &domainAccount.Account{
		UserID: fixtures.UserID,
	})

	s.Require().NoError(err)
	s.Require().NotNil(acc)

	acc2, err := s.accountsRepo.Save(s.Context(), &domainAccount.Account{
		UserID: fixtures.UserID_2,
	})

	s.Require().NoError(err)
	s.Require().NotNil(acc2)

	tc := time.Now()
	tu := tc.Add(time.Second)

	trSpend, err := domainTransaction.NewTransactionSpendFromValues(
		uuid.Nil,
		acc.ID.Value(),
		fixtures.UserID_i64,
		fixtures.OrderID_i64,
		fixtures.ProductID_i64,
		fixtures.AmountPositive_i64,
		domainTransaction.TransactionSpendStatusConfirmed.Value(),
		tc,
		tu,
	)

	s.Require().NoError(err)

	trSpend_saved, err := s.transactionsRepo.SaveTransactionSpend(s.Context(), trSpend)
	s.Require().NoError(err)
	s.Require().NotNil(trSpend_saved)

	trDeposit, err := domainTransaction.NewTransactionDepositFromValues(
		uuid.Nil,
		acc.ID.Value(),
		fixtures.UserID_i64,
		"test",
		domainTransaction.TransactionDepositStatusConfirmed.Value(),
		fixtures.AmountPositive_i64+1,
		tc,
		tu,
	)

	s.Require().NoError(err)

	trDeposit_saved, err := s.transactionsRepo.SaveTransactionDeposit(s.Context(), trDeposit)
	s.Require().NoError(err)
	s.Require().NotNil(trDeposit_saved)

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
	s.Require().NotNil(trTransfer_saved)

	report, err := s.transactionsRepo.GetTransactionsByUserID(s.Context(), fixtures.UserID, report_transactions.GetTransactionsQuery{
		Limit: 10,
	})

	s.Require().NoError(err)
	s.Require().NotNil(report)

	s.Require().Equal(len(report.Transactions), 3)

	for _, transaction := range report.Transactions {
		switch t := transaction.(type) {
		case domainTransaction.TransactionSpend:
			s.Require().Equal(trSpend_saved.Amount, t.Amount)
			s.Require().Equal(trSpend_saved.ID, t.ID)
			s.Require().Equal(trSpend_saved.CreatedAt, t.CreatedAt)
			s.Require().Equal(trSpend_saved.UpdatedAt, t.UpdatedAt)
			s.Require().Equal(trSpend_saved.Status, t.Status)
			s.Require().Equal(trSpend_saved.OrderID, t.OrderID)
			s.Require().Equal(trSpend_saved.ProductID, t.ProductID)
			s.Require().Equal(trSpend_saved.UserID, t.UserID)
			s.Require().Equal(trSpend_saved.AccountID, t.AccountID)
		case domainTransaction.TransactionDeposit:
			s.Require().Equal(trDeposit_saved.Amount, t.Amount)
			s.Require().Equal(trDeposit_saved.ID, t.ID)
			s.Require().Equal(trDeposit_saved.CreatedAt, t.CreatedAt)
			s.Require().Equal(trDeposit_saved.UpdatedAt, t.UpdatedAt)
			s.Require().Equal(trDeposit_saved.Status, t.Status)
			s.Require().Equal(trDeposit_saved.UserID, t.UserID)
			s.Require().Equal(trDeposit_saved.AccountID, t.AccountID)
			s.Require().Equal(trDeposit_saved.DepositSource, t.DepositSource)
		case domainTransaction.TransactionTransfer:
			s.Require().Equal(trTransfer_saved.Amount, t.Amount)
			s.Require().Equal(trTransfer_saved.ID, t.ID)
			s.Require().Equal(trTransfer_saved.CreatedAt, t.CreatedAt)
			s.Require().Equal(trTransfer_saved.UpdatedAt, t.UpdatedAt)
			s.Require().Equal(trTransfer_saved.Status, t.Status)
			s.Require().Equal(trTransfer_saved.FromAccountID, t.FromAccountID)
			s.Require().Equal(trTransfer_saved.ToAccountID, t.ToAccountID)
		}
	}
}
