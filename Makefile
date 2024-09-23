SHELL := /bin/bash

air:
	@echo "Starting air..."
	@air -c ./local/.air.toml

run:
	@echo "Running..."
	-@go run main.go

docker-up:
	@echo "Starting docker..."
	@docker compose --file local/docker-compose.yaml up -d

docker-down:
	@echo "Stopping docker..."
	@docker compose --file local/docker-compose.yaml down

migrate-latest:
	@echo "Migrating latest..."
	@set -a; source local/.env; set +a; go run database/main.go -migrate=latest

migrate-truncate:
	@echo "Truncating..."
	@set -a; source local/.env; set +a; go run database/main.go -migrate=0

migrate-rollback:
	@echo "Rolling back..."
	@set -a; source local/.env; set +a; go run database/main.go -migrate=-1

seed-all:
	@echo "Seeding all..."
	@set -a; source local/.env; set +a; go run database/main.go -seed=all

generate:
	@echo "Generating..."
	@go generate ./main.go
	@sqlc generate

build:
	@echo "Building..."
	@go build -o bin/main main.go

env:
	@echo "Loading env..."
	@set -a; source local/.env; set +a;
