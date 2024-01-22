package neptune

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialsSuccess(t *testing.T) {
	jsonData, err := json.Marshal(map[string]string{"api_address": "someAddress"})
	assert.NoError(t, err)

	// Encode to base64
	apiToken := base64.StdEncoding.EncodeToString(jsonData)

	creds, err := NewCredentials(apiToken)

	assert.NoError(t, err)

	assert.Equal(t, creds.tokenOriginAddress, "someAddress")
	assert.Equal(t, creds.apiToken, apiToken)
}

func TestCredentialsFailure(t *testing.T) {
	jsonData, err := json.Marshal(map[string]string{})
	assert.NoError(t, err)

	// Encode to base64
	apiToken := base64.StdEncoding.EncodeToString(jsonData)

	creds, err := NewCredentials(apiToken)

	assert.Error(t, err)
	assert.Nil(t, creds)
}

func TestCredentialsWrongToken(t *testing.T) {
	creds, err := NewCredentials("wrongToken")

	assert.Error(t, err)
	assert.Nil(t, creds)

	os.Unsetenv(NeptuneApiToken)
	creds, err = NewCredentials("")

	assert.Error(t, err)
	assert.Nil(t, creds)
}
