package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/taheri24/helitask/pkg/logger"
)

// Helper provides common functionality for all handlers
type Helper struct {
	defaultLogger logger.Logger
}

var helper *Helper

// NewBaseHandler creates a new instance of BaseHandler
func NewBaseHandler(defaultLogger logger.Logger) *Helper {

	return &Helper{defaultLogger: defaultLogger}
}

// ResponseError sends a standardized error response and logs the error
func (h *Helper) ResponseError(c *gin.Context, status int, message string, err error) {
	// Store logger in a variable for reuse
	logger := h.GetLogger(c)

	if err != nil {
		message = message + ": " + err.Error()
	} else {
		logger.Error(message, err)
	}
	c.JSON(status, gin.H{"error": message})
}

// SendSuccessResponse sends a standardized success response
func (h *Helper) SendSuccessResponse(c *gin.Context, status int, data any) {
	// Store logger in a variable for reuse
	logger := h.GetLogger(c)

	logger.Verbose("Sending success response")
	c.JSON(status, gin.H{"data": data})
}

// GetLogger returns a logger for the incoming request
// The logSource is either passed explicitly or extracted from the X-LOG-SOURCE header
func (h *Helper) GetLogger(c *gin.Context) logger.Logger {
	// Read the X-LOG-SOURCE header to determine the logging source
	logSource := c.GetHeader("X-LOG-SOURCE")
	if logSource != "" {
		return logger.New(logSource)
	}
	return h.defaultLogger
}
