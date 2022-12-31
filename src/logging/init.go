package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logging holds the logging module
type Logging struct {
	Logrus    *logrus.Logger
	LogToFile bool
}

// Init method, does what it says
func Init(logFile string, JSONLog bool) (lg Logging) {
	timeStampFormat := "2006-01-02 15:04:05.000 MST"
	lg.Logrus = logrus.New()

	if JSONLog {
		lg.Logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "date",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "msg",
			},
			TimestampFormat:   timeStampFormat,
			PrettyPrint:       false,
			DisableHTMLEscape: false,
		})
	} else {
		lg.Logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeStampFormat,
			DisableQuote:    true,
			PadLevelText:    true,
		})
	}

	openLogFile, err := os.OpenFile(
		logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644,
	)
	if err != nil {
		lg.LogFatal(
			"Can not open log file",
			logrus.Fields{
				"logfile": logFile,
				"error":   err.Error(),
			},
		)
	}

	if logFile != "/dev/stdout" {
		lg.LogToFile = true
	}

	lg.Logrus.SetOutput(openLogFile)
	return lg
}
