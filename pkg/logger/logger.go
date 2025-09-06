package logger

import "go.uber.org/zap"

type Logger struct {
	logger *zap.Logger
}

func NewZapLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &Logger{
		logger: logger,
	}, nil
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{logger: l.logger.With(fields...)}
}
