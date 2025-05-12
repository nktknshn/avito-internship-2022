package errors

import "errors"

//nolint:errname // errname
type useCaseErrorWithCause struct {
	useCaseError UseCaseError
	cause        error
	error        error
}

func (e useCaseErrorWithCause) Error() string {
	return e.useCaseError.Error()
}

func (e useCaseErrorWithCause) Unwrap() error {
	return e.error
}

func (e useCaseErrorWithCause) Cause() error {
	return e.cause
}

type UseCaseError struct {
	message string
}

func New(message string) UseCaseError {
	return UseCaseError{message: message}
}

func (e UseCaseError) WithCause(cause error) error {
	return useCaseErrorWithCause{
		useCaseError: e,
		cause:        cause,
		error:        errors.Join(e, cause),
	}
}

func (e UseCaseError) Error() string {
	return e.message
}

func IsUseCaseError(err error) bool {
	var ue UseCaseError
	return errors.As(err, &ue)
}
