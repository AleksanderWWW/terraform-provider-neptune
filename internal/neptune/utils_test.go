package neptune

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAllowSelfSignedCertificateEnvSet(t *testing.T) {
	defer os.Unsetenv(NeptuneAllowSelfSignedCertificate)

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
			"1",
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
		{
			"0",
			false,
		},
	}

	for _, tc := range testCases {
		os.Setenv(NeptuneAllowSelfSignedCertificate, tc.envVal)
		assert.Equal(t, shouldAllowSelfSignedCertificates(), tc.result)
	}
}

func TestShouldAllowSelfSignedCertificateEnvNotSet(t *testing.T) {
	os.Unsetenv(NeptuneAllowSelfSignedCertificate)
	assert.Equal(t, shouldAllowSelfSignedCertificates(), false)
}
