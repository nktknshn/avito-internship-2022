package password_hasher

type Hasher interface {
	Hash(password string) (string, error)
}

type HashVerifier interface {
	Verify(password, hash string) (bool, error)
}

type HasherVerifier interface {
	Hasher
	HashVerifier
}
