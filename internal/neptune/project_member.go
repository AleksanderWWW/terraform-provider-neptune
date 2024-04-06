package neptune

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type addProjectMemberRequestBody struct {
	Role   string `json:"role"`
	UserId string `json:"userId"`
}

type deleteProjectMemberRequestBody struct {
	UserId string `json:"userId"`
}

type updateProjectMemberRequestBody struct {
	Role string `json:"role"`
}

func buildAddProjectMemberRequestData(authToken, projectIdentifier, username, role string) (*requestData, error) {
	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	params := map[string]string{
		"projectIdentifier": projectIdentifier,
	}

	body := addProjectMemberRequestBody{
		Role:   role,
		UserId: username,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return &requestData{}, err
	}

	return newRequestData(headers, params, bodyJson), nil
}

func buildDeleteProjectMemberRequestData(authToken, projectIdentifier, username string) (*requestData, error) {
	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	params := map[string]string{
		"projectIdentifier": projectIdentifier,
	}

	body := deleteProjectMemberRequestBody{
		UserId: username,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return &requestData{}, err
	}

	return newRequestData(headers, params, bodyJson), nil
}

func buildUpdateProjectMemberRequestData(authToken, projectIdentifier, username, role string) (*requestData, error) {
	headers := map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", authToken),
	}

	params := map[string]string{
		"projectIdentifier": projectIdentifier,
	}

	body := updateProjectMemberRequestBody{
		Role: role,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return &requestData{}, err
	}

	return newRequestData(headers, params, bodyJson), nil
}

func (c *NeptuneClient) AddProjectMember(project, workspace, username, role string) error {
	if _, ok := roles[role]; !ok {
		return fmt.Errorf("unknown role '%s'", role)
	}

	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	projectIdentifier := buildProjectIdentifier(workspace, project)

	data, err := buildAddProjectMemberRequestData(authToken, projectIdentifier, username, role)

	if err != nil {
		return err
	}

	resp, err := c.do("addProjectMember", data)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	}

	return handleResponseError(resp)
}

func (c *NeptuneClient) DeleteProjectMember(project string, workspace string, username string) error {
	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	projectIdentifier := fmt.Sprintf("%s/%s", workspace, project)

	data, err := buildDeleteProjectMemberRequestData(authToken, projectIdentifier, username)

	if err != nil {
		return err
	}

	var opData operationData
	err = viper.UnmarshalKey("deleteProjectMember", &opData)
	if err != nil {
		return err
	}
	endpoint := opData.Endpoint + "/" + username

	req, err := prepareRequest(c.creds.tokenOriginAddress, endpoint, opData.Method, data)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	}

	return handleResponseError(resp)
}

func (c *NeptuneClient) UpdateProjectMember(project string, workspace string, username string, role string) error {
	if _, ok := roles[role]; !ok {
		return fmt.Errorf("unexpected 'role' argument value: '%s'", role)
	}

	authToken, err := c.getAuthToken()
	if err != nil {
		return err
	}

	projectIdentifier := buildProjectIdentifier(workspace, project)

	data, err := buildUpdateProjectMemberRequestData(authToken, projectIdentifier, username, role)
	if err != nil {
		return err
	}

	var opData operationData
	err = viper.UnmarshalKey("updateProjectMember", &opData)
	if err != nil {
		return err
	}
	endpoint := opData.Endpoint + "/" + username

	req, err := prepareRequest(c.creds.tokenOriginAddress, endpoint, opData.Method, data)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	}

	return handleResponseError(resp)
}
