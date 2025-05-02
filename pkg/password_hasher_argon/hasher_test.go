package password_hasher_argon

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	// сгенерировано https://argon2.online/
	hashFormatString = "$argon2id$v=19$m=32,t=2,p=4$TG5Vd0RUNHVPQjk2MGpwag$/XaTBO+VnMSf8lsQ8NxDrw"
	password         = "password123"
)

func TestHasher_Hash(t *testing.T) {
	ok, err := verifyPasswordFormatString("password123", hashFormatString)
	require.NoError(t, err)
	require.True(t, ok)
}

func TestHasher_ParseArgon2String(t *testing.T) {
	hp, err := parseArgon2FormatString(hashFormatString)
	require.NoError(t, err)
	require.Equal(t, uint32(19), hp.version)
	require.Equal(t, uint32(32), hp.memory)
	require.Equal(t, uint32(2), hp.timeCost)
	require.Equal(t, uint8(4), hp.parallelism)
	require.Equal(t, uint32(16), hp.keyLength)
	require.Equal(t, uint32(16), hp.saltLength)
}

func TestHasher_HashAndVerify(t *testing.T) {
	hp, err := hashPasswordDefault(password, nil)
	require.NoError(t, err)
	ok, err := verifyPasswordFormatString(password, hp.FormatString())
	require.NoError(t, err)
	require.True(t, ok)
}
