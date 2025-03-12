package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	var err error
	Log, err = zap.NewProduction() // or zap.NewDevelopment() for debug
	if err != nil {
		panic(err)
	}
}

func Sync() {
	Log.Sync()
}
