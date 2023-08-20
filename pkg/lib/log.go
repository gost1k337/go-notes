package lib

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"tg_bot/configs"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(cfg *configs.Config) Logger {
	zapConfig := zap.NewDevelopmentConfig()

	if cfg.Environment == "development" {
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if cfg.Environment == "production" && cfg.LogOutput != "" {
		zapConfig.OutputPaths = []string{cfg.LogOutput}
	}

	logger, _ := zapConfig.Build()
	sugar := logger.Sugar()
	return Logger{sugar}
}
