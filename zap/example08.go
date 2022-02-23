package main

import "go.uber.org/zap"

// 与标准日志库搭配使用
// 如果项目一开始使用的是标准日志库log，后面想转为zap。这时不必修改每一个文件。
// 我们可以调用zap.NewStdLog(l *Logger) *log.Logger返回一个标准的log.Logger
func example08() {
	logger := zap.NewExample()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	std := zap.NewStdLog(logger)
	std.Print("standard logger wrapper")
}
