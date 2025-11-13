package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/taheri24/helitask/pkg/adapter/handlers"
	"github.com/taheri24/helitask/pkg/config"
	"github.com/taheri24/helitask/pkg/logger"
	"github.com/taheri24/helitask/pkg/ports/storage/postgres"
)

func TestCreateTodoItemIntegration(t *testing.T) {
	cfg, _ := config.LoadConfig("development")
	db, _ := postgres.NewDB(cfg.DB.DSN)
	repo := postgres.NewTodoRepository(db)
	log := logger.New("")

	h := handlers.NewTodoHandler(repo, log)

	r := gin.Default()
	r.POST("/todo", h.CreateTodoItem)

	input := `{"description": "Test Todo", "due_date": "2025-12-31T23:59:59Z"}`
	req := httptest.NewRequest("POST", "/todo", strings.NewReader(input))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
