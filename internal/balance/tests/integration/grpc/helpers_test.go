package grpc_test

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func withAuthToken(ctx context.Context, token string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"authorization": token,
	}))
}
