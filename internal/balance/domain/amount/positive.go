package amount

// Положительное кол-во копеек
type AmountPositive int64

func NewAmountPositive(amount int64) (AmountPositive, error) {
	if amount <= 0 {
		return 0, ErrInvalidPositiveAmount
	}
	return AmountPositive(amount), nil
}

func (a AmountPositive) Value() int64 {
	return int64(a)
}
