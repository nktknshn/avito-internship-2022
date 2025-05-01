package errors

import "errors"

type UseCaseError struct {
	message string
}

func New(message string) UseCaseError {
	return UseCaseError{message: message}
}

func (e UseCaseError) Error() string {
	return e.message
}

func IsUseCaseError(err error) bool {
	var ue UseCaseError
	return errors.As(err, &ue)
}
