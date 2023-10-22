package log

import "github.com/sirupsen/logrus"

var logger = logrus.New()

func SetVerbose(verbose bool) {
	if verbose {
		logger.SetLevel(logrus.TraceLevel)
	} else {
		logger.SetLevel(logrus.WarnLevel)
	}
}

func Logger() *logrus.Logger {
	return logger
}

var levels = []logrus.Level{
	logrus.PanicLevel,
	logrus.FatalLevel,
	logrus.ErrorLevel,
	logrus.WarnLevel,
	logrus.InfoLevel,
	logrus.DebugLevel,
	logrus.TraceLevel,
}
