package transaction

import domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"

var (
	ErrTransactionAlreadyExists  = domainError.New("transaction already exists")
	ErrTransactionAlreadyPaid    = domainError.New("transaction is already paid")
	ErrTransactionNotFound       = domainError.New("transaction not found")
	ErrTransactionAmountMismatch = domainError.New("transaction amount mismatch")
	ErrTransactionStatusMismatch = domainError.New("transaction status mismatch")
)
