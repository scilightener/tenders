migrate-up:
	@echo "Migrating database up..."
	@SERVER_ADDRESS="localhost:8765" POSTGRES_CONN="postgres://postgres:12345678@localhost:5432/tenders-management?sslmode=disable" go run ./cmd/migrator -migrations-path ./migrations -direction up

migrate-down:
	@echo "Migrating database down..."
	@SERVER_ADDRESS="localhost:8765" POSTGRES_CONN="postgres://postgres:12345678@localhost:5432/tenders-management?sslmode=disable" go run ./cmd/migrator -migrations-path ./migrations -direction down

run:
	@echo "Running server..."
	@SERVER_ADDRESS="localhost:8765" POSTGRES_CONN="postgres://postgres:12345678@localhost:5432/tenders-management?sslmode=disable" go run ./cmd/app

lint:
	@echo "Running linter..."
	@golangci-lint run ./... -c ./config/.golangci.yml
