MODULE = $(shell go list -m)
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
PACKAGES := $(shell go list ./... | grep -v /vendor/)

CONFIG_FILE ?= ./config/local.yml
MIGRATE := migrate -source file://database/postgres/migrations/ -database postgres://postgres:postgres@db/postgres?sslmode=disable

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up