package use_cases_test

import (
	"context"
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/suite"
)

func TestUseCases(t *testing.T) {
	s := new(SuiteTest)
	s.SetPostgresMigrationsDir("../../migrations/postgres")
	suite.Run(t, s)
}

type SuiteTest struct {
	testing_pg.TestSuitePg
	accountsRepo     *accounts_pg.AccountsRepository
	transactionsRepo *transactions_pg.TransactionsRepository
	// use cases
	reserve        *reserve.ReserveUseCase
	deposit        *deposit.DepositUseCase
	reserveCancel  *reserve_cancel.ReserveCancelUseCase
	reserveConfirm *reserve_confirm.ReserveConfirmUseCase
	authSignin     *auth_signin.AuthSigninUseCase
}

func (s *SuiteTest) SetupTest() {
	trm := helpers.GetTrm(&s.TestSuitePg)
	s.accountsRepo = accounts_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
	s.transactionsRepo = transactions_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)

	s.reserve = reserve.New(trm, s.accountsRepo, s.transactionsRepo)
	s.deposit = deposit.New(trm, s.accountsRepo, s.transactionsRepo)
	s.reserveCancel = reserve_cancel.New(trm, s.accountsRepo, s.transactionsRepo)
	s.reserveConfirm = reserve_confirm.New(trm, s.accountsRepo, s.transactionsRepo)
}

func (s *SuiteTest) TearDownTest() {
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *SuiteTest) newAccount(mods ...func(*domainAccount.Account)) *domainAccount.Account {
	acc := must.Must(domainAccount.NewAccount(fixtures.UserID))
	for _, mod := range mods {
		mod(acc)
	}
	return acc
}

func (s *SuiteTest) newAccountSaved(mods ...func(*domainAccount.Account)) *domainAccount.Account {
	acc := s.newAccount(mods...)
	acc, err := s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)
	return acc
}
