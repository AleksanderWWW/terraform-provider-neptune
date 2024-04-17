package neptune

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialsSuccess(t *testing.T) {
	jsonData, err := json.Marshal(map[string]string{
		"api_address": "https://someAddress",
	})
	assert.NoError(t, err)

	// Encode to base64
	apiToken := base64.StdEncoding.EncodeToString(jsonData)

	creds, err := newCredentials(&apiToken)

	assert.NoError(t, err)

	assert.Equal(t, creds.tokenOriginAddress, "someAddress")
	assert.Equal(t, creds.apiTokenStr, apiToken)
}

func TestCredentialsFailure(t *testing.T) {
	jsonData, err := json.Marshal(map[string]string{})
	assert.NoError(t, err)

	// Encode to base64
	apiToken := base64.StdEncoding.EncodeToString(jsonData)

	creds, err := newCredentials(&apiToken)

	assert.Error(t, err)
	assert.Nil(t, creds)
}

func TestCredentialsWrongToken(t *testing.T) {
	token := "wrongToken"
	creds, err := newCredentials(&token)

	assert.Error(t, err)
	assert.Nil(t, creds)

	os.Unsetenv(NeptuneApiToken)
	creds, err = newCredentials(nil)

	assert.Error(t, err)
	assert.Nil(t, creds)
}
