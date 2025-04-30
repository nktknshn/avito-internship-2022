#!/bin/sh

set -e

echo "Running start.sh"


# go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/cli --cfg-path config/balance/config-docker.yaml signup admin admin1234 admin

go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http -cfg_path /config/balance/config-docker.yaml