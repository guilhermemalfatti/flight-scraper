package scraper

import (
	"os"

	"github.com/sirupsen/logrus"
)

type rootLoggerImpl struct {
	*logrus.Logger
}

func NewLogger() *logrus.Logger {

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	//logger.Formatter = &logrus.JSONFormatter{}

	return logger
}
