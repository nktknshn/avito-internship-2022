package amount

import "errors"

// Неотрициательное кол-во копеек
type Amount int64

var (
	ErrInvalidAmount         = errors.New("invalid amount")
	ErrInvalidPositiveAmount = errors.New("invalid positive amount")
	ErrInsufficientAmount    = errors.New("insufficient amount")
)

func NewAmount(amount int64) (Amount, error) {
	if amount < 0 {
		return 0, ErrInvalidAmount
	}

	return Amount(amount), nil
}

func (a Amount) Value() int64 {
	return int64(a)
}

func (a Amount) LessThan(b Amount) bool {
	return a.Value() < b.Value()
}

func (a Amount) LessThanPositive(b AmountPositive) bool {
	return a.Value() < b.Value()
}

func (a Amount) Add(b AmountPositive) Amount {
	return Amount(a.Value() + b.Value())
}

func (a Amount) Sub(b AmountPositive) (Amount, error) {
	if a.Value() < b.Value() {
		return 0, ErrInsufficientAmount
	}
	return Amount(a.Value() - b.Value()), nil
}
