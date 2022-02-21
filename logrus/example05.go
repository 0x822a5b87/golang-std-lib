package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type QLog struct {
	AppName string
}

func (q *QLog) Levels() []logrus.Level {
	fmt.Println("qlog hook levels")
	return logrus.AllLevels
}

func (q *QLog) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = q.AppName
	fmt.Println("qlog hook fire")
	return nil
}

func example05() {
	logrus.AddHook(&QLog{
		AppName: "xxx",
	})
	logrus.Info("hello world!!")
}
