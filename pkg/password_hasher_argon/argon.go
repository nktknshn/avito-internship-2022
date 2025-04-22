package password_hasher_argon

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	saltLength = 16        // длина соли в байтах
	timeCost   = 1         // количество итераций
	memoryCost = 64 * 1024 // объем памяти в килобайтах (примерно 64 MB)
	threads    = 4         // число потоков
	keyLength  = 32        // длина генерируемого хеша в байтах
	// Формат строки хеша: содержит версию, параметры, соль и сам хеш.
	hashFormat        = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s"
	saltHashSeparator = "$"
)

type hashedPassword struct {
	version     uint32
	memory      uint32
	timeCost    uint32
	parallelism uint8
	salt        []byte
	hash        []byte
}

func (hp hashedPassword) String() string {
	b64salt := base64.RawStdEncoding.EncodeToString(hp.salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hp.hash)
	return fmt.Sprintf(hashFormat, hp.version, hp.memory, hp.timeCost, hp.parallelism, b64salt+saltHashSeparator+b64hash)
}

func (hp hashedPassword) StringBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(hp.String()))
}

// hashPassword хеширует пароль с использованием Argon2id и возвращает строку,
// содержащую все необходимые параметры для проверки, соль и хеш.
func hashPassword(password string, salt []byte) (*hashedPassword, error) {
	var err error
	if salt == nil {
		salt, err = generateRandomBytes(saltLength)
		if err != nil {
			return nil, err
		}
	}

	// Вычисление хеша с помощью Argon2id
	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, threads, keyLength)

	// Формирование итоговой строки с параметрами, солью и хешем
	return &hashedPassword{
		version:     argon2.Version,
		memory:      memoryCost,
		timeCost:    timeCost,
		parallelism: threads,
		salt:        salt,
		hash:        hash,
	}, nil
}

// hashPasswordToBase64 хеширует пароль и возвращает его в виде строки, готовой для сохранения в базе данных.
func hashPasswordToBase64(password string) (string, error) {
	hash, err := hashPassword(password, nil)
	if err != nil {
		return "", err
	}
	b64hash := hash.StringBase64()
	// Не требуется дополнительное кодирование — хеш уже содержит salt и закодирован в нужном формате.
	return b64hash, nil
}

// verifyPassword проверяет, соответствует ли указанный пароль сохраненному хешу.
func verifyPasswordBase64(password string, hash string) (bool, error) {
	// Берем строковое представление сохраненного хеша
	argon2StringBytes, err := base64.StdEncoding.DecodeString(hash)

	if err != nil {
		return false, err
	}

	// Извлекаем из него соль (и игнорируем сам хеш, так как нам он понадобится для сравнения)
	hp, err := parseArgon2String(string(argon2StringBytes))
	if err != nil {
		return false, err
	}

	computedHash, err := hashPassword(password, hp.salt)
	if err != nil {
		return false, err
	}

	// Использование константного сравнения для предотвращения утечки времени
	if subtle.ConstantTimeCompare(computedHash.hash, hp.hash) == 1 {
		return true, nil
	}
	return false, nil
}

// parseArgon2String разбирает строку, содержащую параметры, соль и хеш, и возвращает их.
func parseArgon2String(encoded string) (*hashedPassword, error) {
	var saltAndHash string
	var hp hashedPassword
	_, err := fmt.Sscanf(encoded, hashFormat, &hp.version, &hp.memory, &hp.timeCost, &hp.parallelism, &saltAndHash)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(saltAndHash, "$")

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
