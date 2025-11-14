package logger

import (
	"fmt"
	"os"

	"log/slog"
)

// Logger interface for structured logging
type Logger interface {
	Info(msg string, args ...any)
	Verbose(msg string, args ...any)
	Error(msg string, err error)
	With(fields ...any) Logger
}

// slogLogger implements the Logger interface using Go's slog package
type slogLogger struct {
	logger *slog.Logger
}
type noopLogger struct {
}

// New creates a new instance of slogLogger with a logSource
// logSource can be a path to a log file or other future log types
func New(logSource string) Logger {
	var logger *slog.Logger
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	// Check if logSource is a file path ending with .log
	if len(logSource) > 4 && logSource[len(logSource)-4:] == ".log" {
		// Open the log file for appending
		file, err := os.OpenFile(logSource, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("Error opening log file:", err)
			os.Exit(1)
		}
		// Use the file handler for logging
		handler := slog.NewTextHandler(file, opts)
		logger = slog.New(handler)
	} else {
		// Default to console logging
		handler := slog.NewTextHandler(os.Stdout, opts)
		logger = slog.New(handler)
	}

	return &slogLogger{logger: logger}
}
func Default() Logger {
	return NewSlogger(slog.Default())
}

func NewSlogger(slogger *slog.Logger) Logger {
	return &slogLogger{slogger}
}

func Nop() Logger {
	return &noopLogger{}
}

// Info logs an informational message
func (l *slogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Verbose logs a detailed message for debugging or tracking execution
func (l *slogLogger) Verbose(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Warn logs a detailed message for warning
func (l *slogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs an error message
func (l *slogLogger) Error(msg string, err error) {
	l.logger.Error(msg, slog.Any("err", err))
}

// With creates a new logger instance with additional fields
func (l *slogLogger) With(fields ...any) Logger {
	return &slogLogger{
		logger: l.logger.With(fields...),
	}
}

//-----------------------------

// Info logs an informational message
func (l *noopLogger) Info(msg string, args ...any) {
}

// Verbose logs a detailed message for debugging or tracking execution
func (l *noopLogger) Verbose(msg string, args ...any) {
}

// Warn logs a detailed message for warning
func (l *noopLogger) Warn(msg string, args ...any) {
}

// Error logs an error message
func (l *noopLogger) Error(msg string, err error) {
}

// With creates a new logger instance with additional fields
func (l *noopLogger) With(fields ...any) Logger {
	return l
}
