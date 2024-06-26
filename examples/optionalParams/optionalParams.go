package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/khulnasoft/go-intelx/gointelx"
	"github.com/sirupsen/logrus"
)

/*
For this example I'll be using the tag params!
*/
func main() {

	// Configuring the IntelXClient!
	clientOptions := gointelx.IntelXClientOptions{
		Url:         "PUT-YOUR-INTELX-INSTANCE-URL-HERE",
		Token:       "PUT-YOUR-TOKEN-HERE",
		Certificate: "",
	}

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

	// making the tag parameters!
	tagParams := gointelx.TagParams{
		Label: "your super duper cool tag label!",
		Color: "#ffb703",
	}
	createdTag, err := client.TagService.Create(ctx, &tagParams)
	if err != nil {
		fmt.Println(err)
	} else {
		tagJson, err := json.Marshal(createdTag)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(tagJson))
		}
	}

}
