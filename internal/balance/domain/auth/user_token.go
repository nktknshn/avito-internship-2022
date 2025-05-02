package auth

type AuthUserToken string

func NewAuthToken(token string) (AuthUserToken, error) {
	return AuthUserToken(token), nil
}

func (t AuthUserToken) String() string {
	return string(t)
}

type AuthUserTokenClaims struct {
	AuthUserID   int64
	AuthUserRole string
}
