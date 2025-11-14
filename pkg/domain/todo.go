package domain

import (
	"context"
	"time"
)

type TodoItem struct {
	ID          UUID      `gorm:"id,primarykey"`
	Description string    `gorm:"description"`
	DueDate     time.Time `gorm:"due_date"`
}

type TodoRepository interface {
	Create(ctx context.Context, todo *TodoItem) error
	GetByID(ctx context.Context, id UUID) (*TodoItem, error)
}
