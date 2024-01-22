package neptune

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCreds credentials = credentials{
	apiToken:           "someAPIToken",
	tokenOriginAddress: "someAddress",
}

func TestFormProjectBody(t *testing.T) {
	type projectBody struct {
		name        string
		workspaceId string
		visibility  string
		projectKey  string
	}
	type testCase struct {
		pb           projectBody
		expectedOut  string
		expectingErr bool
	}

	for _, test := range []testCase{
		{
			pb: projectBody{
				"Name1",
				"SomeID1",
				"priv",
				"NAM",
			},
			expectedOut:  "{\"name\":\"Name1\",\"organizationId\":\"SomeID1\",\"visibility\":\"priv\",\"projectKey\":\"NAM\"}",
			expectingErr: false,
		},
		{
			pb:           projectBody{},
			expectedOut:  "{\"name\":\"\",\"organizationId\":\"\",\"visibility\":\"\",\"projectKey\":\"\"}",
			expectingErr: false,
		},
	} {
		body, err := formProjectBody(test.pb.name, test.pb.workspaceId, test.pb.visibility, test.pb.projectKey)
		if test.expectingErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

		assert.Equal(t, string(body[:]), test.expectedOut)
	}
}

func TestVerifyCreateProjectArgs(t *testing.T) {
	_, err := verifyCreateProjectArgs("", "some-workspace", "", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Project name length should be at least 3 characters.")

	_, err = verifyCreateProjectArgs("SomeProject", "SomeWorkspace", "", "Wrong-visibility")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Unsupported visibility type")

	project, err := verifyCreateProjectArgs("SomeName", "SomeWorkspace", "", "")
	assert.NoError(t, err)
	assert.Equal(
		t,
		project,
		projectData{
			"SomeWorkspace/SomeName",
			"SOM",
			"priv",
		},
	)
}

func TestCreateProjectInvalidResponseCode(t *testing.T) {
	resps := []*http.Response{
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "{\"accessToken\":\"someToken\"}",
			},
		},
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "[{\"name\": \"Workspace1\", \"id\":\"someId\"}]",
			},
		},
		{
			StatusCode: 409,
			Body: &MockReadCloser{
				data: "Test error",
			},
		},
	}
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, nil},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      testCreds,
	}
	err := nptClient.CreateProject("Name1", "Workspace1", "", "")
	assert.EqualError(t, err, "409: Test error")
}

func TestCreateProjectSuccess(t *testing.T) {
	resps := []*http.Response{
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "{\"accessToken\":\"someToken\"}",
			},
		},
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "[{\"name\": \"Workspace1\", \"id\":\"someId\"}]",
			},
		},
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "",
			},
		},
	}
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, nil},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      testCreds,
	}
	err := nptClient.CreateProject("Name1", "Workspace1", "", "")
	assert.NoError(t, err)
}

func TestDeleteProjectSuccess(t *testing.T) {
	resps := []*http.Response{
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "{\"accessToken\":\"someToken\"}",
			},
		},
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "",
			},
		},
	}
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, nil},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      testCreds,
	}
	err := nptClient.DeleteProject("Name1", "Workspace1")
	assert.NoError(t, err)
}

func TestDeleteProjectError(t *testing.T) {
	resps := []*http.Response{
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "{\"accessToken\":\"someToken\"}",
			},
		},
		{
			StatusCode: 404,
			Body: &MockReadCloser{
				data: "Test error",
			},
		},
	}
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, nil},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      testCreds,
	}
	err := nptClient.DeleteProject("Name1", "Workspace1")
	assert.EqualError(t, err, "404: Test error")
}

func TestCreateProjectTokenError(t *testing.T) {

	resps := []*http.Response{
		{
			StatusCode: 500,
			Body: &MockReadCloser{
				data: "Test error",
			},
		},
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "[{\"name\": \"Workspace1\", \"id\":\"someId\"}]",
			},
		},
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "",
			},
		},
	}
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, nil},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      testCreds,
	}
	err := nptClient.CreateProject("Name1", "Workspace1", "", "")
	assert.EqualError(t, err, "500: Test error")
}

func TestDeleteProjectTokenError(t *testing.T) {

	resps := []*http.Response{
		{
			StatusCode: 500,
			Body: &MockReadCloser{
				data: "Test error",
			},
		},
		{
			StatusCode: 200,
			Body: &MockReadCloser{
				data: "",
			},
		},
	}
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, nil},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      testCreds,
	}
	err := nptClient.DeleteProject("Name1", "Workspace1")
	assert.EqualError(t, err, "500: Test error")
}
