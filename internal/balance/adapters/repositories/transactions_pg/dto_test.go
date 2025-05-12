package transactions_pg

import (
	"testing"

	"github.com/stretchr/testify/require"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

func TestFromTransactionDepositDTO_StripDomainError(t *testing.T) {
	transaction, err := fromTransactionDepositDTO(&transactionDepositDTO{})
	require.Error(t, err)
	require.Nil(t, transaction)
	require.False(t, domainError.IsDomainError(err))
}

func TestFromTransactionSpendDTO_StripDomainError(t *testing.T) {
	transaction, err := fromTransactionSpendDTO(&transactionSpendDTO{})
	require.Error(t, err)
	require.Nil(t, transaction)
	require.False(t, domainError.IsDomainError(err))
}

func TestFromTransactionTransferDTO_StripDomainError(t *testing.T) {
	transaction, err := fromTransactionTransferDTO(&transactionTransferDTO{})
	require.Error(t, err)
	require.Nil(t, transaction)
	require.False(t, domainError.IsDomainError(err))
}

func TestFromReportTransactionDTO_StripDomainError(t *testing.T) {
	transaction, err := fromReportTransactionDTO(&reportTransactionDTO{
		TransactionType: "deposit",
	})
	require.Error(t, err)
	require.Nil(t, transaction)
	require.False(t, domainError.IsDomainError(err))

	transaction, err = fromReportTransactionDTO(&reportTransactionDTO{
		TransactionType: "spend",
	})
	require.Error(t, err)
	require.Nil(t, transaction)
	require.False(t, domainError.IsDomainError(err))

	transaction, err = fromReportTransactionDTO(&reportTransactionDTO{
		TransactionType: "transfer",
	})
	require.Error(t, err)
	require.Nil(t, transaction)
	require.False(t, domainError.IsDomainError(err))

}
