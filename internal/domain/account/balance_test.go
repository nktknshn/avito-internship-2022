package account_test

import (
	"testing"

	"github.com/nktknshn/avito-internship-2022/internal/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/domain/amount"
	"github.com/stretchr/testify/require"
)

// TestAccountBalance_Equality проверяет, что AccountBalance является значением, а не объектом
func TestAccountBalance_Equality(t *testing.T) {
	av, err := amount.NewAmount(10)
	require.NoError(t, err)
	re, err := amount.NewAmount(100)
	require.NoError(t, err)

	b1, err := account.NewAccountBalance(av, re)
	require.NoError(t, err)

	b2, err := account.NewAccountBalance(av, re)
	require.NoError(t, err)

	if b1 != b2 {
		t.Fatalf("AccountBalance is not a real value object")
	}
}
