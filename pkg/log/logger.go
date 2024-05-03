package log

import (
	"log"

	"go.uber.org/zap"
)

func New(debug bool) *zap.Logger {
	var logger *zap.Logger
	var err error

	if debug {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Panicln("[-] Error setting logger")
	}

	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()
	sugar.Infof("[+] Logger Module initialized.")

	return logger
}
