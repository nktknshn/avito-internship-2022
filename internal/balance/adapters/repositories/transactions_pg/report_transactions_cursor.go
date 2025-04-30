package transactions_pg

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
)

type cursorAmount struct {
	ID     uuid.UUID `json:"id"`
	Amount uint64    `json:"amount"`
}

type cursorUpdatedAt struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type cursorUnion struct {
	Amount    *cursorAmount
	UpdatedAt *cursorUpdatedAt
}

func (c *cursorUnion) IsZero() bool {
	return c.Amount == nil && c.UpdatedAt == nil
}

func (c *cursorUnion) IsAmount() bool {
	return c.Amount != nil
}

func (c *cursorUnion) IsUpdatedAt() bool {
	return c.UpdatedAt != nil
}

// декодируем из base64 и парсим json
func unmarshalCursor(cursor report_transactions.Cursor) (*cursorUnion, error) {
	if cursor == report_transactions.CursorEmpty {
		return &cursorUnion{}, nil
	}

	cursorStr, ok := cursor.(string)
	if !ok {
		return nil, report_transactions.ErrCursorInvalid
	}

	cursorBytes, err := base64.StdEncoding.DecodeString(cursorStr)
	if err != nil {
		return nil, report_transactions.ErrCursorInvalid
	}

	if bytes.Contains(cursorBytes, []byte("amount")) {
		var cursorAmount cursorAmount
		err = json.Unmarshal(cursorBytes, &cursorAmount)
		if err != nil {
			return nil, report_transactions.ErrCursorInvalid
		}
		return &cursorUnion{Amount: &cursorAmount}, nil
	}

	if bytes.Contains(cursorBytes, []byte("updated_at")) {
		var cursorUpdatedAt cursorUpdatedAt
		err = json.Unmarshal(cursorBytes, &cursorUpdatedAt)
		if err != nil {
			return nil, report_transactions.ErrCursorInvalid
		}
		return &cursorUnion{UpdatedAt: &cursorUpdatedAt}, nil
	}

	return nil, report_transactions.ErrCursorInvalid
}
