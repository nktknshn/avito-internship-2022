//nolint:testpackage
package decorator

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	commonErrors "github.com/nktknshn/avito-internship-2022/internal/common/errors"
)

type useCase0 struct{}

func (u *useCase0) Handle(ctx context.Context, in int) error {
	panic("panic error")
}
func (u *useCase0) GetName() string {
	return "useCase0"
}

type useCase1 struct{}

func (u *useCase1) Handle(ctx context.Context, in int) (int, error) {
	panic("panic error")
}

func (u *useCase1) GetName() string {
	return "useCase1"
}

type mockRecoverHandler struct {
	mock.Mock
}

func (m *mockRecoverHandler) Handle(ctx context.Context, err error) (errRecovered error) {
	args := m.Called(ctx, err)
	return args.Error(0)
}

func TestDecorator0Recover(t *testing.T) {
	h := &mockRecoverHandler{}
	errRecovered := errors.New("error recovered")
	h.On("Handle", mock.Anything, commonErrors.NewErrPanic("panic error")).Return(errRecovered)
	dec := Decorator0Recover[int]{
		base:           &useCase0{},
		recoverHandler: h.Handle,
	}

	require.NotPanics(t, func() {
		err := dec.Handle(t.Context(), 1)
		require.ErrorIs(t, err, errRecovered)
	})

	h.AssertExpectations(t)
}

func TestDecorator1Recover(t *testing.T) {
	h := &mockRecoverHandler{}
	errRecovered := errors.New("error recovered")
	h.On("Handle", mock.Anything, commonErrors.NewErrPanic("panic error")).Return(errRecovered)
	dec := Decorator1Recover[int, int]{
		base:           &useCase1{},
		recoverHandler: h.Handle,
	}
	require.NotPanics(t, func() {
		res, err := dec.Handle(t.Context(), 1)
		require.ErrorIs(t, err, errRecovered)
		require.Equal(t, 0, res)
	})

	h.AssertExpectations(t)
}

func TestError(t *testing.T) {
	err := commonErrors.NewErrPanic("panic error")
	require.Equal(t, "panic: panic error", err.Error())
	require.NoError(t, err.Unwrap())
}

func TestErrorWrap(t *testing.T) {
	errToWrap := errors.New("error to wrap")
	err := commonErrors.NewErrPanic(errToWrap)
	require.Equal(t, "panic: error to wrap", err.Error())
	require.Equal(t, errToWrap, err.Unwrap())
	require.ErrorIs(t, err, errToWrap)
}
