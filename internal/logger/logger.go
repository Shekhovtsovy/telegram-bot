package logger

import (
	"go.uber.org/zap"
)

// Log is an interface which provides methods for logger
type Log interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type log struct {
	logger   *zap.Logger
	facility string
}

// Info send info log
func (l *log) Info(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("facility", l.facility))
	l.logger.Info(msg, fields...)
}

// Error send info log
func (l *log) Error(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("facility", l.facility))
	l.logger.Error(msg, fields...)
}

// NewLogger return a new logger
func NewLogger(facility string) Log {
	zapLog, _ := zap.NewProduction()
	defer zapLog.Sync()
	return &log{
		logger:   zapLog,
		facility: facility,
	}
}
