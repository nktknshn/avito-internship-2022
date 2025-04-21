package handlers_auth

import (
	"context"
	"errors"
	"net/http"
	"slices"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	ergo "github.com/nktknshn/go-ergo-handler"
)

var (
	ErrUserNotAllowed = errors.New("user is not allowed")
)

type RoleCheckerParser struct {
	roles []domain.AuthUserRole
}

func NewRoleChecker(roles ...domain.AuthUserRole) *RoleCheckerParser {
	return &RoleCheckerParser{roles: roles}
}

type attachedRoleChecker struct {
	Roles []domain.AuthUserRole
	auth  *AttachedAuthParser
}

func (r *RoleCheckerParser) Attach(auth *AttachedAuthParser, builder ergo.ParserAdder) *attachedRoleChecker {
	attached := &attachedRoleChecker{r.roles, auth}
	builder.AddParser(attached)
	return attached
}

func (at *attachedRoleChecker) ParseRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, error) {
	user := at.auth.GetContext(ctx)
	if !slices.Contains(at.Roles, (*user).GetAuthUserRole()) {
		return ctx, ergo.NewError(http.StatusForbidden, ErrUserNotAllowed)
	}
	return ctx, nil
}
