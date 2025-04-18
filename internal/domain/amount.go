package domain

import "errors"

// Неотрициательное кол-во копеек
type Amount int64

var (
	InvalidAmount         = errors.New("invalid amount")
	InvalidPositiveAmount = errors.New("invalid positive amount")
)

func NewAmount(amount int64) (Amount, error) {
	if amount < 0 {
		return 0, InvalidAmount
	}

	return Amount(amount), nil
}

func (a Amount) Value() int64 {
	return int64(a)
}

func (a Amount) Add(b AmountPositive) Amount {
	return Amount(a.Value() + b.Value())
}

func (a Amount) Sub(b AmountPositive) Amount {
	return Amount(a.Value() - b.Value())
}
