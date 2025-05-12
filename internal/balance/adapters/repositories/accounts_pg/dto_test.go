package accounts_pg

import (
	"testing"

	"github.com/stretchr/testify/require"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

func TestFromAccountDTO_StripDomainError(t *testing.T) {
	acc, err := fromAccountDTO(&accountDTO{})
	require.Error(t, err)
	require.Nil(t, acc)
	require.False(t, domainError.IsDomainError(err))
}
