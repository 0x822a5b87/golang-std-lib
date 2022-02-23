package main

import (
	"go.uber.org/zap"
)

func example04() {
	encoderConfig := zap.NewProductionEncoderConfig()

	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	config.Encoding = "console"

	production, err := zap.NewProduction()
	if err != nil {
		return
	}
	production.Debug("this is a debug message!")
	production.Info("this is a info message!")
	production.Warn("this is a warn message!")
}
