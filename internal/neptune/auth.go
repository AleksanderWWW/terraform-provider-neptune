package neptune

import "fmt"

type credentials struct {
	apiToken           string
	tokenOriginAddress string
}

func NewCredentials(apiToken string) (*credentials, error) {
	apiToken, err := getApiToken(apiToken)

	if err != nil {
		return nil, err
	}

	tokenDict, err := decodeAPIToken(apiToken)
	if err != nil {
		return nil, err
	}

	tokenOriginAddress, ok := tokenDict["api_address"]
	if !ok {
		return nil, fmt.Errorf("invalid api token")
	}

	return &credentials{
		apiToken:           apiToken,
		tokenOriginAddress: tokenOriginAddress,
	}, nil
}
