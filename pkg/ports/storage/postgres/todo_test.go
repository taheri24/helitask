package postgres

import (
	"testing"
	"time"

	"github.com/taheri24/helitask/pkg/domain"
)

func TestCreateTodoItem(t *testing.T) {
	mockDB, mockSql := NewMockDB()
	repo := NewTodoRepository(mockDB)

	repo.Create(t.Context(), &domain.TodoItem{
		ID:          domain.NewUUID(),
		Description: "Test Todo",
		DueDate:     time.Now(),
	})

	// we make sure that all expectations were met
	if err := mockSql.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
