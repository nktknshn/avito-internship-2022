package use_cases_test

import (
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func TestUseCases(t *testing.T) {
	s := new(UseCasesSuiteIntegrationTest)
	s.SetPostgresMigrationsDir("../../../migrations/postgres")
	suite.Run(t, s)
}

type UseCasesSuiteIntegrationTest struct {
	testing_pg.TestSuitePg
	trm              trm.Manager
	accountsRepo     *accounts_pg.AccountsRepository
	transactionsRepo *transactions_pg.TransactionsRepository
	// use cases
	reserve            *reserve.ReserveUseCase
	deposit            *deposit.DepositUseCase
	reserveCancel      *reserve_cancel.ReserveCancelUseCase
	reserveConfirm     *reserve_confirm.ReserveConfirmUseCase
	transfer           *transfer.TransferUseCase
	reportTransactions *report_transactions.ReportTransactionsUseCase
	// mocked dependencies

}

func (s *UseCasesSuiteIntegrationTest) SetupTest() {
	s.trm = helpers.GetTrm(&s.TestSuitePg)
	s.accountsRepo = accounts_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
	s.transactionsRepo = transactions_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)

	s.reserve = reserve.New(s.trm, s.accountsRepo, s.transactionsRepo)
	s.deposit = deposit.New(s.trm, s.accountsRepo, s.transactionsRepo)
	s.reserveCancel = reserve_cancel.New(s.trm, s.accountsRepo, s.transactionsRepo)
	s.reserveConfirm = reserve_confirm.New(s.trm, s.accountsRepo, s.transactionsRepo)
	s.transfer = transfer.New(s.trm, s.accountsRepo, s.transactionsRepo)
	s.reportTransactions = report_transactions.New(s.transactionsRepo)
}

func (s *UseCasesSuiteIntegrationTest) TearDownTest() {
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *UseCasesSuiteIntegrationTest) newAccount(mods ...func(*domainAccount.Account)) *domainAccount.Account {
	acc := must.Must(domainAccount.NewAccount(fixtures.UserID))
	for _, mod := range mods {
		mod(acc)
	}
	return acc
}

func (s *UseCasesSuiteIntegrationTest) newAccountSaved(mods ...func(*domainAccount.Account)) *domainAccount.Account {
	acc := s.newAccount(mods...)
	acc, err := s.accountsRepo.Save(s.Context(), acc)
	s.Require().NoError(err)
	return acc
}
