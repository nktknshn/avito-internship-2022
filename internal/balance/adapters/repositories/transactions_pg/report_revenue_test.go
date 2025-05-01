package transactions_pg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetMonthRangeUTC(t *testing.T) {
	type testCase struct {
		year  int
		month int
		t0    time.Time
		t1    time.Time
	}

	testCases := []testCase{
		{
			year:  2025,
			month: 4,
			t0:    time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			t1:    time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			year:  2025,
			month: 12,
			t0:    time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
			t1:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, testCase := range testCases {
		t0, t1 := getMonthRangeUTC(testCase.year, testCase.month)
		require.Equal(t, t0, testCase.t0)
		require.Equal(t, t1, testCase.t1)
		require.Equal(t, t0.Location(), time.UTC)
		require.Equal(t, t1.Location(), time.UTC)
	}
}
