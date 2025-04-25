package errors

import (
	"errors"
)

type DomainError struct {
	message string
}

func New(message string) DomainError {
	return DomainError{message: message}
}

func (e DomainError) Error() string {
	return e.message
}

func IsDomainError(err error) bool {
	var de DomainError
	return errors.As(err, &de)
}
