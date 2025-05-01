#!/bin/sh

set -e

echo "Running start.sh"


go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http -cfg_path /config/balance/config-docker.yaml