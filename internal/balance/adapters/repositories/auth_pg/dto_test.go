package auth_pg

import (
	"testing"

	"github.com/stretchr/testify/require"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

func TestFromAuthUserDTO_StripDomainError(t *testing.T) {
	authUser, err := fromAuthUserDTO(&authUserDTO{})
	require.Error(t, err)
	require.Nil(t, authUser)
	require.False(t, domainError.IsDomainError(err))
}
