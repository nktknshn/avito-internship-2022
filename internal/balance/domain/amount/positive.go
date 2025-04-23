package amount

// Кол-во копеек больше 0
type AmountPositive struct {
	amount int64
}

func NewPositive(amount int64) (AmountPositive, error) {
	if amount <= 0 {
		return AmountPositive{}, ErrInvalidPositiveAmount
	}
	return AmountPositive{amount: amount}, nil
}

func (a AmountPositive) Value() int64 {
	return a.amount
}
