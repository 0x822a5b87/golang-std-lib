package main

import (
	"fmt"
	"go.uber.org/zap"
)

func example02() {
	logger := zap.NewExample()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("logger.Sync() err : ", err)
		}
	}(logger)

	// 我们可以使用zap.Namespace(key string) Field构建一个命名空间，后续的Field都记录在此命名空间中
	logger.Info("tracked some metrics",
		zap.Namespace("metrics"),
		zap.Int("counter0", 1),
	)

	logger2 := logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter1", 1),
	)
	logger2.Info("tracked some metrics")
}
