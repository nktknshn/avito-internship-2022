package http

import (
	"context"

	openapi "github.com/nktknshn/avito-internship-2022-bench/api"
)

var username = "admin"
var password = "admin1234"

var host = "localhost"

func getConfig(host string) *openapi.Configuration {
	conf := openapi.NewConfiguration()
	conf.Servers = openapi.ServerConfigurations{
		{
			URL:         host,
			Description: "No description provided",
		},
	}
	return conf
}

func getClientOpenAPI(host string) (*openapi.APIClient, error) {
	conf := getConfig(host)
	client := openapi.NewAPIClient(conf)
	return client, nil
}

func getClientOpenAPIAuthorized(host string) (*openapi.APIClient, error) {
	client, err := getClientOpenAPI(host)

	if err != nil {
		return nil, err
	}

	req := client.SigninAPI.SignIn(context.Background())

	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersSigninRequestBody{
		Username: &username,
		Password: &password,
	})

	resp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}

	token := resp.Result.Token

	conf := getConfig(host)
	conf.DefaultHeader = map[string]string{
		"Authorization": "Bearer " + *token,
	}

	authClient := openapi.NewAPIClient(conf)

	return authClient, nil
}
