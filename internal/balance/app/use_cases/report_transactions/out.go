package report_transactions

import (
	"time"

	"github.com/google/uuid"
)

type Out struct {
	Transactions []OutTransaction `json:"transactions"`
	Cursor       Cursor           `json:"cursor"`
	HasMore      bool             `json:"has_more"`
}

type OutTransaction interface {
	isOutTransaction()
}

type OutTransactionSpend struct {
	ID           uuid.UUID `json:"id"`
	AccountID    int64     `json:"account_id"`
	OrderID      int64     `json:"order_id"`
	ProductID    int64     `json:"product_id"`
	ProductTitle string    `json:"product_title"`
	Amount       int64     `json:"amount"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (o *OutTransactionSpend) isOutTransaction() {}

type OutTransactionDeposit struct {
	ID        uuid.UUID `json:"id"`
	Source    string    `json:"source"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (o *OutTransactionDeposit) isOutTransaction() {}

type OutTransactionTransfer struct {
	ID        uuid.UUID `json:"id"`
	From      int64     `json:"from"`
	To        int64     `json:"to"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (o *OutTransactionTransfer) isOutTransaction() {}
