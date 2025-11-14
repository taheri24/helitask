# Helitask

Helitask is a Go service built with Uber's [`fx`](https://github.com/uber-go/fx) for dependency injection. It exposes a HTTP API for working with todo items and relies on PostgreSQL for persistence.

## Prerequisites

- [Go 1.25](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/) (recommended for bringing up PostgreSQL quickly)
- GNU Make (optional, but used by the provided helper targets)

## Initial setup

1. Install dependencies:
   ```bash
   go mod download
   ```
2. Create an environment file for the desired environment (the application defaults to `development`). At minimum you must provide a PostgreSQL DSN and the port the HTTP server should bind to.

   Example `.env.development`:
   ```dotenv
   DB_DSN=postgres://postgres:postgres@localhost:5432/helitask?sslmode=disable
   PORT=8080
   ```

   The loader looks for `.env.<environment>` first and falls back to `.env`. Make sure the file exists before starting the service.

3. (Optional) Start a local PostgreSQL instance using Docker Compose:
   ```bash
   docker compose up db
   ```

## Running the application

With a PostgreSQL instance available, you can run Helitask locally with Go:

```bash
APP_ENV=development go run ./main.go
```

Alternatively, you can use the helper target which wires up Docker Compose for you:

```bash
make run
```

This command brings up both the database and the application containers defined in `docker-compose.yml`.

## Applying database migrations

Helitask uses GORM for data access, and schema migrations can be applied via the provided Make target:

```bash
make migrate
```

This target invokes the CLI at `cmd/migrate`, running the application's migration routine against the database configured in your environment variables. Ensure that the database container (or another PostgreSQL instance matching your `DB_DSN`) is already running before executing the command.

You can also call the migration tool directly:

```bash
go run ./cmd/migrate -env=development
```

Omit the `-env` flag to fall back to the `APP_ENV` environment variable (defaulting to `development`).

## Running the tests

Execute the full test suite with Go directly:

```bash
go test ./...
```

Or use the Make wrapper:

```bash
make test
```

## Additional useful commands

- Build all packages:
  ```bash
  go build ./...
  ```
- Stop Docker Compose services started for development:
  ```bash
  docker compose down
  ```

## Troubleshooting

- **Configuration errors:** Ensure that the expected `.env.<environment>` file exists and contains a valid `DB_DSN`.
- **Database connectivity:** When using Docker Compose, PostgreSQL exposes port `5432` locally. Update your DSN if you map it differently.
- **Dependency downloads:** If module downloads fail, verify that you have internet connectivity or that the required modules are already in your module cache.
