package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/jkaveri/goabs/log"
)

type AdapterLogrus struct {
	logger *logrus.Logger
}

// NewAdapter
func NewAdapterLogrus(logger *logrus.Logger) *AdapterLogrus {
	return &AdapterLogrus{
		logger: logger,
	}
}

func (a *AdapterLogrus) Log(fields log.Fields) {
	level := convertToLogrusLevel(fields.Level())
	msg := fields.Message()
	err := fields.Error()
	// delete all know fields
	rest := fields.Rest()
	entry := a.logger.
		WithFields(convertToLogrusFields(rest))

	if err != nil {
		entry = entry.WithError(err)
	}
	entry.Log(level, msg)
}

var _ log.ILogAdapter = (*AdapterLogrus)(nil)

func convertToLogrusFields(fields log.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for k, v := range fields {
		logrusFields[k] = v
	}
	return logrusFields
}

func convertToLogrusLevel(level log.Level) logrus.Level {
	switch level {
	case log.LevelTrace:
		return logrus.TraceLevel
	case log.LevelDebug:
		return logrus.DebugLevel
	case log.LevelInfo:
		return logrus.InfoLevel
	case log.LevelWarn:
		return logrus.WarnLevel
	case log.LevelError:
		return logrus.ErrorLevel
	case log.LevelPanic:
		return logrus.PanicLevel
	case log.LevelFatal:
		return logrus.FatalLevel
	}
	return logrus.TraceLevel
}
