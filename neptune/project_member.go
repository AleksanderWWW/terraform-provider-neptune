package neptune

import (
	"encoding/json"
	"fmt"
)

func (c *NeptuneClient) AddProjectMember(project string, workspace string, username string, role string) error {
	if _, ok := roles[role]; !ok {
		return fmt.Errorf("Unknown role '%s'", role)
	}

	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	projectIdentifier := fmt.Sprintf("%s/%s", workspace, project)

	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	params := map[string]string{
		"projectIdentifier": projectIdentifier,
	}

	type reqBody struct {
		Role   string `json:"role"`
		UserId string `json:"userId"`
	}

	body := reqBody{
		Role:   role,
		UserId: username,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := c.do("addProjectMember", params, headers, bodyJson)
	if err != nil {
		return err
	}

	if resp.StatusCode == 409 {
		return fmt.Errorf("User '%s' already exists in project '%s", username, projectIdentifier)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: response status code %d", resp.StatusCode)
	}
	return nil
}

func (c *NeptuneClient) DeleteProjectMember(project string, workspace string, username string) error {
	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	projectIdentifier := fmt.Sprintf("%s/%s", workspace, project)

	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	params := map[string]string{
		"projectIdentifier": projectIdentifier,
	}

	type reqBody struct {
		UserId string `json:"userId"`
	}

	body := reqBody{
		UserId: username,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	opData := config["deleteProjectMember"]
	endpoint := opData.Endpoint + "/" + username

	req, err := prepareRequest(c.creds.tokenOriginAddress, endpoint, opData.Method, params, headers, bodyJson)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 422 {
		return fmt.Errorf("User '%s' not found in project '%s", username, projectIdentifier)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error: response status code %d", resp.StatusCode)
	}
	return nil
}
