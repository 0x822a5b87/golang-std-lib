package main

import "go.uber.org/zap"

func example05() {
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	logger.Info("hello world")
}
