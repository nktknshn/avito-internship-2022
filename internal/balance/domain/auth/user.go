package auth

import "errors"

type AuthUserID int64

func (id AuthUserID) Value() int64 {
	return int64(id)
}

const (
	minUsernameLength = 3
	minPasswordLength = 8
)

var (
	ErrInvalidAuthUserID = errors.New("invalid auth user id")
)

func NewAuthUserID(id int64) (AuthUserID, error) {
	if id < 0 {
		return 0, ErrInvalidAuthUserID
	}
	return AuthUserID(id), nil
}

type AuthUserUsername string

func (u AuthUserUsername) Value() string {
	return string(u)
}

var (
	ErrInvalidAuthUserUsernameTooShort = errors.New("invalid auth user username: too short")
)

func NewAuthUserUsername(username string) (AuthUserUsername, error) {
	if len(username) < minUsernameLength {
		return "", ErrInvalidAuthUserUsernameTooShort
	}
	return AuthUserUsername(username), nil
}

type AuthUserPassword string

func (p AuthUserPassword) String() string {
	return string(p)
}

var (
	ErrInvalidAuthUserPasswordTooShort = errors.New("invalid auth user password: too short")
)

func NewAuthUserPassword(password string) (AuthUserPassword, error) {
	if len(password) < minPasswordLength {
		return "", ErrInvalidAuthUserPasswordTooShort
	}
	return AuthUserPassword(password), nil
}

type AuthUserPasswordHash string

func (p AuthUserPasswordHash) Value() string {
	return string(p)
}

func NewAuthUserPasswordHash(passwordHash string) (AuthUserPasswordHash, error) {
	return AuthUserPasswordHash(passwordHash), nil
}

type AuthUser struct {
	ID           AuthUserID
	Username     AuthUserUsername
	PasswordHash AuthUserPasswordHash
	Role         AuthUserRole
}

func NewAuthUser(
	id AuthUserID,
	username AuthUserUsername,
	passwordHash AuthUserPasswordHash,
	role AuthUserRole,
) (*AuthUser, error) {
	return &AuthUser{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
	}, nil
}

func NewAuthUserFromValues(
	id int64,
	username string,
	passwordHash string,
	role string,
) (*AuthUser, error) {
	_id, err := NewAuthUserID(id)
	if err != nil {
		return nil, err
	}
	_username, err := NewAuthUserUsername(username)
	if err != nil {
		return nil, err
	}
	_passwordHash, err := NewAuthUserPasswordHash(passwordHash)
	if err != nil {
		return nil, err
	}
	_role, err := NewAuthUserRole(role)
	if err != nil {
		return nil, err
	}

	return NewAuthUser(_id, _username, _passwordHash, _role)
}
