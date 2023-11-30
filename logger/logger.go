package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// TODO: options.
// TODO: test

// Logger represents a Logger object.
type Logger struct {
	logger zerolog.Logger
}

// NewLogger creates a new instance of Logger with default settings.
func NewLogger() *Logger {
	return &Logger{
		logger: loggerString(),
	}
}

func loggerString() zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	return zerolog.New(output).With().Timestamp().Logger()
}

// Debug logs a debug-level message with key-value pairs.
func (l *Logger) Debug(k, v, msg string) {
	l.logger.Debug().Str(k, v).Msg(msg)
}

// Info logs an info-level message with key-value pairs.
func (l *Logger) Info(k, v, msg string) {
	l.logger.Info().Str(k, v).Msg(msg)
}

// Warn logs a warning-level message with key-value pairs.
func (l *Logger) Warn(k, v, msg string) {
	l.logger.Warn().Str(k, v).Msg(msg)
}

// Err logs an error-level message with an associated error.
func (l *Logger) Err(e error, msg string) {
	l.logger.Err(e).Msg(msg)
}

func (l *Logger) Fatal(msg string) {
	l.logger.Fatal().Msg(msg)
}
