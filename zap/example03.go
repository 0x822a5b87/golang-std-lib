package main

import (
	"encoding/json"
	"go.uber.org/zap"
)

func example03() {
	rawJSON := []byte(`{
    "level":"debug",
    "encoding":"json",
    "outputPaths": ["stdout", "server.log"],
    "errorOutputPaths": ["stderr"],
    "initialFields":{"name":"dj"},
    "encoderConfig": {
      "messageKey": "encode_message",
      "levelKey": "level",
      "levelEncoder": "lowercase"
    }
  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	// {"level":"info","encode_message":"server start work successfully!","name":"dj"}
	logger.Info("server start work successfully!")
}
