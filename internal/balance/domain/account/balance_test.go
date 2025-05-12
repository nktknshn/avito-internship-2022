package account_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

// TestAccountBalance_Equality проверяет, что AccountBalance является значением, а не объектом
func TestAccountBalance_Equality(t *testing.T) {
	av, err := amount.New(10)
	require.NoError(t, err)
	re, err := amount.New(100)
	require.NoError(t, err)

	b1, err := account.NewAccountBalance(av, re)
	require.NoError(t, err)

	b2, err := account.NewAccountBalance(av, re)
	require.NoError(t, err)

	if b1 != b2 {
		t.Fatalf("AccountBalance is not a real value object")
	}
}

func TestAccountBalance_Deposit(t *testing.T) {
	b1 := account.NewAccountBalanceZero()
	amount1 := must.Must(amount.NewPositive(10))

	b2, err := b1.Deposit(amount1)
	require.NoError(t, err)
	require.Equal(t, int64(10), b2.GetAvailable().Value())
	require.Equal(t, int64(0), b2.GetReserved().Value())
}

func TestAccountBalance_Reserve(t *testing.T) {
	b1, err := account.NewAccountBalanceFromValues(10, 0)
	require.NoError(t, err)
	amount1 := must.Must(amount.NewPositive(5))

	b2, err := b1.Reserve(amount1)
	require.NoError(t, err)
	require.Equal(t, int64(5), b2.GetAvailable().Value())
	require.Equal(t, int64(5), b2.GetReserved().Value())
}

func TestAccountBalance_ReserveCancel_Success(t *testing.T) {
	b1 := must.Must(account.NewAccountBalanceFromValues(5, 5))
	amount1 := must.Must(amount.NewPositive(5))

	b2, err := b1.ReserveCancel(amount1)
	require.NoError(t, err)
	require.Equal(t, int64(10), b2.GetAvailable().Value())
	require.Equal(t, int64(0), b2.GetReserved().Value())
}

func TestAccountBalance_ReserveCancel_InsufficientBalance(t *testing.T) {
	b1 := must.Must(account.NewAccountBalanceFromValues(5, 0))
	amount1 := must.Must(amount.NewPositive(5))

	_, err := b1.ReserveCancel(amount1)
	require.ErrorIs(t, err, account.ErrInsufficientReserveBalance)

}

func TestAccountBalance_ReserveConfirm_Success(t *testing.T) {
	b1 := must.Must(account.NewAccountBalanceFromValues(0, 5))
	amount1 := must.Must(amount.NewPositive(5))

	b2, err := b1.ReserveConfirm(amount1)
	require.NoError(t, err)

	require.Equal(t, int64(0), b2.GetAvailable().Value())
	require.Equal(t, int64(0), b2.GetReserved().Value())
}

func TestAccountBalance_ReserveConfirm_InsufficientBalance(t *testing.T) {
	b1 := must.Must(account.NewAccountBalanceFromValues(0, 0))
	amount1 := must.Must(amount.NewPositive(5))

	_, err := b1.ReserveConfirm(amount1)
	require.ErrorIs(t, err, account.ErrInsufficientReserveBalance)
}

func TestAccountBalance_Withdraw_Success(t *testing.T) {
	b1 := must.Must(account.NewAccountBalanceFromValues(10, 0))
	amount1 := must.Must(amount.NewPositive(5))

	b2, err := b1.Withdraw(amount1)
	require.NoError(t, err)
	require.Equal(t, int64(5), b2.GetAvailable().Value())
	require.Equal(t, int64(0), b2.GetReserved().Value())
}

func TestAccountBalance_Withdraw_InsufficientBalance(t *testing.T) {
	b1 := must.Must(account.NewAccountBalanceFromValues(5, 0))
	amount1 := must.Must(amount.NewPositive(10))

	_, err := b1.Withdraw(amount1)
	require.ErrorIs(t, err, account.ErrInsufficientAvailableBalance)
}
