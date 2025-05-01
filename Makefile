test:
	go test ./...
.PHONY: test

test-verbose:
	go test -v ./...
.PHONY: test-verbose

cover-html:
	./scripts/coverage.sh
.PHONY: cover-html

cover:
	./scripts/cover.sh
.PHONY: cover

generate-protobuf:
	./scripts/generate-protobuf.sh balance
.PHONY: generate-protobuf
