package neptune

import (
	"encoding/json"
	"fmt"
	"strings"
)

type projectData struct {
	projectIdentifier string
	key               *string
	vis               string
}

func buildCreateProjectRequestData(authToken, name, workspaceId string, project projectData) (*requestData, error) {
	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	body, err := formProjectBody(name, workspaceId, project.vis, project.key)
	if err != nil {
		return &requestData{}, err
	}

	return &requestData{
		headers:  headers,
		params:   nil,
		bodyJson: body,
	}, nil
}

func buildDeleteProjectRequestData(authToken, projectIdentifier string) *requestData {
	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	params := map[string]string{
		"projectIdentifier": projectIdentifier,
	}

	return &requestData{headers: headers, params: params, bodyJson: nil}
}

func verifyCreateProjectArgs(name string, workspace string, key string, vis string) (projectData, error) {
	if len(name) < 3 {
		return projectData{}, fmt.Errorf("project name length should be at least 3 characters. Got %d", len(name))
	}

	var optionalKey *string

	if key == "" {
		optionalKey = nil
	} else {
		optionalKey = &key
	}

	if vis == "" {
		vis = "private"
	}

	_vis, ok := stringToVisibility[strings.ToLower(vis)]
	if !ok {
		return projectData{}, fmt.Errorf(
			"unsupported visibility type '%s'. Available choices (case insensitive) are: 'private', 'public' and 'workspace'",
			vis,
		)
	}

	return projectData{
		projectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
		key:               optionalKey,
		vis:               _vis,
	}, nil
}

func (c *NeptuneClient) CreateProject(name string, workspace string, key string, vis string) error {
	project, err := verifyCreateProjectArgs(name, workspace, key, vis)
	if err != nil {
		return err
	}

	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	workspaceId, err := c.getWorkspaceId(authToken, workspace)
	if err != nil {
		return err
	}

	data, err := buildCreateProjectRequestData(authToken, name, workspaceId, project)
	if err != nil {
		return err
	}

	resp, err := c.do("createProject", data)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		responseString, _ := getResponseMessage(resp)
		return fmt.Errorf(responseString)
	}

	return nil
}

func formProjectBody(name string, workspaceId string, vis string, key *string) ([]byte, error) {
	type body struct {
		Name           string  `json:"name"`
		OrganizationId string  `json:"organizationId"`
		Visibility     string  `json:"visibility"`
		ProjectKey     *string `json:"projectKey"`
	}
	b := body{
		Name:           name,
		OrganizationId: workspaceId,
		Visibility:     vis,
		ProjectKey:     key,
	}
	return json.Marshal(b)
}

func (c *NeptuneClient) DeleteProject(name string, workspace string) error {
	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	projectIdentifier := fmt.Sprintf("%s/%s", workspace, name)

	data := buildDeleteProjectRequestData(authToken, projectIdentifier)

	resp, err := c.do("deleteProject", data)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		responseString, _ := getResponseMessage(resp)
		return fmt.Errorf(responseString)
	}

	return nil
}
