//go:build !android

package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel)
}
