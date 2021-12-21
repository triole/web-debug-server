package logging

import (
	"github.com/sirupsen/logrus"
)

// LogInfo logs an info message
func (lg Logging) LogInfo(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Info(msg)
	default:
		lg.Logrus.Info(msg)
	}
}

// LogWarn logs a warning
func (lg Logging) LogWarn(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Warning(msg)
	default:
		lg.Logrus.Warning(msg)
	}
}

// LogFatal logs fatal and exits
func (lg Logging) LogFatal(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Fatal(msg)
	default:
		lg.Logrus.Fatal(msg)
	}
}

// LogError logs an error message
func (lg Logging) LogError(msg interface{}, fields interface{}) {
	var msgStr string
	switch val := msg.(type) {
	case error:
		msgStr = val.Error()
	default:
		msgStr = val.(string)
	}
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Error(msgStr)
	default:
		lg.Logrus.Error(msgStr)
	}
}
