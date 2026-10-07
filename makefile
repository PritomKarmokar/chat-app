DB_DSN=host=localhost port=5432 user=postgres dbname=chat_app_db password=postgres sslmode=disable
MIGRATIONS_DIR=migrations

.PHONY: migrate-up migrate-down migrate-status migrate-create

migrate-up:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" up

migrate-down:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

migrate-status:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" status

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=<migration_name>"; \
		exit 1; \
	fi
	@goose -dir $(MIGRATIONS_DIR) -s create $(name) sql