#!/bin/sh

go run github.com/swaggo/swag/v2/cmd/swag init --v3.1 --parseDependency --parseInternal -g internal/balance/adapters/http/api.go -o api/openapi
