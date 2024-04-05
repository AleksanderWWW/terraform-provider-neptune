package neptune

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	NeptuneApiToken                   string = "NEPTUNE_API_TOKEN"
	NeptuneAllowSelfSignedCertificate string = "NEPTUNE_ALLOW_SELF_SIGNED_CERTIFICATE"
)

var stringToVisibility = map[string]string{
	"private": "priv",
	"public":  "pub",
}

var roles = map[string]bool{
	"member":  true,
	"viewer":  true,
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
	_apiToken, ok := os.LookupEnv(NeptuneApiToken)
	if !ok {
		return "", fmt.Errorf("Neptune api token not found")
	}
	return _apiToken, nil
}

func getResponseMessage(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responseString := string(responseData)

	return fmt.Sprintf("%d: %s", resp.StatusCode, responseString), nil
}

func shouldAllowSelfSignedCertificates() bool {
	envVal, ok := os.LookupEnv(NeptuneAllowSelfSignedCertificate)
	if !ok || len(envVal) == 0 {
		return false
	}

	envVal = string(strings.ToLower(envVal)[0])

	return (envVal == "t" || envVal == "1")
}
