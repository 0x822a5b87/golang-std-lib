package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func example01() {
	logger := zap.NewExample()
	defer func(logger *zap.Logger) {
		// zap底层 API 可以设置缓存，所以一般使用defer logger.Sync()将缓存同步到文件中。
		err := logger.Sync()
		if err != nil {
			fmt.Println("logger.Sync() err : ", err)
		}
	}(logger)

	// 由于fmt.Printf之类的方法大量使用interface{}和反射，会有不少性能损失，并且增加了内存分配的频次。
	// zap为了提高性能、减少内存分配次数，没有使用反射，而且默认的Logger只支持强类型的、结构化的日志。
	url := "https://example.org/api"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// 每个字段都用方法包一层用起来比较繁琐。zap也提供了便捷的方法SugarLogger，可以使用printf格式符的方式。
	sugar := logger.Sugar()
	// SugarLogger还支持以w结尾的方法，这种方式不需要先创建字段对象，直接将字段名和值依次放在参数中即可，如例子中的 Infow
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
