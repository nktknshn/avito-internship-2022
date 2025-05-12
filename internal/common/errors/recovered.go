package errors

import (
	"errors"
	"fmt"
)

type PanicError struct {
	Arg any
	err error
}

func NewErrPanic(arg any) *PanicError {
	if e, ok := arg.(error); ok {
		return &PanicError{Arg: arg, err: e}
	}
	return &PanicError{Arg: arg, err: nil}
}

func (p *PanicError) Error() string {
	switch e := p.Arg.(type) {
	case error:
		return fmt.Sprintf("panic: %v", e.Error())
	default:
		return fmt.Sprintf("panic: %v", e)
	}
}

func (p *PanicError) Unwrap() error {
	return p.err
}

func IsPanicError(err error) bool {
	var p *PanicError
	return errors.As(err, &p)
}
