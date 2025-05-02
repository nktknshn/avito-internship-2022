package errors

import (
	"errors"
)

type DomainError struct {
	cause   error
	message string
}

func New(message string) DomainError {
	return DomainError{message: message}
}

func (e DomainError) Error() string {
	return e.message
}

func (e DomainError) WithCause(cause error) DomainError {
	e.cause = cause
	return e
}

func (e DomainError) Cause() error {
	return e.cause
}

func IsDomainError(err error) bool {
	var de DomainError
	return errors.As(err, &de)
}

// Strip возвращает доменную ошибку в виде обычной ошибки с тем же сообщением
func Strip(err error) error {
	if !IsDomainError(err) {
		return err
	}
	return errors.New(err.Error())
}
