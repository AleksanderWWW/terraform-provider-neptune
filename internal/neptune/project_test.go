package neptune

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var creds credentials = credentials{
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

	_, err = verifyCreateProjectArgs("SomeProject", "SomeWorkspace", "WrongKey", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key must either be an empty, or a 3-letter string")

	_, err = verifyCreateProjectArgs("SomeProject", "SomeWorkspace", "WK", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Key must either be an empty, or a 3-letter string")

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

func executeCreateProject(resps []*http.Response, creds credentials) error {
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, nil},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      creds,
	}

	return nptClient.CreateProject("Name1", "Workspace1", "", "")
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
	err := executeCreateProject(resps, creds)
	assert.NoError(t, err)
}

func TestCreateProjectIfExists(t *testing.T) {
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
				data: "",
			},
		},
	}
	err := executeCreateProject(resps, creds)
	assert.EqualError(t, err, "Error: project 'Workspace1/Name1' already exists")
}

func TestCreateProjectLimitExceeded(t *testing.T) {
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
			StatusCode: 422,
			Body: &MockReadCloser{
				data: "",
			},
		},
	}
	err := executeCreateProject(resps, creds)
	assert.EqualError(t, err, "Error: project limit exceeded")
}

func TestCreateProjectLInvalidStatusCode(t *testing.T) {
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
			StatusCode: 500,
			Body: &MockReadCloser{
				data: "",
			},
		},
	}
	err := executeCreateProject(resps, creds)
	assert.EqualError(t, err, "Error: response status code 500")
}

func TestCreateProjectWorkspaceNotFound(t *testing.T) {
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
				data: "[{\"name\": \"Workspace42\", \"id\":\"someId\"}]",
			},
		},
		{
			StatusCode: 409,
			Body: &MockReadCloser{
				data: "",
			},
		},
	}
	err := executeCreateProject(resps, creds)
	assert.EqualError(t, err, "Workspace 'Workspace1' not found")
}

func TestCreateProjectErrorWithResponse(t *testing.T) {
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
				data: "",
			},
		},
	}
	client := MockHttpClient{
		resps: resps,
		errs:  []error{nil, nil, fmt.Errorf("Some custom error here")},
	}

	nptClient := NeptuneClient{
		httpClient: &client,
		creds:      creds,
	}
	err := nptClient.CreateProject("Name1", "Workspace1", "", "")
	assert.EqualError(t, err, "Some custom error here")
}
