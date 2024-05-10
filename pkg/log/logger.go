package log

import (
	"log"
	"os"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var zapLog *zap.Logger

func New(debug bool) *zap.Logger {
	var logger *zap.Logger
	var err error

	if debug {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Panicln("[-] error setting logger")
	}

	// flushes buffer, if any
	defer (func() {
		if err := logger.Sync(); err != nil {
			logger.Sugar().Debugln("[-] error: failure syncing logger")
		}
	})()

	sugar := logger.Sugar()
	sugar.Infof("[+] Logger Module initialized.")

	return logger
}

func init() {

	var err error
	debug := os.Getenv("DEBUG") == "1"

	if debug {
		config := zap.NewProductionConfig()
		enccoderConfig := zap.NewProductionEncoderConfig()
		zapcore.TimeEncoderOfLayout("Jan _2 15:04:05.000000000")
		enccoderConfig.StacktraceKey = "" // to hide stacktrace info
		config.EncoderConfig = enccoderConfig
		zapLog, err = config.Build(zap.AddCallerSkip(1))

	} else {
		zapLog, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	zapLog.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	zapLog.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	zapLog.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	zapLog.Fatal(message, fields...)
}

func GetLogger() *zap.Logger {
	return zapLog
}
