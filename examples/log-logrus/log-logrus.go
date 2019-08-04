package main

import (
	"errors"

	adapter "github.com/jkaveri/goabs-adapter-logrus"
	"github.com/sirupsen/logrus"

	"github.com/jkaveri/goabs-log"
)

func main() {
	// create new logrus instance
	logger := logrus.New()
	// set logrus level, this is example of you can
	// do what every you want with logrus instance.
	logger.SetLevel(logrus.InfoLevel)
	// create new adapter to connect logrus with goabs-log
	adapterLogrus := adapter.NewAdapterLogrus(logger)
	// set adapter.
	log.Configure(adapterLogrus)
	// existing code doesn't impact
	// if you switch to another log engine
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
