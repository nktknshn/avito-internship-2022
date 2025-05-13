package transaction

import domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"

var (
	ErrTransactionAlreadyPaid       = domainError.New("transaction is already paid")
	ErrTransactionAlreadyCanceled   = domainError.New("transaction is already canceled")
	ErrTransactionAlreadyReserved   = domainError.New("transaction is already reserved")
	ErrTransactionNotFound          = domainError.New("transaction not found")
	ErrTransactionAmountMismatch    = domainError.New("transaction amount mismatch")
	ErrTransactionStatusMismatch    = domainError.New("transaction status mismatch")
	ErrTransactionProductIDMismatch = domainError.New("transaction productID mismatch")
)
