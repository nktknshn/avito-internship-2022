package errors_test

import (
	"testing"

	useCaseError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestUseCaseError(t *testing.T) {
	causeError := errors.New("cause error")

	de := useCaseError.New("test use case error")
	deWithCause := de.WithCause(causeError)

	require.Equal(t, "test use case error", deWithCause.Error())
	require.Equal(t, "cause error", errors.Cause(deWithCause).Error())

	require.ErrorIs(t, deWithCause, de)
	require.ErrorIs(t, deWithCause, causeError)
	require.True(t, useCaseError.IsUseCaseError(deWithCause))
	require.True(t, errors.Is(errors.Cause(deWithCause), causeError))
}
