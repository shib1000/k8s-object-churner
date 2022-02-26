package koclog

import (
	"fmt"
	"go.uber.org/zap"
)

type KocLogger struct {
	logger *zap.Logger
}

func GetKocLoggerInstance() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	zlog, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	return zlog
}
