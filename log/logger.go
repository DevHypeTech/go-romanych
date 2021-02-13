package log

import   "github.com/sirupsen/logrus"

var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}
