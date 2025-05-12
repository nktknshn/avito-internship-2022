package password_hasher_argon

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	defaultSaltLength = 16        // длина соли в байтах
	defaultTimeCost   = 1         // количество итераций
	defaultMemoryCost = 64 * 1024 // объем памяти в килобайтах (примерно 64 MB)
	defaultThreads    = 4         // число потоков
	defaultKeyLength  = 32        // длина генерируемого хеша в байтах
	// Формат строки хеша: содержит версию, параметры, соль и сам хеш.
	hashFormat        = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s"
	saltHashSeparator = "$"
)

type hashParams struct {
	version     uint32
	memory      uint32
	timeCost    uint32
	parallelism uint8
	salt        []byte
	keyLength   uint32
	saltLength  uint32
}

type hashedPassword struct {
	hashParams
	hash []byte
}

func (hp hashedPassword) FormatString() string {
	b64salt := base64.RawStdEncoding.EncodeToString(hp.salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hp.hash)
	return fmt.Sprintf(hashFormat, hp.version, hp.memory, hp.timeCost, hp.parallelism, b64salt+saltHashSeparator+b64hash)
}

func (hp hashedPassword) String() string {
	return fmt.Sprintf(
		"version: %d, memory: %d, timeCost: %d, parallelism: %d, saltLength: %d, keyLength: %d, salt: %s, hash: %s",
		hp.version,
		hp.memory,
		hp.timeCost,
		hp.parallelism,
		hp.saltLength,
		hp.keyLength,
		hp.salt,
		hp.hash,
	)
}

func (hp hashedPassword) FormatStringBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(hp.FormatString()))
}

// hashPasswordDefault хеширует пароль с использованием Argon2id и возвращает строку,
// содержащую все необходимые параметры для проверки, соль и хеш.
func hashPasswordDefault(password string, salt []byte) (*hashedPassword, error) {
	var err error

	if salt == nil {
		salt, err = generateRandomBytes(defaultSaltLength)
		if err != nil {
			return nil, err
		}
	}

	// Вычисление хеша с помощью Argon2id
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultTimeCost,
		defaultMemoryCost,
		defaultThreads,
		defaultKeyLength,
	)

	// Формирование итоговой строки с параметрами, солью и хешем
	return &hashedPassword{
		hashParams: hashParams{
			version:     argon2.Version,
			memory:      defaultMemoryCost,
			timeCost:    defaultTimeCost,
			parallelism: defaultThreads,
			keyLength:   defaultKeyLength,
			saltLength:  defaultSaltLength,
			salt:        salt,
		},
		hash: hash,
	}, nil
}

func hashPassword(password string, params hashParams) *hashedPassword {
	hash := argon2.IDKey(
		[]byte(password),
		params.salt,
		params.timeCost,
		params.memory,
		params.parallelism,
		params.keyLength,
	)
	return &hashedPassword{
		hashParams: hashParams{
			version:     argon2.Version,
			memory:      params.memory,
			timeCost:    params.timeCost,
			parallelism: params.parallelism,
			salt:        params.salt,
			keyLength:   params.keyLength,
			saltLength:  params.saltLength,
		},
		hash: hash,
	}
}

// hashPasswordToBase64 хеширует пароль и возвращает его в виде строки, готовой для сохранения в базе данных.
func hashPasswordToBase64(password string) (string, error) {
	hash, err := hashPasswordDefault(password, nil)
	if err != nil {
		return "", err
	}
	return hash.FormatStringBase64(), nil
}

func verifyPasswordHashed(password string, hash *hashedPassword) bool {
	computedHash := hashPassword(password, hash.hashParams)

	return subtle.ConstantTimeCompare(computedHash.hash, hash.hash) == 1
}

func verifyPasswordFormatString(password string, formatString string) (bool, error) {
	hp, err := parseArgon2FormatString(formatString)

	if err != nil {
		return false, err
	}

	return verifyPasswordHashed(password, hp), nil
}

// verifyPassword проверяет, соответствует ли указанный пароль сохраненному хешу.
func verifyPasswordBase64(password string, hash string) (bool, error) {
	// Берем строковое представление сохраненного хеша
	argon2StringBytes, err := base64.StdEncoding.DecodeString(hash)

	if err != nil {
		return false, err
	}

	return verifyPasswordFormatString(password, string(argon2StringBytes))
}

// parseArgon2FormatString разбирает строку, содержащую параметры, соль и хеш, и возвращает их.
func parseArgon2FormatString(formatString string) (*hashedPassword, error) {
	var saltAndHash string
	var hp hashedPassword

	_, err := fmt.Sscanf(formatString, hashFormat, &hp.version, &hp.memory, &hp.timeCost, &hp.parallelism, &saltAndHash)

	if err != nil {
		return nil, err
	}

	parts := strings.Split(saltAndHash, "$")

	//nolint: mnd // ...
	if len(parts) != 2 {
		return nil, errors.New("invalid salt and hash")
	}

	b64salt := parts[0]
	b64hash := parts[1]

	salt, err := base64.RawStdEncoding.DecodeString(b64salt)
	if err != nil {
		return nil, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(b64hash)
	if err != nil {
		return nil, err
	}

	hp.salt = salt
	hp.hash = hash

	hashLen := len(hash)
	saltLen := len(salt)

	if hashLen > math.MaxUint32 || saltLen > math.MaxUint32 {
		return nil, errors.New("hash or salt length exceeds maximum allowed size")
	}

	hp.keyLength = uint32(hashLen)
	hp.saltLength = uint32(saltLen)

	return &hp, nil
}

// generateRandomBytes генерирует случайный набор байт указанной длины.
func generateRandomBytes(n int) ([]byte, error) {
	salt := make([]byte, n)
	// rand.Read пишет случайные байты непосредственно в выделенный буфер `salt`
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
