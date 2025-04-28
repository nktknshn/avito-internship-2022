package transactions_pg_test

import (
	"time"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	transactions_pg "github.com/nktknshn/avito-internship-2022/internal/balance/tests/repositories/transactions_pg"
)

func (s *Suite) TestReportTransactions_Success() {
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

	transactionsDeposit := transactions_pg.GenerateTransactionsDeposit(account1.ID, fixtures.UserID, 100, transactions_pg.GenerateTransactionsDepositParams{
		TimeMin: time.Now().Add(-time.Hour * 24),
		TimeMax: time.Now(),

		AmountMin: 100,
		AmountMax: 1000,
	})

	for _, transaction := range transactionsDeposit {
		_, err := s.transactionsRepo.SaveTransactionDeposit(s.Context(), &transaction)
		s.Require().NoError(err)
	}

	transactionsSpend := transactions_pg.GenerateTransactionsSpend(account1.ID, fixtures.UserID, 100, transactions_pg.GenerateTransactionsSpendParams{
		TimeMin: time.Now().Add(-time.Hour * 24),
		TimeMax: time.Now(),

		AmountMin: 100,
		AmountMax: 1000,
	})

	for _, transaction := range transactionsSpend {
		_, err := s.transactionsRepo.SaveTransactionSpend(s.Context(), &transaction)
		s.Require().NoError(err)
	}

	transactionsTransfer := transactions_pg.GenerateTransactionsTransfer(account1.ID, account2.ID, 100, transactions_pg.GenerateTransactionsTransferParams{
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
