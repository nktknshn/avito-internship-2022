test:
	go test ./...
.PHONY: test

test-race:
	./scripts/test-race.sh
.PHONY: test-race

test-verbose:
	go test -v ./...
.PHONY: test-verbose

cover-html:
	./scripts/cover-html.sh
.PHONY: cover-html

cover-html-balance-use-cases:
	./scripts/cover-html.sh github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/...
.PHONY: cover-html-balance-use-cases

cover:
	./scripts/cover.sh
.PHONY: cover

cover-balance-use-cases:
	./scripts/cover.sh github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/...
.PHONY: cover-balance-use-cases

generate-protobuf:
	./scripts/generate-protobuf.sh balance
.PHONY: generate-protobuf

generate-swagger:
	./scripts/swagger.sh && ./scripts/openapi-generate-client.sh
.PHONY: generate-swagger

lint:
	./scripts/lint.sh
.PHONY: lint

lint-nilaway:
	./scripts/lint-nilaway.sh
.PHONY: lint-nilaway

lint-all:
	./scripts/lint.sh
	./scripts/lint-nilaway.sh
.PHONY: lint-all
