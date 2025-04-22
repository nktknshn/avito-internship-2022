package password_hasher_argon

type Hasher struct {
}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(password string) (string, error) {
	return hashPasswordToBase64(password)
}

func (h *Hasher) Verify(password, hash string) (bool, error) {
	return verifyPasswordBase64(password, hash)
}
