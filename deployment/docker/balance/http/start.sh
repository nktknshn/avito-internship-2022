#!/bin/sh

set -e

echo "Running start.sh"

CFG_PATH="/config/balance/config-docker.yaml"

CMD_CLI="go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/cli --cfg-path $CFG_PATH"

CMD_HTTP="go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http -cfg_path $CFG_PATH"


if [ -n "$BALANCE_ADMIN_USERNAME" ]; then
    echo "Creating admin user $BALANCE_ADMIN_USERNAME"
    
    USER_EXISTS=$($CMD_CLI list-users --print-header=false $BALANCE_ADMIN_USERNAME)

    if [ -z "$USER_EXISTS" ]; then
        $CMD_CLI signup $BALANCE_ADMIN_USERNAME $BALANCE_ADMIN_PASSWORD admin
    else
        echo "Admin user $BALANCE_ADMIN_USERNAME already exists"
    fi

fi

$CMD_HTTP