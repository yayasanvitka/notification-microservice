package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var (
	Info     = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	Error    = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    config.EncoderConfig.TimeKey = "timestamp"
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    config.OutputPaths = []string{
    	"log/app.log",
    	"stdout",
	}
	config.ErrorOutputPaths = []string{
    	"log/app.log",
    	"stderr",
	}
    zapLogger, _ := config.Build()
    return zapLogger
}

func Init()  {
	logManager := initZapLog()
	zap.ReplaceGlobals(logManager)
	defer logManager.Sync() // flushes buffer, if any
}