package neptune

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

const NeptuneApiToken string = "NEPTUNE_API_TOKEN"

type credentials struct {
	apiTokenStr        string
	tokenOriginAddress string
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

func newCredentials(apiToken *string) (*credentials, error) {
	var apiTokenStr string

	if apiToken == nil {
		apiTokenStr = os.Getenv(NeptuneApiToken)
	} else {
		apiTokenStr = *apiToken
	}
	tokenDict, err := decodeAPIToken(apiTokenStr)
	if err != nil {
		return nil, err
	}

	tokenOriginAddress, ok := tokenDict["api_address"]
	if !ok {
		return nil, fmt.Errorf("invalid api token")
	}

	if strings.Contains(tokenOriginAddress, "//") {
		tokenOriginAddress = strings.Split(tokenOriginAddress, "//")[1]
	}

	return &credentials{
		apiTokenStr:        apiTokenStr,
		tokenOriginAddress: tokenOriginAddress,
	}, nil
}

type NeptuneManagementAuthenticator struct {
	authToken string
}

func (nma *NeptuneManagementAuthenticator) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	return req.SetHeaderParam("authorization", fmt.Sprintf("Bearer %s", nma.authToken))
}

type DefaultAuthenticator struct {
}

func (da *DefaultAuthenticator) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	return nil
}
