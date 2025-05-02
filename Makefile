test:
	go test ./...
.PHONY: test

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

generate-protobuf:
	./scripts/generate-protobuf.sh balance
.PHONY: generate-protobuf
