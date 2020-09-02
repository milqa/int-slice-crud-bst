package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(loggerLevel string) (*zap.Logger, error) {
	var level zapcore.Level

	if err := level.UnmarshalText([]byte(loggerLevel)); err != nil {
		return nil, fmt.Errorf("cannot unmarshal logger level: %w", err)
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("connot init logger: %w", err)
	}

	return logger, nil
}
