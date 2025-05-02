package auth_validate_token

type In struct {
	token string
}

func NewInFromValues(token string) (In, error) {
	if token == "" {
		return In{}, ErrEmptyToken
	}
	return In{
		token: token,
	}, nil
}
