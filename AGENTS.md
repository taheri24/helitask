# Project Overview & Purpose
This repository hosts the Helitask service, a Go application built with functional-programming principles and test-driven development (TDD) at its core. All contributors should strive to keep business logic pure, deterministic, and easy to reason about.
- Maintain project-wide consistency by evaluating every change against the established functional-programming and TDD rules.

# Core Technologies & Stack
- Go modules managed via `go.mod`/`go.sum`.
- Uber's `fx` for dependency injection and application lifecycle management.
- Standard Go testing framework with table-driven patterns.
- Keep all technology choices and updates aligned with the functional-programming and TDD expectations captured here.

# Dev Environment Tips
- Keep the workspace Go environment tidy by syncing `go.mod` and `go.sum` whenever dependencies change.
- Run `go fmt` on all modified Go files before committing.
- Avoid introducing global mutable state; prefer passing dependencies explicitly.
- Apply these tips uniformly so local development stays consistent with functional purity and test-first workflows.

# Build & Commands
- `go build ./...` to compile the full project.
- `go test ./...` to run the entire test suite.
- Use `make` targets when available for common workflows.
- Ensure command usage reinforces the functional-programming and TDD discipline documented here.

# Code Style
- Write small, composable functions that accept dependencies through parameters.
- Treat structs as immutable once constructed; return modified copies instead of mutating in place.
- Wrap returned errors with `fmt.Errorf` and the `%w` verb to preserve context.
- Check new and existing code for consistent adherence to these functional and TDD-aligned practices.

# Testing Instructions
- Follow TDD: add or update failing tests before implementing code changes.
- Keep tests deterministic and self-contained; do not call external services.
- Favor table-driven tests and descriptive case names for clarity.
- When fixing regressions, add tests that fail without the fix.
- Audit tests for consistency with the overarching functional-programming and TDD rules before merging.

# Architecture
- Structure modules around pure business logic with side effects encapsulated behind interfaces (e.g., repositories, loggers).
- Use `fx` modules to compose dependencies via `fx.Provide`, keeping configuration in the composition root (`main.go` or dedicated providers).
- Review architectural decisions to confirm they align with the shared functional-programming and TDD expectations.

# Security
- Do not introduce hidden state or side effects that could leak sensitive information.
- Handle errors with context while avoiding logging secrets.
- Keep security-sensitive code consistent with functional immutability and test coverage mandates.

# Git Workflow
- Create focused commits with clear messages describing the change.
- Ensure working tree is clean and tests pass before committing.
- Use the commit history to reinforce consistent functional-programming and TDD compliance.

# PR Instructions
- Summarize functional changes and associated tests in the PR description.
- Highlight adherence to TDD and functional design choices.
- Confirm that every PR narrative reflects ongoing consistency with functional-programming and TDD rules.

# Configuration
- Read configuration and environment variables only within composition roots so that functional units remain deterministic.
- Validate configuration changes against the repository's functional-programming and TDD consistency requirements.

# Key Files & Entrypoints
- `main.go`: primary entrypoint wiring the application with `fx`.
- `pkg/`: houses domain logic and supporting modules.
- Keep key files synchronized with functional-programming and TDD-aligned abstractions throughout the codebase.

# Coding Conventions & Style Guide
- Document exported functions and types with GoDoc comments describing behavior and side effects.
- Prefer pure functions and explicit dependency injection across the codebase.
- Standardize coding decisions around these functional-programming and TDD principles.

# Development & Testing Workflow
- Start by updating or writing tests that express the desired behavior.
- Implement production code to satisfy those tests while maintaining functional purity.
- Run `go test ./...` and any relevant `make` targets before pushing.
- Ensure each iteration through this workflow maintains consistent functional-programming and TDD rigor.

# Specific Instructions for AI Collaboration
- When generating code, keep functions pure, inject dependencies, and avoid global mutable state.
- Provide accompanying tests for all behavioral changes.
- Preserve this guideline file and update it only when process expectations change.
- Cross-check AI-generated outputs against these functional-programming and TDD consistency rules before acceptance.
