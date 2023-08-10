package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func selectLoggerLevel (level string) logrus.Level {
	switch level {
	case "ErrorLevel" : return logrus.ErrorLevel
	case "DebugLevel" : return logrus.DebugLevel
	case "InfoLevel" : return logrus.InfoLevel
	case "FatalError" : return logrus.FatalLevel
	default : return logrus.DebugLevel
	}
}

var defaultLogger *logrus.Logger

func init() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{ PrettyPrint: true }

	LoggerFile, LoggerFileOpeningError := os.OpenFile("loggerFile.txt", os.O_WRONLY | os.O_CREATE, 0755)
	if LoggerFileOpeningError != nil {
		log.Fatal("The log file was not created")
		return
	}
	logger.SetOutput(LoggerFile)
	logger.SetLevel(logrus.DebugLevel)
	defaultLogger = logger
	// defer LoggerFile.Close()
}

func ThrowErrorLog(data interface{}) {
	defaultLogger.SetLevel(selectLoggerLevel("ErrorLevel"))
	defaultLogger.Error(data)
}

func ThrowCommonLog(data interface{}) {
	defaultLogger.SetLevel(selectLoggerLevel("InfoLevel"))
	defaultLogger.Info(data)
}

func ThrowDebugLog(data interface{}) {
	defaultLogger.SetLevel(selectLoggerLevel("DebugLevel"))
	defaultLogger.Debug(data)
}