BUILDCOMMIT := $(shell git describe --dirty --always)
BUILDDATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VER_FLAGS=-X main.commit=$(BUILDCOMMIT) -X main.date=$(BUILDDATE)

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.DEFAULT_GOAL:=help

##@ Build

.PHONY: build
build: ## Build the operator
	@go build -ldflags "$(VER_FLAGS)" ./cmd/payments-server

.PHONY: release
release: ## Build a release version of service
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-w -s $(VER_FLAGS)" -o $(GOPATH)/bin/payments-server ./cmd/payments-server

.PHONY: docker-build
docker-build: ## Build a release version of service via multi-stage docker file
	docker build  .

##@ Testing & CI

.PHONY: test
test:   ## Run unit tests
	@git diff --exit-code swagger/swagger.yaml > /dev/null || (echo "Swagger changed, update generated code"; exit 1)
	@go test -v -covermode=count -coverprofile=coverage.out ./pkg/... ./cmd/...

.PHONY: lint
lint: ## Run linting over the codebase
	golangci-lint run -v

.PHONY: ci
ci: test lint ## Target for CI system to invoke to run tests and linting

##@ Code Generation

.PHONY: codegen
codegen:
	swagger generate server -t pkg/gen -f ./swagger/swagger.yaml -A payments --flag-strategy=pflag --exclude-main

##@ Utility

.PHONY: fmt
fmt: ## Format all the source code using gofmt
	@gofmt -l -w $(SRC)

.PHONY: help
help:  ## Display this help. Thanks to https://suva.sh/posts/well-documented-makefiles/
@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
