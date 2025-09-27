
## database migrations using golang-migrate
.PHONY: install-golang-migrate create-migration-file migrate-up migrate-down

# Install golang-migrate if not already installed
install-golang-migrate:
	@echo "Installing golang-migrate..."
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "golang-migrate installed."

# Create a new migration file
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

## build: build the cmd/api application
.PHONY: build
build:
	go build -o=/tmp/bin/app ./cmd/app

## run/dev: run the application with reloading on file changes
.PHONY: run/dev
run/dev:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/app" --build.delay "100" \
		--build.exclude_dir "" \
		--misc.clean_on_exit "true"