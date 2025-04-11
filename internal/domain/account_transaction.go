package domain

type AccountTransactionSpendID int64

func (id AccountTransactionSpendID) Value() int64 {
	return int64(id)
}

func NewAccountTransactionSpendID(id int64) (AccountTransactionSpendID, error) {
	if id < 0 {
		return 0, ErrInvalidAccountTransactionID
	}
	return AccountTransactionSpendID(id), nil
}

type AccountTransactionDepositID int64

func (id AccountTransactionDepositID) Value() int64 {
	return int64(id)
}

func NewAccountTransactionDepositID(id int64) (AccountTransactionDepositID, error) {
	if id < 0 {
		return 0, ErrInvalidAccountTransactionID
	}
	return AccountTransactionDepositID(id), nil
}
