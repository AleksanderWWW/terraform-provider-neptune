package neptune

import (
	"terraform-provider-neptune/internal/neptune/client"
	"terraform-provider-neptune/internal/neptune/client/operations"

	"github.com/go-openapi/strfmt"
)

type neptuneApiClient struct {
	backendClient *client.Neptune
	auth          *NeptuneManagementAuthenticator
}

func newNeptuneApiClient(apiToken *string) (*neptuneApiClient, error) {
	c, err := newCredentials(apiToken)
	if err != nil {
		return nil, err
	}

	client := &neptuneApiClient{
		backendClient: client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{
			Host:    c.tokenOriginAddress,
			Schemes: []string{"https"},
		}),
		auth: nil,
	}

	auth, err := newManagementAuthenticator(client, c.apiTokenStr)
	if err != nil {
		return nil, err
	}

	client.auth = auth

	return client, nil
}

func (ac *neptuneApiClient) getAuthToken(apiTokenStr string) (*string, error) {
	auth := &DefaultAuthenticator{}

	authTokenDTO, err := ac.backendClient.Operations.ExchangeAPIToken(
		&operations.ExchangeAPITokenParams{
			XNeptuneAPIToken: apiTokenStr,
		},
		auth,
	)
	if err != nil {
		return nil, err
	}

	authToken := authTokenDTO.Payload.AccessToken
	return authToken, nil
}

func newManagementAuthenticator(apiClient *neptuneApiClient, apiTokenStr string) (*NeptuneManagementAuthenticator, error) {
	authToken, err := apiClient.getAuthToken(apiTokenStr)
	if err != nil {
		return nil, err
	}

	return &NeptuneManagementAuthenticator{authToken: *authToken}, nil
}
