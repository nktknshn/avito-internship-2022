package transactions_pg

import (
	"math/rand"
	"time"

	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

type GenerateTransactionsDepositParams struct {
	TimeMin time.Time
	TimeMax time.Time

	AmountMin int64
	AmountMax int64
}

func GenerateTransactionsDeposit(
	accountID domainAccount.AccountID,
	userID domain.UserID,
	count int,
	params GenerateTransactionsDepositParams,
) []domainTransaction.TransactionDeposit {

	transactions := make([]domainTransaction.TransactionDeposit, count)

	for i := 0; i < count; i++ {
		amount, err := domainAmount.NewPositive(rand.Int63n(params.AmountMax-params.AmountMin) + params.AmountMin)
		if err != nil {
			panic(err)
		}
		transactions[i] = domainTransaction.TransactionDeposit{
			AccountID:     accountID,
			UserID:        userID,
			CreatedAt:     time.Now().Add(time.Duration(i) * time.Second),
			UpdatedAt:     time.Now().Add(time.Duration(i+1) * time.Second),
			DepositSource: "",
			Status:        domainTransaction.TransactionDepositStatusConfirmed,
			Amount:        amount,
		}
	}

	return transactions
}

type GenerateTransactionsSpendParams struct {
	TimeMin time.Time
	TimeMax time.Time

	AmountMin int64
	AmountMax int64
}

func GenerateTransactionsSpend(
	accountID domainAccount.AccountID,
	userID domain.UserID,
	count int,
	params GenerateTransactionsSpendParams,
) []domainTransaction.TransactionSpend {

	transactions := make([]domainTransaction.TransactionSpend, count)

	for i := range transactions {
		amount, err := domainAmount.NewPositive(rand.Int63n(params.AmountMax-params.AmountMin) + params.AmountMin)
		if err != nil {
			panic(err)
		}

		status := domainTransaction.TransactionSpendStatusConfirmed

		if i%2 == 0 {
			status = domainTransaction.TransactionSpendStatusCanceled
		}
		if i%3 == 0 {
			status = domainTransaction.TransactionSpendStatusReserved
		}

		productID := must.Must(domainProduct.NewProductID(int64(i) + 10))

		transactions[i] = domainTransaction.TransactionSpend{
			AccountID: accountID,
			UserID:    userID,
			Amount:    amount,
			OrderID:   must.Must(domain.NewOrderID(int64(i) + 10)),
			Status:    status,
			ProductID: productID,
			CreatedAt: time.Now().Add(time.Duration(i) * time.Second),
			UpdatedAt: time.Now().Add(time.Duration(i+1) * time.Second),
		}
	}

	return transactions
}

type GenerateTransactionsTransferParams struct {
	TimeMin time.Time
	TimeMax time.Time

	AmountMin int64
	AmountMax int64
}

func GenerateTransactionsTransfer(
	fromAccountID domainAccount.AccountID,
	toAccountID domainAccount.AccountID,

	count int,
	params GenerateTransactionsTransferParams,
) []domainTransaction.TransactionTransfer {

	transactions := make([]domainTransaction.TransactionTransfer, count)

	for i := range transactions {
		amount, err := domainAmount.NewPositive(rand.Int63n(params.AmountMax-params.AmountMin) + params.AmountMin)
		if err != nil {
			panic(err)
		}

		transactions[i] = domainTransaction.TransactionTransfer{
			FromAccountID: fromAccountID,
			ToAccountID:   toAccountID,
			Amount:        amount,
			Status:        domainTransaction.TransactionTransferStatusConfirmed,
			CreatedAt:     time.Now().Add(time.Duration(i) * time.Second),
			UpdatedAt:     time.Now().Add(time.Duration(i+1) * time.Second),
		}
	}

	return transactions
}
