package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLog() {

	log.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})

}

func getLoggerLevel(value string) log.Level {
	switch value {
	case "DEBUG":
		return log.DebugLevel
	case "TRACE":
		return log.TraceLevel
	default:
		return log.InfoLevel
	}
}
