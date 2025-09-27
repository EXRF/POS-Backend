# POS Backend

A Point of Sale (POS) backend system built with Go, Gin, and PostgreSQL.

## Project Structure

```
POS-Backend/
├── api/                    # API documentation and OpenAPI specs
├── cmd/
│   └── app/
│       └── main.go         # Application entry point
├── internal/
│   ├── delivery/
│   │   └── http/           # HTTP handlers, routing
│   ├── entities/           # Data structures and models
│   ├── repository/         # Database interfaces and implementations
│   └── usecases/           # Business logic
├── migration/              # Database migration files
├── pkg/                    # Shared packages
│   ├── middlewares/        # Global middleware
│   ├── postgres/           # PostgreSQL connection utilities
│   └── utils/              # Utility functions
├── go.mod, go.sum          # Go module files
└── .env                    # Environment variables
```

## Architecture

The project follows a clean architecture pattern:

- **Entities**: Define core business objects
- **Repository**: Handle data persistence operations
- **Use Cases**: Implement business logic
- **Delivery**: Handle HTTP requests and responses

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd POS-Backend
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up PostgreSQL database:
```bash
# Create the database
createdb pos_backend

# Run migrations
migrate -path migration -database "postgres://postgres:postgres@localhost:5432/pos_backend?sslmode=disable" up
```

4. Set environment variables:
```bash
cp .env.example .env
# Edit .env with your database configuration
```

5. Run the application:
```bash
go run cmd/app/main.go
```

## API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### User Management
- `POST /api/v1/users/register` - Register a new user
- `POST /api/v1/users/login` - Login user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

## Environment Variables

- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: pos_backend)
- `PORT` - Server port (default: 8080)
- `GIN_MODE` - Gin mode (debug/release)

## Database Migrations

To create a new migration:
```bash
make create-migration-file
```

To run migrations:
```bash
# Run migrations up
migrate -path migration -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

# Run migrations down
migrate -path migration -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down
```