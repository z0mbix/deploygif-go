.DEFAULT_GOAL := help

migrate: ## Destroy all resources
	./bin/migrate

db: ## start local redis server
	redis-server /usr/local/etc/redis.conf

build: ## Build the project
	go build

help: ## See all the Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help migrate db build
