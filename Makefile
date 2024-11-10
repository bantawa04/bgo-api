include .env
#docker
#MIGRATE=docker-compose exec web migrate -path=migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose

MIGRATE=migrate -path=migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose

migrate-up:
		$(MIGRATE) up

migrate-down:
		$(MIGRATE) down

force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

migrate-create:
		@echo "Enter the name of the migration:"; \
		read NAME; \
		$(MIGRATE) create -ext sql -dir database/migration $$NAME

swag-generate:
		swag fmt
		swag init --parseDependency --parseInternal

.PHONY: dev migrate-up migrate-down force goto drop create swag-generate