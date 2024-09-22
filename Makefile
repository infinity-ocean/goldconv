include infra.env
-include .env
export

# Директория, в которой хранятся исполняемые
# файлы проекта и зависимости, необходимые для сборки.
LOCAL_BIN := $(CURDIR)/bin
MIGRATIONS_DIR := ./migrations

run:
	go run .

start-infra:
	docker-compose up -d

stop-infra:
	docker-compose down
	
print-dsn:
	echo $(POSTGRES_DSN)
# Создать миграцию
migration:
	mkdir -p $(MIGRATIONS_DIR)
	$(LOCAL_BIN)/goose -dir $(MIGRATIONS_DIR) create $(shell bash -c 'read -p "Migration name: " migration_name; echo $$migration_name') sql

migration-up:
	$(LOCAL_BIN)/goose $(opts) -allow-missing -dir $(MIGRATIONS_DIR) postgres "$(POSTGRES_DSN)" up

migration-down:
	$(LOCAL_BIN)/goose $(opts) -dir $(MIGRATIONS_DIR) postgres "$(POSTGRES_DSN)" down
migration-status:
	$(LOCAL_BIN)/goose $(opts) -dir $(MIGRATIONS_DIR) postgres "$(POSTGRES_DSN)" status

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.18.0