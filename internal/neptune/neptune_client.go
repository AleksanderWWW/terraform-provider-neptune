package neptune

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type operationID string

type operationData struct {
	Endpoint string `mapstructure:"endpoint"`
	Method   string `mapstructure:"method"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			panic(err)
		}
	}
}

type NeptuneClient struct {
	httpClient *http.Client
	creds      credentials
}

func NewNeptuneClient(apiKey string, timeout int64) (*NeptuneClient, error) {
	timeoutDuration := time.Duration(timeout) * time.Second

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
	var opData operationData
	err := viper.UnmarshalKey(string(id), &opData)
	if err != nil {
		return nil, err
	}

	req, err := prepareRequest(c.creds.tokenOriginAddress, opData.Endpoint, opData.Method, params, headers, body)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(req)
}
