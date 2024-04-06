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

	data := &requestData{headers: headers, params: nil, bodyJson: nil}

	resp, err := c.do("auth", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", handleResponseError(resp)
	}

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

	data := &requestData{headers: headers, params: nil, bodyJson: nil}

	resp, err := c.do("listOrganizations", data)

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
	return "", fmt.Errorf("workspace '%s' not found", workspaceName)
}
