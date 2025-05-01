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
	authUsecase   authUsecase
	methodToRoles map[string][]domainAuth.AuthUserRole
}

type authUsecase interface {
	Handle(ctx context.Context, in auth_validate_token.In) (auth_validate_token.Out, error)
}

func NewAuthInterceptor(authUsecase authUsecase, methodToRoles map[string][]domainAuth.AuthUserRole) *AuthInterceptor {
	if authUsecase == nil {
		panic("authUsecase is nil")
	}

	if methodToRoles == nil {
		panic("accessibleRoles is nil")
	}

	return &AuthInterceptor{authUsecase: authUsecase, methodToRoles: methodToRoles}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		accessibleRoles, ok := i.methodToRoles[info.FullMethod]

		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing accessible roles")
		}

		if len(accessibleRoles) == 0 {
			return handler(ctx, req)
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

		if !slices.Contains(accessibleRoles, out.Role) {
			return nil, status.Error(codes.PermissionDenied, "user has no permission")
		}

		return handler(ctx, req)
	}
}
