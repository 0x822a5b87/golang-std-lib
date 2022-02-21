package main

import "github.com/sirupsen/logrus"

func example01() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)

	logger := logrus.WithFields(logrus.Fields{
		"name": "xxx",
		"age":  18,
	})
	logger.Info("hello world!")
	logger.Warn("hello world again!!")

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")

	logrus.Info("this message will not show because logrus already exit")
}
