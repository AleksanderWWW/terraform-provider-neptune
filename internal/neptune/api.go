package neptune

import (
	"fmt"
	"terraform-provider-neptune/internal/neptune/client/operations"
	"terraform-provider-neptune/internal/neptune/models"

	"github.com/go-openapi/strfmt"
)

func getWorkspaceID(name string, apiClient *neptuneApiClient) (strfmt.UUID, error) {
	organizationDTOs, err := apiClient.backendClient.Operations.ListOrganizations(
		&operations.ListOrganizationsParams{}, apiClient.auth,
	)
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

func CreateProject(name, workspace, visibility string, apiToken *string) error {
	apiClient, err := newNeptuneApiClient(apiToken)
	if err != nil {
		return err
	}

	visDTO := getVisibilityDTO(visibility)

	orgId, err := getWorkspaceID(workspace, apiClient)
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
	_, err = apiClient.backendClient.Operations.CreateProject(&params, apiClient.auth)
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

func DeleteProject(name, workspace string, apiToken *string) error {
	apiClient, err := newNeptuneApiClient(apiToken)
	if err != nil {
		return err
	}

	params := &operations.DeleteProjectParams{
		ProjectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
	}

	_, err = apiClient.backendClient.Operations.DeleteProject(params, apiClient.auth)
	return err
}

func AddProjectMember(name, workspace, username, role string, apiToken *string) error {
	apiClient, err := newNeptuneApiClient(apiToken)
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

	_, err = apiClient.backendClient.Operations.AddProjectMember(params, apiClient.auth)
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

func UpdateProjectMember(name, workspace, username, role string, apiToken *string) error {
	apiClient, err := newNeptuneApiClient(apiToken)
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

	_, err = apiClient.backendClient.Operations.UpdateProjectMember(params, apiClient.auth)
	return err
}

func DeleteProjectMember(name, workspace, username string, apiToken *string) error {
	apiClient, err := newNeptuneApiClient(apiToken)
	if err != nil {
		return err
	}

	params := &operations.DeleteProjectMemberParams{
		ProjectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
		UserID:            username,
	}

	_, err = apiClient.backendClient.Operations.DeleteProjectMember(params, apiClient.auth)
	return err
}

func ListProjectMembers(name, workspace string, apiToken *string) ([]*models.ProjectMemberDTO, error) {
	apiClient, err := newNeptuneApiClient(apiToken)
	if err != nil {
		return nil, err
	}

	params := &operations.ListProjectMembersParams{
		ProjectIdentifier: fmt.Sprintf("%s/%s", workspace, name),
	}

	res, err := apiClient.backendClient.Operations.ListProjectMembers(params, apiClient.auth)
	if err != nil {
		return nil, err
	}

	return res.Payload, nil
}
