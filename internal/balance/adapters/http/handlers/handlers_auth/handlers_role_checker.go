package handlers_auth

import (
	"context"
	"errors"
	"net/http"
	"slices"

	ergo "github.com/nktknshn/go-ergo-handler"

	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

var (
	ErrUserNotAllowed = errors.New("user is not allowed")
)

type RoleCheckerParser struct {
	roles []domainAuth.AuthUserRole
}

func NewRoleChecker(roles ...domainAuth.AuthUserRole) *RoleCheckerParser {
	return &RoleCheckerParser{roles: roles}
}

type AttachedRoleChecker struct {
	Roles []domainAuth.AuthUserRole
	auth  *AttachedAuthParser
}

func (r *RoleCheckerParser) Attach(auth *AttachedAuthParser, builder ergo.ParserAdder) *AttachedRoleChecker {
	attached := &AttachedRoleChecker{r.roles, auth}
	builder.AddParser(attached)
	return attached
}

func (at *AttachedRoleChecker) ParseRequest(ctx context.Context, _ http.ResponseWriter, _ *http.Request) (context.Context, error) {
	user := at.auth.GetContext(ctx)
	if !slices.Contains(at.Roles, user.GetRole()) {
		return ctx, ergo.NewError(http.StatusForbidden, ErrUserNotAllowed)
	}
	return ctx, nil
}
