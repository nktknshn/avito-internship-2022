package cli

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_list_users"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
)

type CliAdapter struct {
	app *app.Application
}

func NewCliAdapter(app *app.Application) *CliAdapter {
	return &CliAdapter{app: app}
}

func (a *CliAdapter) SignUp(ctx context.Context, username string, password string, role string) error {
	in, err := auth_signup.NewInFromValues(username, password, role)
	if err != nil {
		return err
	}
	err = a.app.AuthSignup.Handle(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

func (a *CliAdapter) SignIn(ctx context.Context, username string, password string) (string, error) {
	in, err := auth_signin.NewInFromValues(username, password)
	if err != nil {
		return "", err
	}
	out, err := a.app.AuthSignin.Handle(ctx, in)
	if err != nil {
		return "", err
	}
	return out.Token.String(), nil
}

func (a *CliAdapter) ListUsers(ctx context.Context) ([]auth_list_users.OutUser, error) {
	in := auth_list_users.In{}
	out, err := a.app.AuthListUsers.Handle(ctx, in)
	if err != nil {
		return nil, err
	}
	return out.Users, nil
}
