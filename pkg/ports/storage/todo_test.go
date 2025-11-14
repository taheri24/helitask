package storage

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/taheri24/helitask/pkg/domain"
	"github.com/taheri24/helitask/pkg/logger"
	"github.com/taheri24/helitask/pkg/ports/storage/postgres"

	"github.com/taheri24/helitask/pkg/logger/testinglogger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func setupApp(t *testing.T, options ...fx.Option) (sqlmock.Sqlmock, *fxtest.App) {
	db, sqlMock := postgres.NewMockDB()
	testLogger := logger.NewSlogger(testinglogger.NewTestLogger(t))
	opts := append([]fx.Option{
		fx.NopLogger, // remove this to getting more details on DI problems
		fx.Provide(func() logger.Logger {
			return testLogger
		}), fx.Supply(db, sqlMock), Module}, options...)
	fxApp := fxtest.New(t, opts...)
	return sqlMock, fxApp
}
func TestCreateTodoItem(t *testing.T) {
	var repo domain.TodoRepository
	mockSql, fxApp := setupApp(t, fx.Populate(&repo))
	fxApp.RequireStart()
	defer fxApp.RequireStop()
	freshItem := &domain.TodoItem{
		ID:          domain.NewUUID(),
		Description: "Test Todo",
		DueDate:     time.Now(),
	}

	mockSql.ExpectExec(`^INSERT INTO.+todo_items.+`).WithArgs(freshItem.ID, freshItem.Description, freshItem.DueDate).WillReturnResult(sqlmock.NewResult(0, 1))
	if err := repo.Create(t.Context(), freshItem); err != nil {
		t.Errorf("repo.Create failed  ,%s", err)
		return
	}

	// we make sure that all expectations were met
	if err := mockSql.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
