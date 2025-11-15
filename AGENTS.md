## Project Overview & Purpose
This repository hosts the Helitask service, a Go application built with functional-programming principles and test-driven development (TDD) at its core. All contributors should strive to keep business logic pure, deterministic, and easy to reason about.

## Core Technologies & Stack
- Uber's `fx` for dependency injection and application lifecycle management.
- Standard Go testing framework with table-driven(TestCase) patterns.
- Keep all technology choices and updates aligned with the functional-programming and TDD expectations captured here.

## Dev Environment Tips
- Keep the workspace Go environment tidy by syncing `go.mod` and `go.sum` whenever dependencies change.
- Run `go fmt` on all modified Go files before committing.
- Avoid introducing global mutable state; prefer passing dependencies explicitly.
- Apply these tips uniformly so local development stays consistent with functional purity and test-first workflows.

## Build & Commands
- `go build ./...` to compile the full project.
- `go test ./...` to run the entire test suite.
- Use `make` targets when available for common workflows.

## Code Style
- Write small, composable functions that accept dependencies through parameters.
- Treat structs as immutable once constructed; return modified copies instead of mutating in place.
- Wrap returned errors with `fmt.Errorf` and the `%w` verb to preserve context.
- Check new and existing code for consistent adherence to these functional and TDD-aligned practices.

## Testing Instructions
- Follow TDD: add or update failing tests before implementing code changes.
- Keep tests deterministic and self-contained; do not call external services.
- Favor table-driven tests and descriptive case names for clarity.
- When fixing regressions, add tests that fail without the fix.

## Architecture
- Structure modules around pure business logic with side effects encapsulated behind interfaces (e.g., repositories, loggers).
- Use `fx` modules to compose dependencies via `fx.Provide`, keeping configuration in the composition root (`main.go` or dedicated providers).

### Hexagonal Architecture (Ports and Adapters) Principles
- Design modules with clear separation between core business logic (domain) and external dependencies (adapters).
- Define interfaces (ports) for external interactions, such as repositories, services, or APIs, within the domain layer.
- Implement adapters to fulfill these interfaces, keeping them isolated from the core logic.
- Ensure the domain layer is unaware of the implementation details of adapters.
- Use dependency injection to bind ports to their respective adapters at the composition root.
- Write tests for the domain layer using mock implementations of ports to maintain isolation and determinism.
- Keep adapters thin and focused on translating between external systems and the domain layer.
- Avoid coupling the domain layer to specific frameworks or libraries to maintain portability and testability.
- Document the purpose and usage of each port and adapter to ensure clarity and maintainability.
- Regularly review the architecture to ensure adherence to these principles as the codebase evolves.
## Security
- Do not introduce hidden state or side effects that could leak sensitive information.
- Handle errors with context while avoiding logging secrets.
- Keep security-sensitive code consistent with functional immutability and test coverage mandates.

## Git Workflow
- Create focused commits with clear messages describing the change.
- Ensure working tree is clean and tests pass before committing.
- Use descriptive branch names that reflect the purpose of the changes (e.g., `feat/add-auth`, `bugfix/fix-crash`, `chore/update-deps`).
- Follow the convention: `<type>/<description>` where `<type>` is one of `feat`, `bugfix`, `chore`, or `hotfix`.
- Keep branch names lowercase and use hyphens to separate words.

## PR Instructions
- Summarize functional changes and associated tests in the PR description.

## Configuration
- Read configuration and environment variables only within composition roots so that functional units remain deterministic.

## Key Files & Entrypoints
- `main.go`: primary entrypoint wiring the application with `fx`.
- `pkg/`: houses domain logic and supporting modules.


## Coding Conventions & Style Guide
- Document exported functions and types with GoDoc comments describing behavior and side effects.
- Prefer pure functions and explicit dependency injection across the codebase.
- Standardize coding decisions around these functional-programming and TDD principles.

## Development & Testing Workflow
- Start by updating or writing tests that express the desired behavior.
- Implement production code to satisfy those tests while maintaining functional purity.
- Run `go test ./...` and any relevant `make` targets before pushing.

## Specific Instructions for AI Collaboration
- When generating code, keep functions pure, inject dependencies, and avoid global mutable state.
- Provide accompanying tests for all behavioral changes.
- Preserve this guideline file and update it only when process expectations change.
