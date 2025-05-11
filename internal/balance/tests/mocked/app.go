//nolint:nilaway
package mocked

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue_export"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
	"github.com/stretchr/testify/mock"
)

type AppMocked struct {
	app.Application
	GetBalanceUseCaseMock          *GetBalanceUseCaseMock
	ReserveUseCaseMock             *ReserveUseCaseMock
	DepositUseCaseMock             *DepositUseCaseMock
	TransferUseCaseMock            *TransferUseCaseMock
	ReserveCancelUseCaseMock       *ReserveCancelUseCaseMock
	ReserveConfirmUseCaseMock      *ReserveConfirmUseCaseMock
	AuthSigninUseCaseMock          *AuthSigninUseCaseMock
	AuthSignupUseCaseMock          *AuthSignupUseCaseMock
	AuthValidateTokenUseCaseMock   *AuthValidateTokenUseCaseMock
	ReportTransactionsUseCaseMock  *ReportTransactionsUseCaseMock
	ReportRevenueUseCaseMock       *ReportRevenueUseCaseMock
	ReportRevenueExportUseCaseMock *ReportRevenueExportUseCaseMock
}

func NewMockedApp() *AppMocked {
	a := &AppMocked{
		GetBalanceUseCaseMock:          &GetBalanceUseCaseMock{},
		ReserveUseCaseMock:             &ReserveUseCaseMock{},
		DepositUseCaseMock:             &DepositUseCaseMock{},
		TransferUseCaseMock:            &TransferUseCaseMock{},
		ReserveCancelUseCaseMock:       &ReserveCancelUseCaseMock{},
		ReserveConfirmUseCaseMock:      &ReserveConfirmUseCaseMock{},
		AuthSigninUseCaseMock:          &AuthSigninUseCaseMock{},
		AuthSignupUseCaseMock:          &AuthSignupUseCaseMock{},
		AuthValidateTokenUseCaseMock:   &AuthValidateTokenUseCaseMock{},
		ReportTransactionsUseCaseMock:  &ReportTransactionsUseCaseMock{},
		ReportRevenueUseCaseMock:       &ReportRevenueUseCaseMock{},
		ReportRevenueExportUseCaseMock: &ReportRevenueExportUseCaseMock{},
	}

	a.Application.GetBalance = a.GetBalanceUseCaseMock
	a.Application.Reserve = a.ReserveUseCaseMock
	a.Application.Deposit = a.DepositUseCaseMock
	a.Application.Transfer = a.TransferUseCaseMock
	a.Application.ReserveCancel = a.ReserveCancelUseCaseMock
	a.Application.ReserveConfirm = a.ReserveConfirmUseCaseMock
	a.Application.ReportTransactions = a.ReportTransactionsUseCaseMock
	a.Application.ReportRevenue = a.ReportRevenueUseCaseMock

	a.Application.AuthSignin = a.AuthSigninUseCaseMock
	a.Application.AuthSignup = a.AuthSignupUseCaseMock
	a.Application.AuthValidateToken = a.AuthValidateTokenUseCaseMock
	a.Application.ReportRevenueExport = a.ReportRevenueExportUseCaseMock
	return a
}

type GetBalanceUseCaseMock struct {
	mock.Mock
}

func (m *GetBalanceUseCaseMock) Handle(ctx context.Context, in get_balance.In) (get_balance.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(get_balance.Out), args.Error(1)
}

func (m *GetBalanceUseCaseMock) GetName() string {
	return use_cases.NameGetBalance
}

type ReserveUseCaseMock struct {
	mock.Mock
}

func (m *ReserveUseCaseMock) Handle(ctx context.Context, in reserve.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *ReserveUseCaseMock) GetName() string {
	return use_cases.NameReserve
}

type DepositUseCaseMock struct {
	mock.Mock
}

func (m *DepositUseCaseMock) Handle(ctx context.Context, in deposit.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *DepositUseCaseMock) GetName() string {
	return use_cases.NameDeposit
}

type TransferUseCaseMock struct {
	mock.Mock
}

func (m *TransferUseCaseMock) Handle(ctx context.Context, in transfer.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *TransferUseCaseMock) GetName() string {
	return use_cases.NameTransfer
}

type ReserveCancelUseCaseMock struct {
	mock.Mock
}

func (m *ReserveCancelUseCaseMock) Handle(ctx context.Context, in reserve_cancel.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *ReserveCancelUseCaseMock) GetName() string {
	return use_cases.NameReserveCancel
}

type ReserveConfirmUseCaseMock struct {
	mock.Mock
}

func (m *ReserveConfirmUseCaseMock) Handle(ctx context.Context, in reserve_confirm.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *ReserveConfirmUseCaseMock) GetName() string {
	return use_cases.NameReserveConfirm
}

type AuthSigninUseCaseMock struct {
	mock.Mock
}

func (m *AuthSigninUseCaseMock) Handle(ctx context.Context, in auth_signin.In) (auth_signin.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(auth_signin.Out), args.Error(1)
}

func (m *AuthSigninUseCaseMock) GetName() string {
	return use_cases.NameAuthSignin
}

type AuthSignupUseCaseMock struct {
	mock.Mock
}

func (m *AuthSignupUseCaseMock) Handle(ctx context.Context, in auth_signup.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *AuthSignupUseCaseMock) GetName() string {
	return use_cases.NameAuthSignup
}

type AuthValidateTokenUseCaseMock struct {
	mock.Mock
}

func (m *AuthValidateTokenUseCaseMock) Handle(ctx context.Context, in auth_validate_token.In) (auth_validate_token.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(auth_validate_token.Out), args.Error(1)
}

func (m *AuthValidateTokenUseCaseMock) GetName() string {
	return use_cases.NameAuthValidateToken
}

func (app *AppMocked) SetupAuth(token string, returnOut auth_validate_token.Out, returnErr error) {
	authIn := must.Must(auth_validate_token.NewInFromValues(token))
	app.AuthValidateTokenUseCaseMock.On("Handle", mock.Anything, authIn).Return(returnOut, returnErr)
}

type ReportTransactionsUseCaseMock struct {
	mock.Mock
}

func (m *ReportTransactionsUseCaseMock) Handle(ctx context.Context, in report_transactions.In) (report_transactions.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(report_transactions.Out), args.Error(1)
}

func (m *ReportTransactionsUseCaseMock) GetName() string {
	return use_cases.NameReportTransactions
}

type ReportRevenueUseCaseMock struct {
	mock.Mock
}

func (m *ReportRevenueUseCaseMock) Handle(ctx context.Context, in report_revenue.In) (report_revenue.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(report_revenue.Out), args.Error(1)
}

func (m *ReportRevenueUseCaseMock) GetName() string {
	return use_cases.NameReportRevenue
}

type ReportRevenueExportUseCaseMock struct {
	mock.Mock
}

func (m *ReportRevenueExportUseCaseMock) Handle(ctx context.Context, in report_revenue_export.In) (report_revenue_export.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(report_revenue_export.Out), args.Error(1)
}

func (m *ReportRevenueExportUseCaseMock) GetName() string {
	return use_cases.NameReportRevenueExport
}
