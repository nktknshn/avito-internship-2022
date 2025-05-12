#!/bin/sh

set -e

echo "Running start.sh"

CFG_PATH="/config/balance/config-docker.yaml"

if [ -n "$BALANCE_ADMIN_USERNAME" ]; then
    echo "Creating admin user $BALANCE_ADMIN_USERNAME"
    
    go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/cli --cfg-path $CFG_PATH signup $BALANCE_ADMIN_USERNAME $BALANCE_ADMIN_PASSWORD admin
fi

go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http -cfg_path $CFG_PATH