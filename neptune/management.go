package neptune

import (
	"encoding/json"
	"fmt"
)

func (c *NeptuneClient) getAuthToken() (string, error) {
	type respData struct {
		AccessToken string `json:"accessToken"`
	}
	var respJson respData

	headers := map[string]string{
		"X-Neptune-Api-Token": c.creds.apiToken,
	}

	resp, err := c.do("auth", nil, headers, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respJson)
	if err != nil {
		return "", err
	}

	return respJson.AccessToken, nil
}

func (c *NeptuneClient) getWorkspaceId(authToken string, workspaceName string) (string, error) {
	type workspaceInfo struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}

	var info []workspaceInfo

	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	resp, err := c.do("listOrganizations", nil, headers, nil)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return "", err
	}

	for _, workspace := range info {
		if workspace.Name == workspaceName {
			return workspace.Id, nil
		}
	}
	return "", fmt.Errorf("Workspace '%s' not found", workspaceName)
}

func (c *NeptuneClient) CreateProject(name string, workspace string, key string, vis string) error {
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
	body, err := formProjectBody(name, workspaceId, vis, key)
	if err != nil {
		return err
	}

	resp, err := c.do("createProject", nil, headers, body)

	if err != nil {
		return err
	}

	if resp.StatusCode == 422 {
		return fmt.Errorf("Error: project '%s' already exists", name)
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
