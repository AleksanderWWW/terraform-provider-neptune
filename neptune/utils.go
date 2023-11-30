package neptune

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const neptuneApiToken string = "NEPTUNE_API_TOKEN"

var stringToVisibility = map[string]string {
	"private": "priv",
	"public": "pub",
}

var roles = map[string]bool {
	"member": true,
	"viewer": true,
	"manager": true,
}

func prepareRequest(host string, endpoint string, method string, params map[string]string, headers map[string]string, body []byte) (*http.Request, error) {
	baseURL := fmt.Sprintf("%s/%s", host, endpoint)
	req, err := http.NewRequest(method, baseURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Default request headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authority", host)
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Connection", "keep-alive")

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func loadConfig(c interface{}) error {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteValue, c)
}

func decodeAPIToken(apiToken string) (map[string]string, error) {
	// Decode base64
	decodedBytes, err := base64.StdEncoding.DecodeString(apiToken)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	// Unmarshal JSON
	var result map[string]string
	err = json.Unmarshal(decodedBytes, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return result, nil
}

func getApiToken(apiToken string) (string, error) {
	if apiToken != "" {
		return apiToken, nil
	}
	_apiToken, ok := os.LookupEnv(neptuneApiToken)
	if !ok {
		return "", fmt.Errorf("Neptune api token not found")
	}
	return _apiToken, nil
}
