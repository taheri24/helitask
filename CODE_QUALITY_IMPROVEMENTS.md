# Code Quality Improvement Suggestions

This document outlines specific improvements to enhance code quality, maintainability, and adherence to functional programming and TDD principles.

## Table of Contents
1. [Architecture & Design](#architecture--design)
2. [Error Handling](#error-handling)
3. [Code Organization](#code-organization)
4. [Testing](#testing)
5. [Documentation](#documentation)
6. [Code Cleanliness](#code-cleanliness)
7. [Security & Validation](#security--validation)
8. [Performance](#performance)

---

## Architecture & Design

### 1. Remove Domain Layer Dependency on GORM
**Issue**: The domain layer imports `gorm.io/gorm` for `ErrRecordNotFound`, violating hexagonal architecture principles.

**Location**: `pkg/domain/legacy.go:16`

**Current Code**:
```go
var ErrRecordNotFound = gorm.ErrRecordNotFound
```

**Recommendation**: Define domain-specific errors in the domain package:
```go
package domain

import "errors"

var ErrRecordNotFound = errors.New("record not found")
```

**Impact**: High - Maintains proper separation of concerns and makes the domain layer framework-agnostic.

---

### 2. Eliminate Global Mutable State
**Issue**: Global `helper` variable in `pkg/adapter/handlers/helper.go:15` violates functional programming principles.

**Location**: `pkg/adapter/handlers/helper.go:15`

**Current Code**:
```go
var helper *Helper = NewBaseHandler(nil)
```

**Recommendation**: 
- Make `Helper` a dependency injected into handlers
- Remove global variable and pass `Helper` instance through dependency injection
- Update `TodoHandler` to accept `Helper` as a dependency

**Impact**: High - Aligns with functional programming principles and improves testability.

---

### 3. Remove Unused Error Variable
**Issue**: `ErrTodoNotFound` is defined but never used.

**Location**: `pkg/ports/storage/todo.go:15`

**Recommendation**: Remove the unused variable or use it consistently if it's intended for future use.

**Impact**: Low - Code cleanliness.

---

## Error Handling

### 4. Improve Configuration Error Handling
**Issue**: In `main.go`, configuration loading errors are only logged as warnings, but the application continues with potentially nil config.

**Location**: `main.go:28-31`

**Current Code**:
```go
cfg, err := config.LoadConfig(env)
if err != nil {
    slog.Warn(fmt.Sprintf("Error loading configuration files(.env.%s .env)", env), slog.Any("err", err))
}
```

**Recommendation**: Fail fast if configuration cannot be loaded:
```go
cfg, err := config.LoadConfig(env)
if err != nil {
    slog.Error("Failed to load configuration", slog.Any("err", err))
    os.Exit(1)
}
```

**Impact**: High - Prevents runtime errors from missing configuration.

---

### 5. Replace Panic with Proper Error Handling
**Issue**: Using `panic` in `di.go` for nil config check is not idiomatic Go.

**Location**: `pkg/di/di.go:27`

**Current Code**:
```go
if cfg == nil {
    panic("cfg==nil")
}
```

**Recommendation**: Return an error instead:
```go
if cfg == nil {
    return nil, fmt.Errorf("configuration is required but was nil")
}
```

**Impact**: Medium - Better error handling and prevents unexpected crashes.

---

### 6. Improve Error Wrapping Consistency
**Issue**: Inconsistent error message formatting (extra spaces, inconsistent comma placement).

**Locations**: 
- `pkg/adapter/handlers/todo.go:44` - `"failed to save todo item ,%w"` (space before comma)
- `pkg/ports/storage/todo.go:33` - `"failed to save todo item, %w"` (consistent)

**Recommendation**: Standardize error message format:
```go
fmt.Errorf("failed to save todo item: %w", err)
```

**Impact**: Low - Code consistency and readability.

---

### 7. Handle Context Propagation
**Issue**: In `storage/todo.go:39`, the parameter is named `context` instead of `ctx`, and context is not used for database operations.

**Location**: `pkg/ports/storage/todo.go:39`

**Current Code**:
```go
func (r *PostgresTodoRepository) GetByID(context context.Context, id domain.UUID) (*domain.TodoItem, error) {
    // context is not used
}
```

**Recommendation**: 
- Rename parameter to `ctx` for consistency
- Use context for database operations: `r.DB.WithContext(ctx).First(...)`

**Impact**: Medium - Proper context propagation for cancellation and timeouts.

---

## Code Organization

### 8. Remove Commented Code
**Issue**: Commented-out logger line in production code.

**Location**: `pkg/ports/storage/todo.go:32`

**Current Code**:
```go
///r.logger.Error("Failed to save todo item", err)
```

**Recommendation**: Remove commented code or uncomment if needed. If logging is desired, use it properly:
```go
r.logger.Error("Failed to save todo item", err)
```

**Impact**: Low - Code cleanliness.

---

### 9. Fix Incorrect Function Documentation
**Issue**: `GetTodoItem` function has incorrect GoDoc comment (copied from `CreateTodoItem`).

**Location**: `pkg/adapter/handlers/todo.go:51`

**Current Code**:
```go
// CreateTodoItem handles creating a new TodoItem
func (h *TodoHandler) GetTodoItem(c *gin.Context) {
```

**Recommendation**: Update the comment:
```go
// GetTodoItem retrieves a TodoItem by ID
func (h *TodoHandler) GetTodoItem(c *gin.Context) {
```

**Impact**: Low - Documentation accuracy.

---

### 10. Fix Test Function Documentation
**Issue**: Test function comment says it tests `CreateTodoItem` but it actually tests `GetTodoItem`.

**Location**: `pkg/adapter/handlers/todo_test.go:46`

**Recommendation**: Update the comment to match the actual test.

**Impact**: Low - Documentation accuracy.

---

### 11. Standardize Parameter Naming
**Issue**: Inconsistent parameter naming (`context` vs `ctx`).

**Recommendation**: Use `ctx` consistently across all functions that accept `context.Context`.

**Impact**: Low - Code consistency.

---

## Testing

### 12. Add Input Validation Tests
**Issue**: Missing tests for input validation (empty description, invalid dates, etc.).

**Recommendation**: Add test cases for:
- Empty description
- Description exceeding maximum length
- Invalid date formats
- Missing required fields

**Impact**: Medium - Better test coverage and validation.

---

### 13. Add Error Path Tests
**Issue**: Limited testing of error scenarios in repository layer.

**Recommendation**: Add tests for:
- Database connection failures
- Constraint violations
- Timeout scenarios using context cancellation

**Impact**: Medium - Improved error handling validation.

---

### 14. Improve Test Organization
**Issue**: Test helper functions could be better organized.

**Recommendation**: Consider extracting common test setup into a test package or using table-driven tests more consistently.

**Impact**: Low - Better test maintainability.

---

## Documentation

### 15. Add Missing GoDoc Comments
**Issue**: Some exported functions and types lack GoDoc comments.

**Locations**:
- `pkg/adapter/handlers/module.go:10` - `ProvideTodoHandler`
- `pkg/ports/storage/module.go` - Module variable
- `pkg/utils/utils.go:9` - `IsNumber`

**Recommendation**: Add comprehensive GoDoc comments for all exported symbols.

**Impact**: Medium - Better API documentation.

---

### 16. Document Error Conditions
**Issue**: Functions don't document when they return errors.

**Recommendation**: Add GoDoc comments explaining error conditions:
```go
// Create creates a new TodoItem in the repository.
// Returns an error if the database operation fails or if the todo item is invalid.
func (r *PostgresTodoRepository) Create(ctx context.Context, todo *domain.TodoItem) error
```

**Impact**: Medium - Better API understanding.

---

## Code Cleanliness

### 17. Remove Unnecessary Blank Lines
**Issue**: Inconsistent blank line usage.

**Location**: `pkg/ports/storage/todo.go:29-30`

**Recommendation**: Remove unnecessary blank lines for consistency.

**Impact**: Low - Code formatting consistency.

---

### 18. Fix Error Message Formatting
**Issue**: Inconsistent error message formatting with extra spaces.

**Location**: `pkg/adapter/handlers/todo.go:44`

**Recommendation**: Remove extra space before comma in error messages.

**Impact**: Low - Code consistency.

---

## Security & Validation

### 19. Add Input Validation
**Issue**: No validation for description length, potentially allowing extremely long strings.

**Location**: `pkg/adapter/handlers/todo.go:23-26`

**Recommendation**: Add validation:
```go
const (
    MaxDescriptionLength = 1000
    MinDescriptionLength = 1
)

if len(input.Description) < MinDescriptionLength {
    helper.ResponseError(c, http.StatusBadRequest, "description is required", nil)
    return
}
if len(input.Description) > MaxDescriptionLength {
    helper.ResponseError(c, http.StatusBadRequest, "description exceeds maximum length", nil)
    return
}
```

**Impact**: High - Prevents potential DoS attacks and data issues.

---

### 20. Validate Due Date
**Issue**: No validation for due date (could be in the past, too far in future, etc.).

**Recommendation**: Add business logic validation for due dates based on requirements.

**Impact**: Medium - Data integrity.

---

### 21. Sanitize Error Messages
**Issue**: Error messages may expose internal implementation details.

**Location**: `pkg/adapter/handlers/todo.go:44`

**Recommendation**: Avoid exposing internal error details to clients in production. Log detailed errors server-side, return generic messages to clients.

**Impact**: Medium - Security best practice.

---

## Performance

### 22. Use Context for Database Operations
**Issue**: Database operations don't use context for cancellation/timeouts.

**Location**: `pkg/ports/storage/todo.go:29,39`

**Recommendation**: Use context in GORM operations:
```go
func (r *PostgresTodoRepository) Create(ctx context.Context, todo *domain.TodoItem) error {
    if err := r.DB.WithContext(ctx).Create(todo).Error; err != nil {
        return fmt.Errorf("failed to save todo item: %w", err)
    }
    return nil
}
```

**Impact**: Medium - Better resource management and cancellation support.

---

### 23. Consider Connection Pooling Configuration
**Issue**: No explicit database connection pool configuration visible.

**Recommendation**: Configure connection pool settings in `di.ProvideDB`:
```go
sqlDB, err := db.DB()
if err != nil {
    return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
}

sqlDB.SetMaxOpenConns(25)
sqlDB.SetMaxIdleConns(5)
sqlDB.SetConnMaxLifetime(5 * time.Minute)
```

**Impact**: Medium - Better database connection management.

---

## Summary of Priority

### High Priority
1. Remove domain layer dependency on GORM (#1)
2. Eliminate global mutable state (#2)
3. Improve configuration error handling (#4)
4. Add input validation (#19)

### Medium Priority
5. Replace panic with proper error handling (#5)
6. Handle context propagation (#7)
7. Add input validation tests (#12)
8. Add error path tests (#13)
9. Add missing GoDoc comments (#15)
10. Validate due date (#20)
11. Sanitize error messages (#21)
12. Use context for database operations (#22)
13. Consider connection pooling (#23)

### Low Priority
14. Remove unused error variable (#3)
15. Improve error wrapping consistency (#6)
16. Remove commented code (#8)
17. Fix incorrect function documentation (#9, #10)
18. Standardize parameter naming (#11)
19. Improve test organization (#14)
20. Document error conditions (#16)
21. Remove unnecessary blank lines (#17)
22. Fix error message formatting (#18)

---

## Implementation Notes

- Follow TDD: Write tests first for each improvement
- Maintain backward compatibility where possible
- Update related tests when making changes
- Run `go fmt` and `go test ./...` after each change
- Keep functions pure and dependencies explicit
- Document breaking changes if any occur

