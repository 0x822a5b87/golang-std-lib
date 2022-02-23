package main

import "go.uber.org/zap"

// {"level":"info","msg":"global Logger after"}
// {"level":"info","msg":"global SugaredLogger after"}
func example06() {
	// 全局的Logger默认并不会记录日志！它是一个无实际效果的Logger
	zap.L().Info("global Logger before")
	zap.S().Info("global SugaredLogger before")

	logger := zap.NewExample()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	zap.ReplaceGlobals(logger)
	zap.L().Info("global Logger after")
	zap.S().Info("global SugaredLogger after")
}
