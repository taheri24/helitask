package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/taheri24/helitask/pkg/ports/storage/postgres"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

// TestCreateTodoItem tests the CreateTodoItem handler
func TestCreateTodoItem(t *testing.T) {
	db, sqlMock := postgres.NewMockDB()
	app := gin.Default()
	fxApp := fxtest.New(t, fx.Supply(app), fx.Supply(db, sqlMock), postgres.Module, Module)
	defer fxApp.RequireStart().RequireStop()

	input := `{"description": "Test Todo", "due_date": "2025-12-31T23:59:59Z"}`

	req := httptest.NewRequest("POST", "/app/v0/todo", strings.NewReader(input))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)

}
