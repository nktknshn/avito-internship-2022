package auth

import (
	"context"
	"slices"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	authUsecase     authUsecase
	accessibleRoles map[string][]domainAuth.AuthUserRole
}

type authUsecase interface {
	Handle(ctx context.Context, in auth_validate_token.In) (auth_validate_token.Out, error)
}

func NewAuthInterceptor(authUsecase authUsecase, accessibleRoles map[string][]domainAuth.AuthUserRole) *AuthInterceptor {
	if authUsecase == nil {
		panic("authUsecase is nil")
	}

	if accessibleRoles == nil {
		panic("accessibleRoles is nil")
	}

	for method, roles := range accessibleRoles {
		if len(roles) == 0 {
			panic("roles is empty for method: " + method)
		}
	}

	return &AuthInterceptor{authUsecase: authUsecase, accessibleRoles: accessibleRoles}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing auth token")
		}

		token := md.Get("authorization")

		if len(token) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing auth token")
		}

		in, err := auth_validate_token.NewInFromValues(token[0])
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid auth token")
		}

		out, err := i.authUsecase.Handle(ctx, in)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid auth token")
		}

		accessibleRoles, ok := i.accessibleRoles[info.FullMethod]

		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing accessible roles")
		}

		if !slices.Contains(accessibleRoles, out.Role) {
			return nil, status.Error(codes.PermissionDenied, "missing permission")
		}

		return handler(ctx, req)
	}
}
