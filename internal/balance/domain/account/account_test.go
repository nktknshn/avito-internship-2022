package account_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

func TestAccount_Transfer_Success(t *testing.T) {
	acc1 := must.Must(domainAccount.NewAccountFromValues(1, 1, 10, 0))
	acc2 := must.Must(domainAccount.NewAccountFromValues(2, 2, 0, 0))
	amount := must.Must(amount.NewPositive(10))

	err := acc1.Transfer(acc2, amount)
	require.NoError(t, err)

	require.Equal(t, int64(0), acc1.Balance.GetAvailable().Value())
	require.Equal(t, int64(0), acc1.Balance.GetReserved().Value())
	require.Equal(t, int64(10), acc2.Balance.GetAvailable().Value())
	require.Equal(t, int64(0), acc2.Balance.GetReserved().Value())
}

func TestAccount_Transfer_SameAccount(t *testing.T) {
	acc1 := must.Must(domainAccount.NewAccountFromValues(1, 1, 10, 0))
	amount := must.Must(amount.NewPositive(10))

	err := acc1.Transfer(acc1, amount)
	require.ErrorIs(t, err, domainAccount.ErrSameAccount)
}

func TestAccount_Transfer_InsufficientAvailable(t *testing.T) {
	acc1 := must.Must(domainAccount.NewAccountFromValues(1, 1, 10, 0))

	acc2 := must.Must(domainAccount.NewAccountFromValues(2, 2, 0, 0))

	amount := must.Must(amount.NewPositive(11))

	err := acc1.Transfer(acc2, amount)
	require.ErrorIs(t, err, domainAccount.ErrInsufficientAvailableBalance)
}
