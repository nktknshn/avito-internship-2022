package errors_test

import (
	"errors"
	"testing"

	commonErrors "github.com/nktknshn/avito-internship-2022/internal/common/errors"
	"github.com/stretchr/testify/require"
)

func TestIsErrPanic(t *testing.T) {
	err := commonErrors.NewErrPanic("panic error")
	require.True(t, commonErrors.IsErrPanic(err))

	err2 := commonErrors.NewErrPanic(errors.New("error"))
	require.True(t, commonErrors.IsErrPanic(err2))

	err3 := errors.New("error")
	require.False(t, commonErrors.IsErrPanic(err3))

}
