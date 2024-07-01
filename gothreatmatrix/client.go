// go-threatmatrix provides an SDK to easily integrate intelx with your own set of tools.

// go-threatmatrix makes it easy to automate, configure, and use intelx with your own set of tools
// with its Idiomatic approach making an analysis is easy as just writing one line of code!
package gothreatmatrix

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// IntelXError represents an error that has occurred when communicating with IntelX.
type IntelXError struct {
	StatusCode int
	Message    string
	Response   *http.Response
}

// Error lets you implement the error interface.
// This is used for making custom go errors.
func (intelXError *IntelXError) Error() string {
	errorMessage := fmt.Sprintf("Status Code: %d \n Error: %s", intelXError.StatusCode, intelXError.Message)
	return errorMessage
}

// newIntelXError lets you easily create new IntelXErrors.
func newIntelXError(statusCode int, message string, response *http.Response) *IntelXError {
	return &IntelXError{
		StatusCode: statusCode,
		Message:    message,
		Response:   response,
	}
}

type successResponse struct {
	StatusCode int
	Data       []byte
}

// IntelXClientOptions represents the fields needed to configure and use the IntelXClient
type IntelXClientOptions struct {
	Url   string `json:"url"`
	Token string `json:"token"`
	// Certificate represents your SSL cert: path to the cert file!
	Certificate string `json:"certificate"`
	// Timeout is in seconds
	Timeout uint64 `json:"timeout"`
}

// IntelXClient handles all the communication with your IntelX instance.
type IntelXClient struct {
	options          *IntelXClientOptions
	client           *http.Client
	TagService       *TagService
	JobService       *JobService
	AnalyzerService  *AnalyzerService
	ConnectorService *ConnectorService
	UserService      *UserService
	Logger           *IntelXLogger
}

// TLP represents an enum for the TLP attribute used in IntelX's REST API.
//
// IntelX docs: https://intelx.readthedocs.io/en/latest/Usage.html#tlp-support
type TLP int

// Values of the TLP enum.
const (
	WHITE TLP = iota + 1
	GREEN
	AMBER
	RED
)

// TLPVALUES represents a map to easily access the TLP values.
var TLPVALUES = map[string]int{
	"WHITE": 1,
	"GREEN": 2,
	"AMBER": 3,
	"RED":   4,
}

// Overriding the String method to get the string representation of the TLP enum
func (tlp TLP) String() string {
	switch tlp {
	case WHITE:
		return "WHITE"
	case GREEN:
		return "GREEN"
	case AMBER:
		return "AMBER"
	case RED:
		return "RED"
	}
	return "WHITE"
}

// ParseTLP is used to easily make a TLP enum
func ParseTLP(s string) TLP {
	s = strings.TrimSpace(s)
	value, ok := TLPVALUES[s]
	if !ok {
		return TLP(0)
	}
	return TLP(value)
}

// Implementing the MarshalJSON interface to make our custom Marshal for the enum
func (tlp TLP) MarshalJSON() ([]byte, error) {
	return json.Marshal(tlp.String())
}

// Implementing the UnmarshalJSON interface to make our custom Unmarshal for the enum
func (tlp *TLP) UnmarshalJSON(data []byte) (err error) {
	var tlpString string
	if err := json.Unmarshal(data, &tlpString); err != nil {
		return err
	}
	if *tlp = ParseTLP(tlpString); err != nil {
		return err
	}
	return nil
}

// NewIntelXClient lets you easily create a new IntelXClient by providing IntelXClientOptions, http.Clients, and LoggerParams.
func NewIntelXClient(options *IntelXClientOptions, httpClient *http.Client, loggerParams *LoggerParams) IntelXClient {

	var timeout time.Duration

	if options.Timeout == 0 {
		timeout = time.Duration(10) * time.Second
	} else {
		timeout = time.Duration(options.Timeout) * time.Second
	}

	// configuring the http.Client
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: timeout,
		}
	}

	// configuring the client
	client := IntelXClient{
		options: options,
		client:  httpClient,
	}

	// Adding the services
	client.TagService = &TagService{
		client: &client,
	}
	client.JobService = &JobService{
		client: &client,
	}
	client.AnalyzerService = &AnalyzerService{
		client: &client,
	}
	client.ConnectorService = &ConnectorService{
		client: &client,
	}
	client.UserService = &UserService{
		client: &client,
	}

	// configuring the logger!
	client.Logger = &IntelXLogger{}
	client.Logger.Init(loggerParams)

	return client
}

// NewIntelXClientThroughJsonFile lets you create a new IntelXClient through a JSON file that contains your IntelXClientOptions
func NewIntelXClientThroughJsonFile(filePath string, httpClient *http.Client, loggerParams *LoggerParams) (*IntelXClient, error) {
	optionsBytes, err := os.ReadFile(filePath)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not read %s", filePath)
		intelXError := newIntelXError(400, errorMessage, nil)
		return nil, intelXError
	}

	intelXClientOptions := &IntelXClientOptions{}
	if unmarshalError := json.Unmarshal(optionsBytes, &intelXClientOptions); unmarshalError != nil {
		return nil, unmarshalError
	}

	intelXClient := NewIntelXClient(intelXClientOptions, httpClient, loggerParams)

	return &intelXClient, nil
}

// buildRequest is used for building requests.
func (client *IntelXClient) buildRequest(ctx context.Context, method string, contentType string, body io.Reader, url string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)

	tokenString := fmt.Sprintf("token %s", client.options.Token)

	request.Header.Set("Authorization", tokenString)
	return request, nil
}

// newRequest is used for making requests.
func (client *IntelXClient) newRequest(ctx context.Context, request *http.Request) (*successResponse, error) {
	response, err := client.client.Do(request)

	// Checking for context errors such as reaching the deadline and/or Timeout
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer response.Body.Close()

	msgBytes, err := ioutil.ReadAll(response.Body)
	statusCode := response.StatusCode
	if err != nil {
		errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
		intelXError := newIntelXError(statusCode, errorMessage, response)
		return nil, intelXError
	}

	if statusCode < http.StatusOK || statusCode >= http.StatusBadRequest {
		errorMessage := string(msgBytes)
		intelXError := newIntelXError(statusCode, errorMessage, response)
		return nil, intelXError
	}

	sucessResp := successResponse{
		StatusCode: statusCode,
		Data:       msgBytes,
	}

	return &sucessResp, nil
}
