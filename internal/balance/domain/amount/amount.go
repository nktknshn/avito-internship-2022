package amount

import (
	"math"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

// Неотрициательное кол-во копеек
type Amount struct {
	amount int64
}

func Zero() Amount {
	return Amount{amount: 0}
}

var (
	ErrInvalidAmount         = domainError.New("invalid amount")
	ErrInvalidPositiveAmount = domainError.New("invalid positive amount")
	ErrInsufficientAmount    = domainError.New("insufficient amount")
	ErrIntegerOverflow       = domainError.New("integer overflow")
)

func New(amount int64) (Amount, error) {
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

func (a Amount) Add(b AmountPositive) (Amount, error) {

	if a.Value() > math.MaxInt64-b.Value() {
		return Amount{}, ErrIntegerOverflow
	}

	return Amount{amount: a.Value() + b.Value()}, nil
}

func (a Amount) Sub(b AmountPositive) (Amount, error) {
	if a.Value() < b.Value() {
		return Amount{}, ErrInsufficientAmount
	}
	return Amount{amount: a.Value() - b.Value()}, nil
}
