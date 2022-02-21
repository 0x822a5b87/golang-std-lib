package main

import (
	"io/ioutil"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

func init() {
	hookConfig := logredis.HookConfig{
		Host:     "test.dc.data.woa.com",
		Port:     16379,
		Key:      "logrus||hook",
		Format:   "v0",
		App:      "awesome",
		Hostname: "test.dc.data.woa.com",
		TTL:      3600,
		Password: "123456",
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

func example06() {
	logrus.Info("just some info logging...")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"foo":    "bar",
		"this":   "that",
	}).Info("additional fields are being logged as well")

	logrus.SetOutput(ioutil.Discard)
	logrus.Info("This will only be sent to Redis")
}
