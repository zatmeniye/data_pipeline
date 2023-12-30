package logger

import (
	"go.uber.org/zap"
)

func New() *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{"stdout", "./logs/.log"}
	logger, _ := cfg.Build()
	return logger.Sugar()
}
