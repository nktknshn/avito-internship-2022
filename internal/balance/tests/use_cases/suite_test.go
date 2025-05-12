package use_cases_test

import (
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/mocked"
)

type SuiteTest struct {
	suite.Suite
	trm trm.Manager

	mockedAccountsRepo     *mocked.AccountRepositoryMock
	mockedTransactionsRepo *mocked.TransactionRepositoryMock
	// use cases with mocked dependencies
	mockedReserve        *reserve.ReserveUseCase
	mockedDeposit        *deposit.DepositUseCase
	mockedReserveCancel  *reserve_cancel.ReserveCancelUseCase
	mockedReserveConfirm *reserve_confirm.ReserveConfirmUseCase
	// auth
}

func (s *SuiteTest) SetupTest() {
	s.trm = &mocked.TrmManagerMock{}

	// mocked dependencies
	s.mockedAccountsRepo = &mocked.AccountRepositoryMock{}
	s.mockedTransactionsRepo = &mocked.TransactionRepositoryMock{}

	// use cases with mocked dependencies
	s.mockedReserve = reserve.New(s.trm, s.mockedAccountsRepo, s.mockedTransactionsRepo)
	s.mockedDeposit = deposit.New(s.trm, s.mockedAccountsRepo, s.mockedTransactionsRepo)
	s.mockedReserveCancel = reserve_cancel.New(s.trm, s.mockedAccountsRepo, s.mockedTransactionsRepo)
	s.mockedReserveConfirm = reserve_confirm.New(s.trm, s.mockedAccountsRepo, s.mockedTransactionsRepo)
}
