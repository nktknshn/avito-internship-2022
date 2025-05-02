package errors

import "errors"

type UseCaseError struct {
	cause   error
	message string
}

func New(message string) UseCaseError {
	return UseCaseError{message: message}
}

func (e UseCaseError) WithCause(cause error) UseCaseError {
	e.cause = cause
	return e
}

func (e UseCaseError) Cause() error {
	return e.cause
}

func (e UseCaseError) Error() string {
	return e.message
}

func IsUseCaseError(err error) bool {
	var ue UseCaseError
	return errors.As(err, &ue)
}
