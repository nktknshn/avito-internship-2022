package report_transactions

import "time"

type Out struct {
	Transactions []OutTransaction `json:"transactions"`
	Cursor       Cursor           `json:"cursor"`
	HasMore      bool             `json:"has_more"`
}

type OutTransaction any

type OutTransactionSpend struct {
	ID           int64     `json:"id"`
	AccountID    int64     `json:"account_id"`
	OrderID      int64     `json:"order_id"`
	ProductID    int64     `json:"product_id"`
	ProductTitle string    `json:"product_title"`
	Amount       int64     `json:"amount"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OutTransactionDeposit struct {
	ID        int64     `json:"id"`
	Source    string    `json:"source"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OutTransactionTransfer struct {
	ID        int64     `json:"id"`
	From      int64     `json:"from"`
	To        int64     `json:"to"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
