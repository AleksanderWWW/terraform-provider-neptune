package neptune

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProjectMember(t *testing.T) {
	type testCase struct {
		label                              string
		resps                              []*http.Response
		errs                               []error
		expectedError                      error
		project, workspace, username, role string
	}

	for _, tc := range []testCase{
		{
			label:         "test invalid role",
			resps:         []*http.Response{},
			expectedError: fmt.Errorf("unknown role 'TestRole'"),
			role:          "TestRole",
		},
		{
			label: "token extraction error",
			resps: []*http.Response{
				{
					StatusCode: 500,
					Body: &MockReadCloser{
						data: "Test Error (token)",
					},
				},
			},
			errs:          []error{nil},
			expectedError: fmt.Errorf("500: Test Error (token)"),
			role:          "member",
		},
		{
			label: "member adding error",
			resps: []*http.Response{
				{
					StatusCode: 200,
					Body: &MockReadCloser{
						data: "{\"accessToken\":\"someToken\"}",
					},
				},
				{
					StatusCode: 500,
					Body: &MockReadCloser{
						data: "Test Error (adding member)",
					},
				},
			},
			errs:          []error{nil, nil},
			expectedError: fmt.Errorf("500: Test Error (adding member)"),
			role:          "member",
		},
		{
			label: "request failure",
			resps: []*http.Response{
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
			},
			errs:          []error{nil, fmt.Errorf("error sending request")},
			expectedError: fmt.Errorf("error sending request"),
			role:          "member",
		},
		{
			label: "success adding member",
			resps: []*http.Response{
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
			},
			errs:          []error{nil, nil},
			expectedError: nil,
			role:          "member",
		},
	} {
		client := MockHttpClient{
			resps: tc.resps,
			errs:  tc.errs,
		}

		nptClient := NeptuneClient{
			httpClient: &client,
			creds:      testCreds,
		}

		err := nptClient.AddProjectMember(tc.project, tc.workspace, tc.username, tc.role)
		if tc.expectedError != nil {
			assert.EqualError(t, err, tc.expectedError.Error())
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestDeleteProjectMember(t *testing.T) {
	type testCase struct {
		label                        string
		resps                        []*http.Response
		errs                         []error
		expectedError                error
		project, workspace, username string
	}

	for _, tc := range []testCase{
		{
			label: "token extraction error",
			resps: []*http.Response{
				{
					StatusCode: 500,
					Body: &MockReadCloser{
						data: "Test Error (token)",
					},
				},
			},
			errs:          []error{nil},
			expectedError: fmt.Errorf("500: Test Error (token)"),
		},
		{
			label: "member deleting error",
			resps: []*http.Response{
				{
					StatusCode: 200,
					Body: &MockReadCloser{
						data: "{\"accessToken\":\"someToken\"}",
					},
				},
				{
					StatusCode: 500,
					Body: &MockReadCloser{
						data: "Test Error (deleting member)",
					},
				},
			},
			errs:          []error{nil, nil},
			expectedError: fmt.Errorf("500: Test Error (deleting member)"),
		},
		{
			label: "request failure",
			resps: []*http.Response{
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
			},
			errs:          []error{nil, fmt.Errorf("Error sending request")},
			expectedError: fmt.Errorf("Error sending request"),
		},
		{
			label: "success deleting member",
			resps: []*http.Response{
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
			},
			errs:          []error{nil, nil},
			expectedError: nil,
		},
	} {
		client := MockHttpClient{
			resps: tc.resps,
			errs:  tc.errs,
		}

		nptClient := NeptuneClient{
			httpClient: &client,
			creds:      testCreds,
		}

		err := nptClient.DeleteProjectMember(tc.project, tc.workspace, tc.username)
		if tc.expectedError != nil {
			assert.EqualError(t, err, tc.expectedError.Error())
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestUpdateProjectMember(t *testing.T) {
	type testCase struct {
		label                              string
		resps                              []*http.Response
		errs                               []error
		expectedError                      error
		project, workspace, username, role string
	}

	for _, tc := range []testCase{
		{
			label:         "test invalid role",
			resps:         []*http.Response{},
			expectedError: fmt.Errorf("unexpected 'role' argument value: 'TestRole'"),
			role:          "TestRole",
		},
		{
			label: "token extraction error",
			resps: []*http.Response{
				{
					StatusCode: 500,
					Body: &MockReadCloser{
						data: "Test Error (token)",
					},
				},
			},
			errs:          []error{nil},
			expectedError: fmt.Errorf("500: Test Error (token)"),
			role:          "member",
		},
		{
			label: "member updating error",
			resps: []*http.Response{
				{
					StatusCode: 200,
					Body: &MockReadCloser{
						data: "{\"accessToken\":\"someToken\"}",
					},
				},
				{
					StatusCode: 500,
					Body: &MockReadCloser{
						data: "Test Error (updating member)",
					},
				},
			},
			errs:          []error{nil, nil},
			expectedError: fmt.Errorf("500: Test Error (updating member)"),
			role:          "member",
		},
		{
			label: "request failure",
			resps: []*http.Response{
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
			},
			errs:          []error{nil, fmt.Errorf("Error sending request")},
			expectedError: fmt.Errorf("Error sending request"),
			role:          "member",
		},
		{
			label: "success updating member",
			resps: []*http.Response{
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
			},
			errs:          []error{nil, nil},
			expectedError: nil,
			role:          "member",
		},
	} {
		client := MockHttpClient{
			resps: tc.resps,
			errs:  tc.errs,
		}

		nptClient := NeptuneClient{
			httpClient: &client,
			creds:      testCreds,
		}

		err := nptClient.UpdateProjectMember(tc.project, tc.workspace, tc.username, tc.role)
		if tc.expectedError != nil {
			assert.EqualError(t, err, tc.expectedError.Error())
		} else {
			assert.NoError(t, err)
		}
	}
}
