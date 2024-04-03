package neptune

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAllowSelfSignedCertificateEnvSet(t *testing.T) {
	defer os.Unsetenv(NeptuneAllowSelfSignedCcertificate)

	type testCase struct {
		envVal string
		result bool
	}

	testCases := []testCase{
		{
			"True",
			true,
		},
		{
			"true",
			true,
		},
		{
			"tRUe",
			true,
		},
		{
			"False",
			false,
		},
		{
			"",
			false,
		},
	}

	for _, tc := range testCases {
		os.Setenv(NeptuneAllowSelfSignedCcertificate, tc.envVal)
		assert.Equal(t, shouldAllowSelfSignedCertificates(), tc.result)
	}
}

func TestShouldAllowSelfSignedCertificateEnvNotSet(t *testing.T) {
	os.Unsetenv(NeptuneAllowSelfSignedCcertificate)
	assert.Equal(t, shouldAllowSelfSignedCertificates(), false)
}
