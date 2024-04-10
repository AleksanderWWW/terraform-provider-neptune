package neptune

import (
	"fmt"
	"terraform-provider-neptune/internal/neptune/client"
	"terraform-provider-neptune/internal/neptune/client/operations"
	"terraform-provider-neptune/internal/neptune/models"

	"github.com/go-openapi/strfmt"
)

const NeptuneApiToken string = "NEPTUNE_API_TOKEN"

type NeptuneClient struct {
	apiClient *client.Neptune
	creds     *credentials
}

func NewNeptuneClient(creds *credentials) *NeptuneClient {
	return &NeptuneClient{
		apiClient: client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{
			Host:    creds.tokenOriginAddress,
			Schemes: []string{"https"},
		}),
		creds: creds,
	}
}

func (nc *NeptuneClient) NewManagementAuthenticator() (*NeptuneManagementAuthenticator, error) {
	auth := &DefaultAuthenticator{}
	authTokenDTO, err := nc.apiClient.Operations.ExchangeAPIToken(
		&operations.ExchangeAPITokenParams{
			XNeptuneAPIToken: nc.creds.apiToken,
		},
		auth,
	)
	if err != nil {
		return nil, err
	}
	authToken := authTokenDTO.Payload.AccessToken
	return &NeptuneManagementAuthenticator{authToken: *authToken}, nil
}

func (nc *NeptuneClient) getWorkspaceID(name string) (strfmt.UUID, error) {
	auth, err := nc.NewManagementAuthenticator()
	if err != nil {
		return "", err
	}

	organizationDTOs, err := nc.apiClient.Operations.ListOrganizations(&operations.ListOrganizationsParams{}, auth)
	if err != nil {
		return "", err
	}

	for _, dto := range organizationDTOs.Payload {
		if *dto.Organization.Name == name {
			return *dto.Organization.ID, nil
		}
	}

	return "", fmt.Errorf("Workspace '%s' not found", name)
}

func (nc *NeptuneClient) CreateProject(name, workspace, visibility string) error {
	visDTO := getVisibilityDTO(visibility)

	auth, err := nc.NewManagementAuthenticator()
	if err != nil {
		return err
	}

	orgId, err := nc.getWorkspaceID(workspace)
	if err != nil {
		return err
	}

	params := operations.CreateProjectParams{
		ProjectToCreate: &models.NewProjectDTO{
			Name:           &name,
			OrganizationID: &orgId,
			Visibility:     *models.NewProjectVisibilityDTO(visDTO),
		},
	}
	_, err = nc.apiClient.Operations.CreateProject(&params, auth)
	return err
}

func getVisibilityDTO(visibility string) models.ProjectVisibilityDTO {
	var visDTO models.ProjectVisibilityDTO

	switch visibility {
	case "private":
		visDTO = models.ProjectVisibilityDTOPriv
	case "public":
		visDTO = models.ProjectVisibilityDTOPub
	case "workspace":
		visDTO = models.ProjectVisibilityDTOWorkspace
	default:
		visDTO = models.ProjectVisibilityDTOPriv
	}

	return visDTO
}

func (nc *NeptuneClient) DeleteProject(name, workspace string) error {
	auth, err := nc.NewManagementAuthenticator()
	if err != nil {
		return err
	}

	params := &operations.DeleteProjectParams{
		ProjectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
	}

	_, err = nc.apiClient.Operations.DeleteProject(params, auth)
	return err
}

func (nc *NeptuneClient) AddProjectMember(name, workspace, username, role string) error {
	auth, err := nc.NewManagementAuthenticator()
	if err != nil {
		return err
	}

	roleDTO := getProjectMemberRoleDTO(role)

	memberDTO := &models.NewProjectMemberDTO{
		Role:   &roleDTO,
		UserID: &username,
	}

	params := &operations.AddProjectMemberParams{
		Member:            memberDTO,
		ProjectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
	}

	_, err = nc.apiClient.Operations.AddProjectMember(params, auth)
	return err
}

func getProjectMemberRoleDTO(role string) models.ProjectRoleDTO {
	switch role {
	case "member":
		return models.ProjectRoleDTOMember
	case "viewer":
		return models.ProjectRoleDTOViewer
	case "manager":
		return models.ProjectRoleDTOManager
	default:
		return models.ProjectRoleDTOMember
	}

}

func (nc *NeptuneClient) UpdateProjectMember(name, workspace, username, role string) error {
	auth, err := nc.NewManagementAuthenticator()
	if err != nil {
		return err
	}

	roleDTO := getProjectMemberRoleDTO(role)

	member := models.ProjectMemberUpdateDTO{
		Role: &roleDTO,
	}

	params := &operations.UpdateProjectMemberParams{
		Member:            &member,
		ProjectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
		UserID:            username,
	}

	_, err = nc.apiClient.Operations.UpdateProjectMember(params, auth)
	return err
}

func (nc *NeptuneClient) DeleteProjectMember(name, workspace, username string) error {
	auth, err := nc.NewManagementAuthenticator()
	if err != nil {
		return err
	}

	params := &operations.DeleteProjectMemberParams{
		ProjectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
		UserID:            username,
	}

	_, err = nc.apiClient.Operations.DeleteProjectMember(params, auth)
	return err
}
