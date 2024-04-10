package neptune

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type credentials struct {
	apiToken           string
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

func NewCredentials(apiToken string) (*credentials, error) {
	tokenDict, err := decodeAPIToken(apiToken)
	if err != nil {
		return nil, err
	}

	tokenOriginAddress, ok := tokenDict["api_address"]
	if !ok {
		return nil, fmt.Errorf("invalid api token")
	}

	tokenOriginAddress = strings.Split(tokenOriginAddress, "//")[1]

	return &credentials{
		apiToken:           apiToken,
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
