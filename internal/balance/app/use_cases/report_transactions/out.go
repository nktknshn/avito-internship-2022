package report_transactions

import (
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type Out struct {
	Transactions []OutTransaction
	Cursor       Cursor
	HasMore      bool
}

type OutTransaction interface {
	isOutTransaction()
}

type OutTransactionSpend struct {
	ID           domainTransaction.TransactionSpendID
	AccountID    domainAccount.AccountID
	OrderID      domain.OrderID
	ProductID    domainProduct.ProductID
	ProductTitle domainProduct.ProductTitle
	Amount       domainAmount.AmountPositive
	Status       domainTransaction.TransactionSpendStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (o *OutTransactionSpend) isOutTransaction() {}

type OutTransactionDeposit struct {
	ID        domainTransaction.TransactionDepositID
	Source    domainTransaction.TransactionDepositSource
	Amount    domainAmount.AmountPositive
	Status    domainTransaction.TransactionDepositStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *OutTransactionDeposit) isOutTransaction() {}

type OutTransactionTransfer struct {
	ID        domainTransaction.TransactionTransferID
	From      domainAccount.AccountID
	To        domainAccount.AccountID
	Amount    domainAmount.AmountPositive
	Status    domainTransaction.TransactionTransferStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *OutTransactionTransfer) isOutTransaction() {}
