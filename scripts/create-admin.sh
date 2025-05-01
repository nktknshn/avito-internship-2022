#!/bin/bash

docker exec -it microservice-balance-balance-service-http-1 go run github.com/nktknshn/avito-internship-2022/internal/balance/cmd/cli --cfg-path config/balance/config-docker.yaml signup admin admin1234 admin