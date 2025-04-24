package domain

import domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"

var (
	ErrInvalidOrderID = domainError.New("invalid order id")
)

type OrderID int64

func (o OrderID) Value() int64 {
	return int64(o)
}
func NewOrderID(id int64) (OrderID, error) {
	if id <= 0 {
		return 0, ErrInvalidOrderID
	}
	return OrderID(id), nil
}
