.DEFAULT_GOAL := help

migrate: ## Run database migrations
	./bin/migrate

db: ## Start docker redis server
	docker run -it --name deploygif-db -p 6379:6379 redis:alpine

db-local: ## Start local redis-server
	redis-server /usr/local/etc/redis.conf

build: ## Build the project
	go build -o deploygif-be

build-image: ## Build the container image
	docker build -t z0mbix/deploygif-be .

run: ## Run the server
	go run main.go

tidy: ## Tidy stuff
	go mod tidy
	rm deploygif-be

help: ## See all the Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help migrate db db-local build build-image run tidy
