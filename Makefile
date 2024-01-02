SHELL := /usr/bin/env bash -eEuo pipefail -c
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
.DEFAULT_GOAL := build

.PHONY: generate
generate: ## Generate the project
	@echo "Generating the project..."
	@go generate ./...

.PHONY: build
build: generate ## Build the project
	@echo "Building the project..."
	@go build -o bin/ ./...

.PHONY: test
test: ## Test the project
	@echo "Testing the project..."
	@go test -v ./...

.PHONY: install-tools
install-tools: ## Install tools
	@echo "Installing tools..."
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: update-tools
update-tools: ## Update tools
	@echo "Updating tools..."
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go get -u %

.PHONY: tidy
tidy: ## Tidy the project
	@echo "Tidying the project..."
	@go mod tidy