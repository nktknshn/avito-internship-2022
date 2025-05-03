package http

import (
	"context"
	"net/http"

	"github.com/nktknshn/avito-internship-2022-bench/client_http"
	"github.com/nktknshn/avito-internship-2022-bench/logger"
)

var username = "admin"
var password = "admin1234"
var host = "http://localhost:5050"

func getClient(host string) (*client_http.ClientWithResponses, error) {
	client, err := client_http.NewClientWithResponses(host)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getClientAuthorized(host string) (*client_http.ClientWithResponses, error) {
	client, err := getClient(host)
	if err != nil {
		return nil, err
	}

	respSignIn, err := client.SignInWithResponse(context.Background(), client_http.SignInJSONRequestBody{
		Username: &username,
		Password: &password,
	})

	if err != nil {
		logger.GetLogger().Error("Failed to sign in", "error", err)
		return nil, err
	}

	token := respSignIn.JSON200.Result.Token

	authClient, err := client_http.NewClientWithResponses(host, client_http.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+*token)
		return nil
	}))

	if err != nil {
		return nil, err
	}

	return authClient, nil
}
