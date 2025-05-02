package errors_test

import (
	"testing"

	"github.com/pkg/errors"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	"github.com/stretchr/testify/require"
)

func TestDomainError(t *testing.T) {
	causeError := errors.New("cause error")

	de := domainError.New("test domain error")
	deWithCause := de.WithCause(causeError)

	require.Equal(t, "test domain error", deWithCause.Error())
	require.Equal(t, "cause error", errors.Cause(deWithCause).Error())

	require.ErrorIs(t, deWithCause, de)
	require.ErrorIs(t, deWithCause, causeError)
	require.True(t, domainError.IsDomainError(deWithCause))
	require.True(t, errors.Is(errors.Cause(deWithCause), causeError))
}

func TestDomainError_Strip(t *testing.T) {
	de := domainError.New("test domain error")
	stripped := domainError.Strip(de)
	require.True(t, domainError.IsDomainError(de))
	require.False(t, domainError.IsDomainError(stripped))

	require.Equal(t, de.Error(), stripped.Error())
}
