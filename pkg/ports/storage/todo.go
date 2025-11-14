package storage

import (
	"context"
	"errors"

	"fmt"

	"github.com/taheri24/helitask/pkg/domain"
	"github.com/taheri24/helitask/pkg/logger"
	"gorm.io/gorm"
)

// Custom error for not found Todo items
var ErrTodoNotFound = errors.New("todo not found")

// PostgresTodoRepository implements the TodoRepository interface
type PostgresTodoRepository struct {
	DB     *gorm.DB
	logger logger.Logger
}

// NewTodoRepository creates a new instance of the PostgresTodoRepository
func NewTodoRepository(db *gorm.DB, logger logger.Logger) domain.TodoRepository {
	return &PostgresTodoRepository{DB: db, logger: logger}
}

// Create implements the TodoRepository interface for PostgreSQL
func (r *PostgresTodoRepository) Create(ctx context.Context, todo *domain.TodoItem) error {

	if err := r.DB.Create(todo).Error; err != nil {
		///r.logger.Error("Failed to save todo item", err)
		return fmt.Errorf("failed to save todo item, %w", err)
	}
	return nil
}

// GetByID retrieves a TodoItem by ID
func (r *PostgresTodoRepository) GetByID(context context.Context, id domain.UUID) (*domain.TodoItem, error) {
	var todo domain.TodoItem
	key := id.String()
	if err := r.DB.First(&todo, "id=?", key).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
