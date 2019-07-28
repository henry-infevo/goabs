package main

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/jkaveri/goabs/log"
	logadapter "github.com/jkaveri/goabs/log/adapter-logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	log.Configure(
		logadapter.NewAdapterLogrus(logger),
	)

	log.Info("this is log message")

	log.Info(
		"this is formatted message: %d",
		log.WithFormatArg(10),
	)

	log.Info(
		"this is full feature of %s",
		log.WithFormatArg("logging"),
		log.WithFields(log.Fields{
			"username": "some username",
			"score":    15.2,
		}),
		log.WithField("age", 10),
		log.WithError(errors.New("test error")),
	)
}
