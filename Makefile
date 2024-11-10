include .env
#docker
#MIGRATE=docker-compose exec web migrate -path=migration -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose
DB_DSN="${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"
MIGRATE=migrate -path=database/migration -database ${DB_TYPE}"://"${DB_DSN} -verbose

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

dao:
		@command -v gentool >/dev/null 2>&1 || (echo "Installing gentool..." && go install gorm.io/gen/tools/gentool@latest)
		gentool -dsn ${DB_DSN} -fieldNullable -fieldWithIndexTag -fieldWithTypeTag -fieldSignable -onlyModel -outPath "./database/dao" -modelPkgName "dao"

swag-generate:
		swag fmt
		swag init --parseDependency --parseInternal

.PHONY: dev migrate-up migrate-down force goto drop create swag-generate dao