package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

// GetLogger returns the logger for the bot
func GetLogger(configLevel string) *logrus.Logger {

	var logger = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     getLogLevel(configLevel),
	}

	return logger
}

func getLogLevel(configLevel string) logrus.Level {
	switch configLevel {
	case "DEBUG":
		return logrus.DebugLevel
	case "FATAL":
		return logrus.FatalLevel
	case "INFO":
		return logrus.InfoLevel
	case "PANIC":
		return logrus.PanicLevel
	case "TRACE":
		return logrus.TraceLevel
	case "WARN":
		return logrus.WarnLevel
	default:
		return logrus.ErrorLevel
	}
}
