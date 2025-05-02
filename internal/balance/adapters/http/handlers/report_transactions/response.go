package report_transactions

import (
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
)

type responseBody struct {
	Transactions []any  `json:"transactions"`
	Cursor       string `json:"cursor"`
	HasMore      bool   `json:"has_more"`
}

type responseTransactionSpend struct {
	ID           string `json:"id"`
	AccountID    int64  `json:"account_id"`
	OrderID      int64  `json:"order_id"`
	ProductID    int64  `json:"product_id"`
	ProductTitle string `json:"product_title"`
	Amount       int64  `json:"amount"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type responseTransactionDeposit struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type responseTransactionTransfer struct {
	ID        string `json:"id"`
	From      int64  `json:"from"`
	To        int64  `json:"to"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func outToResult(out report_transactions.Out) *responseBody {
	result := responseBody{
		Transactions: make([]any, len(out.Transactions)),
		Cursor:       out.Cursor.Value(),
		HasMore:      out.HasMore,
	}

	for i, transaction := range out.Transactions {
		switch tr := transaction.(type) {
		case *report_transactions.OutTransactionSpend:
			result.Transactions[i] = responseTransactionSpend{
				ID:           tr.ID.Value().String(),
				AccountID:    tr.AccountID.Value(),
				OrderID:      tr.OrderID.Value(),
				ProductID:    tr.ProductID.Value(),
				ProductTitle: tr.ProductTitle.Value(),
				Amount:       tr.Amount.Value(),
				Status:       tr.Status.Value(),
				CreatedAt:    tr.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    tr.UpdatedAt.Format(time.RFC3339),
			}
		case *report_transactions.OutTransactionDeposit:
			result.Transactions[i] = responseTransactionDeposit{
				ID:        tr.ID.Value().String(),
				Source:    tr.Source.Value(),
				Amount:    tr.Amount.Value(),
				Status:    tr.Status.Value(),
				CreatedAt: tr.CreatedAt.Format(time.RFC3339),
				UpdatedAt: tr.UpdatedAt.Format(time.RFC3339),
			}
		case *report_transactions.OutTransactionTransfer:
			result.Transactions[i] = responseTransactionTransfer{
				ID:        tr.ID.Value().String(),
				From:      tr.From.Value(),
				To:        tr.To.Value(),
				Amount:    tr.Amount.Value(),
				Status:    tr.Status.Value(),
				CreatedAt: tr.CreatedAt.Format(time.RFC3339),
				UpdatedAt: tr.UpdatedAt.Format(time.RFC3339),
			}
		}
	}

	return &result
}
