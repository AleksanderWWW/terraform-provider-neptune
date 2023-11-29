package neptune

import (
	"fmt"
	"net/http"
	"time"
)

type operationID string

type operationData struct {
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}

var config map[operationID]operationData

func init() {
	loadConfig(&config)
}

const configPath = "config.json"

type NeptuneClient struct {
	httpClient *http.Client
	creds      credentials
}

func NewNeptuneClient(apiKey string, timeout int64) (*NeptuneClient, error) {
	timeoutD := time.Duration(timeout)
	timeoutDuration := timeoutD * time.Second

	apiKey, err := getApiToken(apiKey)
	if err != nil {
		return nil, err
	}

	credentials, err := NewCredentials(apiKey)
	if err != nil {
		return nil, err
	}

	return &NeptuneClient{
		httpClient: &http.Client{
			Timeout: timeoutDuration,
		},
		creds: *credentials,
	}, nil
}

func (c *NeptuneClient) do(id operationID, params map[string]string, headers map[string]string, body []byte) (*http.Response, error) {
	opData, ok := config[id]
	if !ok {
		return nil, fmt.Errorf("Invalid operation '%s'", id)
	}

	req, err := prepareRequest(c.creds.tokenOriginAddress, opData.Endpoint, opData.Method, params, headers, body)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(req)
}