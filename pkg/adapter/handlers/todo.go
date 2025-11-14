package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taheri24/helitask/pkg/domain"
)

// TodoHandler struct for HTTP requests
type TodoHandler struct {
	repository domain.TodoRepository
}

// CreateTodoItem handles creating a new TodoItem
func (h *TodoHandler) CreateTodoItem(c *gin.Context) {
	logger := helper.GetLogger(c)

	ctx := c.Request.Context()
	var input struct {
		Description string    `json:"description"`
		DueDate     time.Time `json:"due_date"`
	}

	logger.Verbose("Received request to create TodoItem")

	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ResponseError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	todo := domain.TodoItem{
		ID:          domain.NewUUID(),
		Description: input.Description,
		DueDate:     input.DueDate,
	}

	logger.Verbose("Creating TodoItem with ID:", todo.ID)

	if err := h.repository.Create(ctx, &todo); err != nil {
		wrappedErr := fmt.Errorf("failed to save todo item ,%w", err)
		helper.ResponseError(c, http.StatusInternalServerError, "Failed to save todo item", wrappedErr)
		return
	}
	helper.SendCreatedResponse(c, todo.ID.String())
}

// CreateTodoItem handles creating a new TodoItem
func (h *TodoHandler) GetTodoItem(c *gin.Context) {
	ctx := c.Request.Context()
	key := c.Param("id")
	uuid, err := domain.ParseUUID(key)
	if err != nil {
		helper.ResponseError(c, http.StatusBadRequest, "Invalid UUID "+key, err)
		return
	}
	dao, err := h.repository.GetByID(ctx, domain.UUID(uuid))
	if err != nil {
		if errors.Is(err, domain.ErrRecordNotFound) {
			helper.ResponseError(c, http.StatusNotFound, "record not found", err)
			return
		}
		helper.ResponseError(c, http.StatusInternalServerError, "Failed to fetch todo item", err)
		return
	}
	var output = struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
	}{dao.ID.String(), dao.Description, dao.DueDate.String()}

	helper.SendSuccessResponse(c, http.StatusOK, output)
}
