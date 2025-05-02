package password_hasher_argon

type Hasher struct {
}

func New() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(password string) (string, error) {
	return hashPasswordToBase64(password)
}

func (h *Hasher) Verify(password, hashBase64 string) (bool, error) {
	return verifyPasswordBase64(password, hashBase64)
}
