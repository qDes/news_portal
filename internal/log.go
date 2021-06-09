package internal

import "go.uber.org/zap"

func InitLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.OutputPaths = []string{"stdout"}
	//cfg.Development = true
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

