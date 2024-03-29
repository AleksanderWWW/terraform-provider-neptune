package neptune

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func (c *NeptuneClient) AddProjectMember(project, workspace, username, role string) error {
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

	if resp.StatusCode != 200 {
		responseString, _ := getResponseMessage(resp)
		return fmt.Errorf(responseString)
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

	var opData operationData
	err = viper.UnmarshalKey("deleteProjectMember", &opData)
	if err != nil {
		return err
	}
	endpoint := opData.Endpoint + "/" + username

	req, err := prepareRequest(c.creds.tokenOriginAddress, endpoint, opData.Method, params, headers, bodyJson)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		responseString, _ := getResponseMessage(resp)
		return fmt.Errorf(responseString)
	}

	return nil
}

func (c *NeptuneClient) UpdateProjectMember(project string, workspace string, username string, role string) error {
	role = strings.ToLower(role)
	if _, ok := map[string]bool{
		"viewer":  true,
		"member":  true,
		"manager": true,
	}[role]; !ok {
		return fmt.Errorf("Unexpected 'role' argument value: '%s'", role)
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
		Role string `json:"role"`
	}

	body := reqBody{
		Role: role,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	var opData operationData
	err = viper.UnmarshalKey("updateProjectMember", &opData)
	if err != nil {
		return err
	}
	endpoint := opData.Endpoint + "/" + username

	req, err := prepareRequest(c.creds.tokenOriginAddress, endpoint, opData.Method, params, headers, bodyJson)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		responseString, _ := getResponseMessage(resp)
		return fmt.Errorf(responseString)
	}

	return nil
}
