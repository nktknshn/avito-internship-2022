#!/bin/bash

set -e 

rm -rf coverage
mkdir coverage

DEFAULT_COVERAGE_PKG="github.com/nktknshn/avito-internship-2022/internal/balance/..."

COVERAGE_PKGS=()

for pkg in "$@"; do
    COVERAGE_PKGS+=("$pkg")
done

if [ ${#COVERAGE_PKGS[@]} -eq 0 ]; then
    COVERAGE_PKGS+=("$DEFAULT_COVERAGE_PKG")
fi

go test ./internal/balance/... \
    -coverpkg ${COVERAGE_PKGS[@]} \
    -coverprofile=coverage/coverage.txt

go tool cover -func=coverage/coverage.txt

rm -rf coverage