package transactions_pg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReportTransactionsCursor_Error_InvalidCursor(t *testing.T) {
	cursor, err := unmarshalCursor("invalid")
	require.Error(t, err)
	require.Nil(t, cursor)
}

func TestReportTransactionsCursor_EmptyCursor(t *testing.T) {
	cursor, err := unmarshalCursor("")
	require.NoError(t, err)
	require.Equal(t, cursor, emptyCursor)
}
