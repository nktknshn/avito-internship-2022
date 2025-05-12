#!/bin/bash

set -e

go test -race ./internal/balance/...

