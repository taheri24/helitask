# Helitask

Build a Go (1.24+) Todo service using hexagonal architecture. Implement a PostgreSQL-backed domain for TodoItems (id, description, dueDate) and expose RESTful APIs for handling todo items

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

## API integration tests (Hurl)

The `api_tests/` directory contains end-to-end suites expressed in the [Hurl](https://hurl.dev) DSL. These scripts exercise the
Helitask REST API using the same request / response layout described in the inline examples inside `api_tests/README.md`.

1. Install the Hurl CLI locally. Refer to the [installation guide](https://hurl.dev/docs/installation.html) for platform-specific
   instructions.
2. Start a Helitask instance (defaults to `http://localhost:8080`):
   ```bash
   APP_ENV=development go run ./main.go
   ```
3. Execute the suites with Hurl:
   ```bash
   hurl --test api_tests/todo_success.hurl
   hurl --test api_tests/todo_error_cases.hurl
   ```

The suites default to the local development base URL defined via their `[Options]` block. If your server runs elsewhere,
override it when invoking Hurl:

```bash
hurl --variable base_url=https://example.org/api/v0 --test api_tests/todo_success.hurl
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
