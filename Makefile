## DEVELOPMENT COMMANDS

.PHONY: help clean full_clean check test build updatedeps cover

default: test

help: ## Shows help screen.
	@echo "\n[$(APP_NAME)]\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
	@echo ""

clean: ## Cleans the project.
	@go clean

full_clean: ## Cleans the project and all it's artifacts.
	@rm -f bin/*
	@go clean

check: ## Checks codestyle and correctness.
	@which gometalinter >/dev/null; if [ $$? -eq 1 ]; then \
		go get -u github.com/alecthomas/gometalinter; \
		gometalinter --install --update; \
	fi
	gometalinter --vendor --vendored-linters --disable-all --enable=vet --enable=golint ./...

test: ## Runs all unit tests of the project.
	@go test -cover ./...

build:
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -o bin/terraform-provider-ciscoasa

updatedeps: ## Updates the vendored Go dependencies
	@dep ensure -update

cover: ## Shows test coverage.
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out
