.DEFAULT_GOAL := help

APP_NAME := curanest-appointment-service
APP_DEBUG := $(APP_NAME)-debug
MAIN_FILE := main.go
GCFLAGS := all=-N -l
SERVICE_NAME := appointment_service
DOCKER_OWNER := pardes29
IMAGE_VER := v1

.PHONY: help build run build-debug debug up down tag push clean migrate-up migrate-reset migrate-down migrate-status migrate-create

help: ## Show all available commands
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -E 's/:.*##/:/g'

build: ## Build the application
	@if ! [ -f swag ]; then go install github.com/swaggo/swag/cmd/swag@latest; fi
	swag fmt
	swag init
	go build -o $(APP_NAME) $(MAIN_FILE)

run: build ## Run the application
	@./$(APP_NAME)

build-debug: ## Build the application in debug mode
	@go build -gcflags="$(GCFLAGS)" -o $(APP_DEBUG) $(MAIN_FILE)

debug: build-debug ## Run the application in debug mode
	@if ! [ -f dlv ]; then go install github.com/go-delve/delve/cmd/dlv@latest; fi
	@dlv exec ./$(APP_DEBUG)

up: ## Start Docker containers
	docker compose up -d

down: ## Stop Docker containers
	docker compose down

tag: ## Tag the Docker image
	docker tag $(SERVICE_NAME):$(IMAGE_VER) $(DOCKER_OWNER)/$(SERVICE_NAME):$(IMAGE_VER)

push: ## Push the Docker image to a registry
	docker push $(DOCKER_OWNER)/$(SERVICE_NAME):$(IMAGE_VER)

clean: ## Clean build files
	@rm -f $(APP_NAME) $(APP_DEBUG)

GOOSE_CMD := goose
DB_DRIVER := mysql
DB_DSN := curanestdev:curanestdev@tcp(103.162.14.143:3306)/appointment_db?multiStatements=True&parseTime=True&loc=Local
MIGRATION_DIR := db/migrations

migrate-up: ## Run all up migrations
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" up

migrate-reset: ## Rollback all migrations and re-run them
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" reset

migrate-down: ## Rollback the last migration
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" down

migrate-status: ## Show current migration status
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_DRIVER) "$(DB_DSN)" status

migrate-create: ## Create a new migration (use `make migrate-create name=add_users_table`)
	@if [ -z "$(name)" ]; then \
		echo "❌ Please provide a name, e.g., make migrate-create name=add_users_table"; \
	else \
		$(GOOSE_CMD) -dir $(MIGRATION_DIR) create $(name) sql; \
	fi
