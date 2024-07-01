package gothreatmatrix

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// LoggerParams represents the fields to configure your logger.
type LoggerParams struct {
	File      io.Writer
	Formatter logrus.Formatter
	Level     logrus.Level
}

// IntelXLogger represents a logger to be used by the developer.
// IntelXLogger implements the Logrus logger.
//
// Logrus docs: https://github.com/sirupsen/logrus
type IntelXLogger struct {
	Logger *logrus.Logger
}

// Init initializes the IntelXLogger via LoggerParams
func (intelXLogger *IntelXLogger) Init(loggerParams *LoggerParams) {
	logger := logrus.New()

	// Where to log the data!
	if loggerParams.File == nil {
		logger.SetOutput(os.Stdout)
	} else {
		logger.Out = loggerParams.File
	}

	if loggerParams.Formatter != nil {
		logger.SetFormatter(loggerParams.Formatter)
	}

	logger.SetLevel(loggerParams.Level)
	intelXLogger.Logger = logger
}
