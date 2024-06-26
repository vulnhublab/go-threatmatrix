package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/khulnasoft/go-intelx/gointelx"
	"github.com/sirupsen/logrus"
)

func main() {

	// Configuring the IntelXClient!
	clientOptions := gointelx.IntelXClientOptions{
		Url:         "PUT-YOUR-INTELX-INSTANCE-URL-HERE",
		Token:       "PUT-YOUR-TOKEN-HERE",
		Certificate: "",
		Timeout:     0,
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

	basicAnalysisParams := gointelx.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelx.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}

	observableAnalysisParams := gointelx.ObservableAnalysisParams{
		BasicAnalysisParams:      basicAnalysisParams,
		ObservableName:           "192.168.69.42",
		ObservableClassification: "ip",
	}

	analyzerResponse, err := client.CreateObservableAnalysis(ctx, &observableAnalysisParams)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		analyzerResponseJSON, _ := json.Marshal(analyzerResponse)
		fmt.Println("JOB ID")
		fmt.Println(analyzerResponse.JobID)
		fmt.Println("JOB ID END")
		fmt.Println("========== ANALYZER RESPONSE ==========")
		fmt.Println(string(analyzerResponseJSON))
		fmt.Println("========== ANALYZER RESPONSE END ==========")
	}
}
