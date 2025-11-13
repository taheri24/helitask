package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UUID uuid.UUID
type TodoItem struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
}

func NewUUID() uuid.UUID {
	return uuid.New()
}

type TodoRepository interface {
	Create(ctx context.Context, todo *TodoItem) error
	GetByID(ctx context.Context, id UUID) (*TodoItem, error)
}
