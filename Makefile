
# Makefile for database migrations using golang-migrate
.PHONY: create-migration-file migrate-up migrate-down install-golang-migrate
create-migration-file:
	@command -v migrate >/dev/null 2>&1 || $(MAKE) install-golang-migrate
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migration $$name
	@echo "Migration file created."

# Run all up migrations
migrate-up:
	@command -v migrate >/dev/null 2>&1 || $(MAKE) install-golang-migrate
	. .env && migrate -path migration -database $$DATABASE_URL up

# Rollback (run down migrations)
migrate-down:
	@command -v migrate >/dev/null 2>&1 || $(MAKE) install-golang-migrate
	. .env && migrate -path migration -database $$DATABASE_URL down

install-golang-migrate:
	@echo "Installing golang-migrate..."
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "golang-migrate installed."