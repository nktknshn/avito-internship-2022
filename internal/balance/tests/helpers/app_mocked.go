package helpers

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/stretchr/testify/mock"
)

type AppMocked struct {
	app.Application
	GetBalanceUseCaseMock        *GetBalanceUseCaseMock
	ReserveUseCaseMock           *ReserveUseCaseMock
	DepositUseCaseMock           *DepositUseCaseMock
	TransferUseCaseMock          *TransferUseCaseMock
	ReserveCancelUseCaseMock     *ReserveCancelUseCaseMock
	ReserveConfirmUseCaseMock    *ReserveConfirmUseCaseMock
	AuthSigninUseCaseMock        *AuthSigninUseCaseMock
	AuthSignupUseCaseMock        *AuthSignupUseCaseMock
	AuthValidateTokenUseCaseMock *AuthValidateTokenUseCaseMock
}

func NewMockedApp() *AppMocked {
	a := &AppMocked{
		GetBalanceUseCaseMock:        &GetBalanceUseCaseMock{},
		ReserveUseCaseMock:           &ReserveUseCaseMock{},
		DepositUseCaseMock:           &DepositUseCaseMock{},
		TransferUseCaseMock:          &TransferUseCaseMock{},
		ReserveCancelUseCaseMock:     &ReserveCancelUseCaseMock{},
		ReserveConfirmUseCaseMock:    &ReserveConfirmUseCaseMock{},
		AuthSigninUseCaseMock:        &AuthSigninUseCaseMock{},
		AuthSignupUseCaseMock:        &AuthSignupUseCaseMock{},
		AuthValidateTokenUseCaseMock: &AuthValidateTokenUseCaseMock{},
	}

	a.Application.GetBalance = a.GetBalanceUseCaseMock
	a.Application.Reserve = a.ReserveUseCaseMock
	a.Application.Deposit = a.DepositUseCaseMock
	a.Application.Transfer = a.TransferUseCaseMock
	a.Application.ReserveCancel = a.ReserveCancelUseCaseMock
	a.Application.ReserveConfirm = a.ReserveConfirmUseCaseMock
	a.Application.AuthSignin = a.AuthSigninUseCaseMock
	a.Application.AuthSignup = a.AuthSignupUseCaseMock
	a.Application.AuthValidateToken = a.AuthValidateTokenUseCaseMock

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
	return use_cases.GetBalance
}

type ReserveUseCaseMock struct {
	mock.Mock
}

func (m *ReserveUseCaseMock) Handle(ctx context.Context, in reserve.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *ReserveUseCaseMock) GetName() string {
	return use_cases.Reserve
}

type DepositUseCaseMock struct {
	mock.Mock
}

func (m *DepositUseCaseMock) Handle(ctx context.Context, in deposit.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *DepositUseCaseMock) GetName() string {
	return use_cases.Deposit
}

type TransferUseCaseMock struct {
	mock.Mock
}

func (m *TransferUseCaseMock) Handle(ctx context.Context, in transfer.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *TransferUseCaseMock) GetName() string {
	return use_cases.Transfer
}

type ReserveCancelUseCaseMock struct {
	mock.Mock
}

func (m *ReserveCancelUseCaseMock) Handle(ctx context.Context, in reserve_cancel.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *ReserveCancelUseCaseMock) GetName() string {
	return use_cases.ReserveCancel
}

type ReserveConfirmUseCaseMock struct {
	mock.Mock
}

func (m *ReserveConfirmUseCaseMock) Handle(ctx context.Context, in reserve_confirm.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *ReserveConfirmUseCaseMock) GetName() string {
	return use_cases.ReserveConfirm
}

type AuthSigninUseCaseMock struct {
	mock.Mock
}

func (m *AuthSigninUseCaseMock) Handle(ctx context.Context, in auth_signin.In) (auth_signin.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(auth_signin.Out), args.Error(1)
}

func (m *AuthSigninUseCaseMock) GetName() string {
	return use_cases.AuthSignin
}

type AuthSignupUseCaseMock struct {
	mock.Mock
}

func (m *AuthSignupUseCaseMock) Handle(ctx context.Context, in auth_signup.In) error {
	args := m.Called(ctx, in)
	return args.Error(0)
}

func (m *AuthSignupUseCaseMock) GetName() string {
	return use_cases.AuthSignup
}

type AuthValidateTokenUseCaseMock struct {
	mock.Mock
}

func (m *AuthValidateTokenUseCaseMock) Handle(ctx context.Context, in auth_validate_token.In) (auth_validate_token.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(auth_validate_token.Out), args.Error(1)
}

func (m *AuthValidateTokenUseCaseMock) GetName() string {
	return use_cases.AuthValidateToken
}
