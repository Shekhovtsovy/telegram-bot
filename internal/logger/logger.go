package logger

import (
	"bot/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	"os"
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

// Error send error log
func (l *log) Error(msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("facility", l.facility))
	l.logger.Error(msg, fields...)
}

// NewLogger returns new logger
func NewLogger(facility string) Log {
	// encode
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// writer
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(os.Stdout)
	// add additional writer
	cfg := config.GetConfig()
	if cfg.Log.GraylogEnable == true {
		gelfWriter, _ := gelf.NewUDPWriter(cfg.Log.GraylogUdpUri)
		writerSyncer := zapcore.AddSync(gelfWriter)
		multiWriteSyncer = zapcore.NewMultiWriteSyncer(writerSyncer, os.Stdout)
	}
	core := zapcore.NewCore(encoder, multiWriteSyncer, zapcore.InfoLevel)
	zapLog := zap.New(core)

	return &log{
		logger:   zapLog,
		facility: facility,
	}
}
