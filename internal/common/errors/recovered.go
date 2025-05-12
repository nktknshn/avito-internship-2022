package errors

import (
	"errors"
	"fmt"
)

type ErrPanic struct {
	Arg any
	err error
}

func NewErrPanic(arg any) *ErrPanic {
	if e, ok := arg.(error); ok {
		return &ErrPanic{Arg: arg, err: e}
	}
	return &ErrPanic{Arg: arg, err: nil}
}

func (p *ErrPanic) Error() string {
	switch e := p.Arg.(type) {
	case error:
		return fmt.Sprintf("panic: %v", e.Error())
	default:
		return fmt.Sprintf("panic: %v", e)
	}
}

func (p *ErrPanic) Unwrap() error {
	return p.err
}

func IsErrPanic(err error) bool {
	var p *ErrPanic
	return errors.As(err, &p)
}
