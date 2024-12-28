MIGRATIONS_PATH = ./cmd/migrations

DB_URL := $(shell yq eval '"postgres://" + .postgres.User + ":" + .postgres.Password + "@" + .postgres.Host + ":" + (.postgres.Port | tostring) + "/" + .postgres.Name + "?sslmode=require"' config.yml)

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_URL) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_URL) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: print-config
print-config:
	@echo "DB URL: $(DB_URL)"
