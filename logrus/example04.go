package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"time"
)

func example04() {
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		ShowFullLevel:   false,
		TimestampFormat: time.RFC3339,
	})

	logrus.Info("info msg")
}
