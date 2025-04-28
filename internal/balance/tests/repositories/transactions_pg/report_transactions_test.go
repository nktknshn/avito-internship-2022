package transactions_pg_test

import (
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
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
	s.generateTransactions()

	// rows := s.ExecSqlMust(sqlQuery)

	report, err := s.transactionsRepo.GetTransactionsByUserID(s.Context(), fixtures.UserID, report_transactions.GetTransactionsQuery{
		Limit: 10,
	})

	s.Require().NoError(err)
	s.Require().NotNil(report)
}
