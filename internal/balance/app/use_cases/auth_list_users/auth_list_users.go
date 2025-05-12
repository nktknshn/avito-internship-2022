package auth_list_users

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

type In struct {
}

type Out struct {
	Users []OutUser
}

type OutUser struct {
	ID       domainAuth.AuthUserID
	Username domainAuth.AuthUserUsername
	Role     domainAuth.AuthUserRole
}

func newOutUser(user *domainAuth.AuthUser) OutUser {
	return OutUser{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}

type authRepo interface {
	ListUsers(ctx context.Context) ([]*domainAuth.AuthUser, error)
}

type AuthListUsersUseCase struct {
	authRepo authRepo
}

func New(authRepo authRepo) *AuthListUsersUseCase {
	if authRepo == nil {
		panic("authRepo is nil")
	}

	return &AuthListUsersUseCase{
		authRepo: authRepo,
	}
}

func (u *AuthListUsersUseCase) Handle(ctx context.Context, _ In) (Out, error) {
	users, err := u.authRepo.ListUsers(ctx)

	if err != nil {
		return Out{}, err
	}

	out := Out{
		Users: make([]OutUser, len(users)),
	}

	for i, user := range users {
		out.Users[i] = newOutUser(user)
	}

	return out, nil
}

func (u *AuthListUsersUseCase) GetName() string {
	return use_cases.NameAuthListUsers
}
