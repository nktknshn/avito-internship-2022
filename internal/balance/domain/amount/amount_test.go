package amount_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

func TestAmount_New(t *testing.T) {
	amount, err := domainAmount.New(10)
	require.NoError(t, err)
	require.Equal(t, int64(10), amount.Value())

	_, err = domainAmount.New(-10)
	require.Error(t, err)

}

func TestAmountPositive_New(t *testing.T) {
	amount, err := domainAmount.NewPositive(10)
	require.NoError(t, err)
	require.Equal(t, int64(10), amount.Value())

	_, err = domainAmount.NewPositive(-10)
	require.Error(t, err)

	_, err = domainAmount.NewPositive(0)
	require.Error(t, err)
}

func TestAmount_Add(t *testing.T) {
	amount1 := must.Must(domainAmount.New(10))
	amount2 := must.Must(domainAmount.NewPositive(20))

	amount, err := amount1.Add(amount2)
	require.NoError(t, err)
	require.Equal(t, int64(30), amount.Value())
}

func TestAmount_Add_Overflow(t *testing.T) {
	amount1 := must.Must(domainAmount.New(math.MaxInt64 - 1))
	amount2 := must.Must(domainAmount.NewPositive(2))

	_, err := amount1.Add(amount2)
	require.ErrorIs(t, err, domainAmount.ErrIntegerOverflow)
}

// TestAmount_Sub
func TestAmount_Sub(t *testing.T) {
	amount1 := must.Must(domainAmount.New(10))
	amount2 := must.Must(domainAmount.NewPositive(20))

	amount, err := amount1.Sub(amount2)
	require.ErrorIs(t, err, domainAmount.ErrInsufficientAmount)

	amount1 = must.Must(domainAmount.New(30))
	amount, err = amount1.Sub(amount2)
	require.NoError(t, err)
	require.Equal(t, int64(10), amount.Value())
}

func TestAmount_LessThan(t *testing.T) {
	amount1 := must.Must(domainAmount.New(10))
	amount2 := must.Must(domainAmount.New(20))

	require.True(t, amount1.LessThan(amount2))

	amount1 = must.Must(domainAmount.New(20))
	amount2 = must.Must(domainAmount.New(10))

	require.False(t, amount1.LessThan(amount2))
}

func TestAmount_Zero(t *testing.T) {
	amount := domainAmount.Zero()
	require.Equal(t, int64(0), amount.Value())
}
