
.PHONY: test
.DEFAULT: help

all: ## Run test, fmt, vet and build
	test fmt vet build

help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run go test against the code
	go test ./...

fmt:  ## Run go fmt against the code
	go fmt ./...

vet:  ## Run go vet against the code	
	go vet ./...

build:
	go build ./...

import: ## Import the MongoDB dump files located at ./resources/dump into the DB
	/bin/bash ./scripts/importer.sh

run: ##  Run docker-compose services
	docker-compose up -d

down: ## Remove all the running services
	docker-compose down



