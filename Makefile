GOBIN_PATH = $$PWD/.bin
ENV_VARS = GOBIN="$(GOBIN_PATH)" PATH="$(GOBIN_PATH):$$PATH"
.SILENT:

export SHELL:=/bin/bash
export SHELLOPTS:=$(if $(SHELLOPTS),$(SHELLOPTS):)pipefail:errexit

.ONESHELL:

.PHONY: build
build:
	$(ENV_VARS) goreleaser build --clean

.PHONY: build-snapshot
build-snapshot: tools
	$(ENV_VARS) goreleaser build --snapshot --clean

.PHONY: lint
lint:
	$(ENV_VARS) go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./...

.PHONY: release-snapshot
release-snapshot:
	$(ENV_VARS) goreleaser release --snapshot --clean

.PHONY: test
test:
	go test -v -count=1 -cover -tags unit -coverprofile coverage.out ./...

.PHONY: tools
tools:
	$(ENV_VARS) go install $$(go list -f '{{join .Imports " "}}' tools.go)