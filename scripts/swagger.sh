#!/bin/sh

go run github.com/swaggo/swag/cmd/swag init --parseDependency --parseInternal -g internal/balance/adapters/http/api.go -o api/openapi
