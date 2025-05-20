package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/khulnasoft/go-threatmatrix/gothreatmatrix"
	"github.com/sirupsen/logrus"
)

// Configure example.go() to run this funtion.

func FilePlaybookAnalysis() {
	// Configuring the Client!
	clientOptions := gothreatmatrix.ClientOptions{
		Url:         "http://localhost:80",
		Token:       "feaed162aefa6ac35bdbbcb2b93c4bdfb5db88c0",
		Certificate: "",
		Timeout:     0,
	}

	loggerParams := &gothreatmatrix.LoggerParams{
		File:      nil,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.DebugLevel,
	}

	// Making the client!
	client := gothreatmatrix.NewClient(
		&clientOptions,
		nil,
		loggerParams,
	)

	ctx := context.Background()

	basicAnalysisParams := gothreatmatrix.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gothreatmatrix.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}

	file, err := os.Open("exampleFiles/sample.jpeg")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	observableAnalysisParams := gothreatmatrix.FilePlaybookAnalysisParams{
		BasicAnalysisParams: basicAnalysisParams,
		PlaybookRequested:   "Sample_Static_Analysis",
		File:                file,
	}

	analyzerResponse, err := client.CreateFilePlaybookAnalysis(ctx, &observableAnalysisParams)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		analyzerResponseJSON, _ := json.Marshal(analyzerResponse)

		fmt.Println("========== ANALYZER RESPONSE ==========")
		fmt.Println(string(analyzerResponseJSON))
		fmt.Println("========== ANALYZER RESPONSE END ==========")
	}
}
