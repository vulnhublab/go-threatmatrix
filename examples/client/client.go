package main

import (
	"context"

	"github.com/khulnasoft/go-intelx/gointelx"
	"github.com/sirupsen/logrus"
)

func main() {
	/*
		Making a new client through NewIntelXClient:
		This takes the following parameters:
			1. IntelXClientOptions
			2. A *http.Client (if you do not provide one. One will be made by default)
			3. LoggerParams
		These are parameters that allow you to easily configure your IntelXClient to your liking.
		For a better understanding you can read it in the documentation: https://github.com/khulnasoft/go-intelx/tree/main/examples/optionalParams
	*/

	// Configuring the IntelXClient!
	clientOptions := gointelx.IntelXClientOptions{
		Url:         "PUT-YOUR-INTELX-INSTANCE-URL-HERE",
		Token:       "PUT-YOUR-TOKEN-HERE",
		Certificate: "",
		Timeout:     0,
	}

	// Configuring the logger
	loggerParams := &gointelx.LoggerParams{
		File:      nil,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.DebugLevel,
	}

	// Making the client!
	client := gointelx.NewIntelXClient(
		&clientOptions,
		nil,
		loggerParams,
	)

	ctx := context.Background()

	tags, err := client.TagService.List(ctx)

	if err != nil {
		client.Logger.Logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("An error occurred")
	} else {
		client.Logger.Logger.WithFields(logrus.Fields{
			"tags": *tags,
		}).Info("These are your tags")
	}

}
