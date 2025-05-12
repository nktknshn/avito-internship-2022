package errors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	commonErrors "github.com/nktknshn/avito-internship-2022/internal/common/errors"
)

func TestIsErrPanic(t *testing.T) {
	err := commonErrors.NewErrPanic("panic error")
	require.True(t, commonErrors.IsPanicError(err))

	err2 := commonErrors.NewErrPanic(errors.New("error"))
	require.True(t, commonErrors.IsPanicError(err2))

	err3 := errors.New("error")
	require.False(t, commonErrors.IsPanicError(err3))

}
