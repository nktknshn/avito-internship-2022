package domain

import "errors"

var (
	ErrTransactionAlreadyExists  = errors.New("transaction already exists")
	ErrTransactionNotFound       = errors.New("transaction not found")
	ErrTransactionAmountMismatch = errors.New("transaction amount mismatch")
)
