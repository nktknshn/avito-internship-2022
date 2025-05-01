package report_transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCursorEmpty(t *testing.T) {
	assert.True(t, CursorEmpty == "")
}
