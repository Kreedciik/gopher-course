MIGRATE_CMD=migrate
DB_URL=postgresql://postgres:root@localhost:5432/blogpost?sslmode=disable
MIGRATION_DIR=schema

.PHONY: clean

migrate-up:
	@$(MIGRATE_CMD) -database $(DB_URL) -path $(MIGRATION_DIR) up

migrate-down:
	@$(MIGRATE_CMD) -database $(DB_URL) -path $(MIGRATION_DIR) down 1

migrate-force:
	@echo "Usage: make migrate-force VERSION=<version>"
	@$(MIGRATE_CMD) -database $(DB_URL) -path $(MIGRATION_DIR) force $(VERSION)

migrate-create:
	@echo "Usage: make migrate-create NAME=<name>"
	@test -n "$(NAME)" || (echo "Error: NAME is required" && exit 1)
	@$(MIGRATE_CMD) create -ext sql -dir $(MIGRATION_DIR) -seq $(NAME)

migrate-status:
	@$(MIGRATE_CMD) -database $(DB_URL) -path $(MIGRATION_DIR) version
	
run:
	go run cmd/main.go

tidy:
	go mod tidy