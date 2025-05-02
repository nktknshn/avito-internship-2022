package errors

import (
	"errors"
)

type domainErrorWithCause struct {
	domainError DomainError
	cause       error
	error       error
}

func (e domainErrorWithCause) Error() string {
	return e.domainError.Error()
}

func (e domainErrorWithCause) Unwrap() error {
	return e.error
}

func (e domainErrorWithCause) Cause() error {
	return e.cause
}

type DomainError struct {
	message string
}

func New(message string) DomainError {
	return DomainError{message: message}
}

func (e DomainError) Error() string {
	return e.message
}

// Скрывает cause из Error(). Причина остается доступна через errors.Cause(err).
// Обе ошибки матчатся через errors.Is
func (e DomainError) WithCause(cause error) error {
	return domainErrorWithCause{
		domainError: e,
		cause:       cause,
		error:       errors.Join(e, cause),
	}
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
