package infrastructure

import (
	"go.uber.org/zap"
	"strings"
)

type LoggerManager struct {
	Logger *zap.Logger
}

func NewLoggerManager(Environment string) *LoggerManager {
	if strings.ToLower(Environment) == "development" {
		devLogger, err := zap.NewDevelopment()
		defer devLogger.Sync()
		if err != nil {
			return nil
		}
		return &LoggerManager{Logger: devLogger}
	}

	prodLogger, err := zap.NewProduction()
	defer prodLogger.Sync()
	if err != nil {
		return nil
	}
	return &LoggerManager{Logger: prodLogger}
}
