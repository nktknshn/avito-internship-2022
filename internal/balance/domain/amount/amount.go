package amount

import "errors"

// Неотрициательное кол-во копеек
type Amount struct {
	amount int64
}

func Zero() Amount {
	return Amount{amount: 0}
}

var (
	ErrInvalidAmount         = errors.New("invalid amount")
	ErrInvalidPositiveAmount = errors.New("invalid positive amount")
	ErrInsufficientAmount    = errors.New("insufficient amount")
)

func NewAmount(amount int64) (Amount, error) {
	if amount < 0 {
		return Amount{}, ErrInvalidAmount
	}

	return Amount{amount: amount}, nil
}

func (a Amount) Value() int64 {
	return a.amount
}

func (a Amount) LessThan(b Amount) bool {
	return a.Value() < b.Value()
}

func (a Amount) LessThanPositive(b AmountPositive) bool {
	return a.Value() < b.Value()
}

func (a Amount) Add(b AmountPositive) Amount {
	return Amount{amount: a.Value() + b.Value()}
}

func (a Amount) Sub(b AmountPositive) (Amount, error) {
	if a.Value() < b.Value() {
		return Amount{}, ErrInsufficientAmount
	}
	return Amount{amount: a.Value() - b.Value()}, nil
}
