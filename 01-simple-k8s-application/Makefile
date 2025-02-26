 .DEFAULT_GOAL := help

NAME = example
GOLANGCI_LINT_TIMEOUT ?= 1m

.PHONY: pre-commit-install
pre-commit-hooks-install: ## Install the pre-commit hooks
	pre-commit install

.PHONY: build
build: build-bin build-image ## Build Go binaries & Docker image

.PHONY: build-bin
build-bin: ## Build Go server for present architecture
	go build -o build/$(NAME) ./cmd

.PHONY: build-image
build-image: ## Build Docker image
	docker build -t thylong/$(NAME):latest .
	docker compose build example

.PHONY: push-docker-image
push-docker-image: ## Push Docker image
	docker trust sign thylong/$(NAME):latest

.PHONY: scan-docker-image
scan-docker-image: ## Scan latest local example image (using docker scan from Snyk)
	docker scan --dependency-tree --severity=low thylong/example

.PHONY: docker-cleanup
docker-cleanup: ## Delete all the docker-compose containers
	docker-compose down --remove-orphans

.PHONY: test
test: lint license-check scan-docker-image test-bench test-unit test-cleanup ## Launch all tests sequentially

.PHONY: lint
lint: ## Scan repository with linters
	golangci-lint run -c .golangci.yaml --timeout $(GOLANGCI_LINT_TIMEOUT)

.PHONY: test-bench
test-bench: ## Launch Go benchmark tests
	go test -bench -benchmem ./...

.PHONY: test-unit
test-unit: ## Launch Go unit tests
	go test -cover ./...

.PHONY: doc
doc: ## Update documentation
	go doc -http:=6060

.PHONY: license-check
license-check: ## Check if licence headers are missing on any files
	docker run -it --rm -v $(PWD):/github/workspace apache/skywalking-eyes header check

.PHONY: license-fix
license-fix: ## Fix missing licence headers on any files
	docker run -it --rm -v $(PWD):/github/workspace apache/skywalking-eyes header fix

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
