package domain

import "errors"

var (
	ErrAccountNotFound                 = errors.New("account not found")
	ErrTransactionAlreadyExists        = errors.New("transaction already exists")
	ErrTransactionNotFound             = errors.New("transaction not found")
	ErrTransactionAmountMismatch       = errors.New("transaction amount mismatch")
	ErrInvalidAccountTransactionID     = errors.New("invalid account transaction id")
	ErrInvalidAccountTransactionStatus = errors.New("invalid account transaction status")
	ErrInvalidProductID                = errors.New("invalid product id")
)
