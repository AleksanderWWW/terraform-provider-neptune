package neptune

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *NeptuneClient) CreateProject(name string, workspace string, key string, vis string) error {
	if len(name) < 3 {
		return fmt.Errorf("Project name length should be at least 3 characters. Got %d", len(name))
	}

	if key == "" {
		key = strings.ToUpper(name[:3])
	}

	if vis == "" {
		vis = "private"
	}

	_vis, ok := stringToVisibility[strings.ToLower(vis)]
	if !ok {
		return fmt.Errorf("Unsupported visibility type '%s'. Available choices (case insensitive) are: 'private', 'public' and 'workspace'", vis)
	}

	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	workspaceId, err := c.getWorkspaceId(authToken, workspace)
	if err != nil {
		return err
	}
	body, err := formProjectBody(name, workspaceId, _vis, key)
	if err != nil {
		return err
	}

	resp, err := c.do("createProject", nil, headers, body)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: response status code %d", resp.StatusCode)
	}

	return nil
}

func formProjectBody(name string, workspaceId string, vis string, key string) ([]byte, error) {
	type body struct {
		Name           string `json:"name"`
		OrganizationId string `json:"organizationId"`
		Visibility     string `json:"visibility"`
		ProjectKey     string `json:"projectKey"`
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

	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	params := map[string]string{
		"projectIdentifier": projectIdentifier,
	}

	resp, err := c.do("deleteProject", params, headers, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode == 404 {
		return fmt.Errorf("Error: project '%s' does not exist", projectIdentifier)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: response status code %d", resp.StatusCode)
	}

	return nil
}
