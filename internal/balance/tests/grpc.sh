#!/bin/sh

# Примеры команды для ручного тестирования

TOKEN=$(grpcurl -plaintext -d '{"username":"admin","password":"admin1234"}' 127.0.0.1:8083 balance.BalanceService.AuthSignIn | jq -r '.token')

# reflection
grpcurl -plaintext 127.0.0.1:8083 list balance.BalanceService

# auth sign in
grpcurl -plaintext -d '{"username":"admin","password":"admin1234"}' 127.0.0.1:8083 balance.BalanceService.AuthSignIn

# get balance
grpcurl -plaintext -H authorization:$TOKEN -d '{"user_id":1}' 127.0.0.1:8083 balance.BalanceService.GetBalance

# reserve (missing token)
grpcurl -plaintext -d '{"user_id":1,"product_id":1,"order_id":1,"amount":100}' 127.0.0.1:8083 balance.BalanceService.Reserve


