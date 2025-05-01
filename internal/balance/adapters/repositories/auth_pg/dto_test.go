package auth_pg

import (
	"testing"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	"github.com/stretchr/testify/require"
)

func TestFromAuthUserDTO_StripDomainError(t *testing.T) {
	authUser, err := fromAuthUserDTO(&authUserDTO{})
	require.Error(t, err)
	require.Nil(t, authUser)
	require.False(t, domainError.IsDomainError(err))
}
