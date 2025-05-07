#!/bin/sh

rm -rf bench/balance/api

docker run \
    --rm \
    -u "$(id -u):$(id -g)" \
    -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/api/openapi/swagger.yaml \
    --git-user-id "nktknshn" \
    --git-repo-id "avito-internship-2022-bench" \
    --package-name "openapi" \
    --additional-properties=withGoMod=false \
    --global-property apiTests=false,modelTests=false \
    -g go \
    -o /local/bench/balance/api/