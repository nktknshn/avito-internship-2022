#!/bin/bash

set -e 

rm -rf coverage
mkdir coverage

go test ./internal/balance/... \
    -coverpkg github.com/nktknshn/avito-internship-2022/internal/balance/... \
    -coverprofile=coverage/coverage.txt

go tool cover -html=coverage/coverage.txt

rm -rf coverage
