package testinglogger

import (
	"context"
	"log/slog"
	"testing"
)

// Custom handler for pairing slog with testing.T
type TestHandler struct {
	t *testing.T
}

func (h *TestHandler) Handle(ctx context.Context, record slog.Record) error {
	// Format the log entry to string (you can customize the format)
	msg := record.Message

	// Use the appropriate method on t based on log level
	switch record.Level {
	case slog.LevelDebug, slog.LevelInfo:
		h.t.Log(msg) // Info/Debug level: Log as normal
	case slog.LevelWarn, slog.LevelError:
		h.t.Error(msg) // Error/Warning level: Error log
	default:
		h.t.Log(msg)
	}

	return nil
}

func (h *TestHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (h *TestHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}
func (h *TestHandler) WithGroup(name string) slog.Handler {
	return h
}

// Helper function to create a new slog.Logger with the TestHandler
func NewTestLogger(t *testing.T) *slog.Logger {
	return slog.New(&TestHandler{t})
}
