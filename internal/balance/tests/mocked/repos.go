//nolint:nilaway // используем в defer
package mocked

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) Save(ctx context.Context, account *domainAccount.Account) (*domainAccount.Account, error) {
	args := m.Called(ctx, account)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainAccount.Account), args.Error(1)
}

func (m *AccountRepositoryMock) GetByUserID(ctx context.Context, userID domain.UserID) (*domainAccount.Account, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainAccount.Account), args.Error(1)
}

func (m *AccountRepositoryMock) GetByAccountID(ctx context.Context, accountID domainAccount.AccountID) (*domainAccount.Account, error) {
	args := m.Called(ctx, accountID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainAccount.Account), args.Error(1)
}

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) SaveTransactionDeposit(
	ctx context.Context,
	transaction *domainTransaction.TransactionDeposit,
) (*domainTransaction.TransactionDeposit, error) {
	args := m.Called(ctx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainTransaction.TransactionDeposit), args.Error(1)
}

func (m *TransactionRepositoryMock) SaveTransactionSpend(
	ctx context.Context,
	transaction *domainTransaction.TransactionSpend,
) (*domainTransaction.TransactionSpend, error) {
	args := m.Called(ctx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainTransaction.TransactionSpend), args.Error(1)
}

func (m *TransactionRepositoryMock) SaveTransactionTransfer(
	ctx context.Context,
	transaction *domainTransaction.TransactionTransfer,
) (*domainTransaction.TransactionTransfer, error) {
	args := m.Called(ctx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainTransaction.TransactionTransfer), args.Error(1)
}

func (m *TransactionRepositoryMock) GetTransactionSpendByOrderID(
	ctx context.Context,
	userID domain.UserID,
	orderID domain.OrderID,
) ([]*domainTransaction.TransactionSpend, error) {
	args := m.Called(ctx, userID, orderID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domainTransaction.TransactionSpend), args.Error(1)
}

type AuthRepositoryMock struct {
	mock.Mock
}

func (m *AuthRepositoryMock) GetUserByUsername(ctx context.Context, username domainAuth.AuthUserUsername) (*domainAuth.AuthUser, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domainAuth.AuthUser), args.Error(1)
}

func (m *AuthRepositoryMock) CreateUser(
	ctx context.Context,
	username domainAuth.AuthUserUsername,
	passwordHash domainAuth.AuthUserPasswordHash,
	role domainAuth.AuthUserRole,
) error {
	args := m.Called(ctx, username, passwordHash, role)
	return args.Error(0)
}

var _ domainAuth.AuthRepository = &AuthRepositoryMock{}
var _ domainAccount.AccountRepository = &AccountRepositoryMock{}
var _ domainTransaction.TransactionRepository = &TransactionRepositoryMock{}
