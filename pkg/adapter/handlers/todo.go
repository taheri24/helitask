package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/taheri24/helitask/pkg/domain"
	"github.com/taheri24/helitask/pkg/logger"
)

// TodoHandler struct for HTTP requests
type TodoHandler struct {
	*Helper
	repository domain.TodoRepository
}

// NewTodoHandler creates a new Handler instance
func NewTodoHandler(repository domain.TodoRepository, logger logger.Logger) *TodoHandler {
	return &TodoHandler{NewBaseHandler(logger), repository}
}

// CreateTodoItem handles creating a new TodoItem
func (h *TodoHandler) CreateTodoItem(c *gin.Context) {
	// Store logger in a variable for reuse
	logger := h.GetLogger(c)

	ctx := c.Request.Context()
	var input struct {
		Description string    `json:"description"`
		DueDate     time.Time `json:"due_date"`
	}

	logger.Verbose("Received request to create TodoItem")

	if err := c.ShouldBindJSON(&input); err != nil {
		h.ResponseError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	todo := domain.TodoItem{
		ID:          domain.NewUUID(),
		Description: input.Description,
		DueDate:     input.DueDate,
	}

	logger.Verbose("Creating TodoItem with ID:", todo.ID)

	if err := h.repository.Create(ctx, &todo); err != nil {
		wrappedErr := errors.Wrap(err, "failed to save todo item")
		h.ResponseError(c, http.StatusInternalServerError, "Failed to save todo item", wrappedErr)
		return
	}

	h.SendSuccessResponse(c, http.StatusCreated, todo)
}
